// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package decoder provides several decoders for differently encoded trice streams.
package decoder

import (
	"encoding/binary"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"

	"github.com/rokath/trice/internal/emitter"
	"github.com/rokath/trice/internal/id"
)

// TestTable ist a struct slice generated by the trice tool -testTable option.
type TestTable []struct {
	In  []byte // byte buffer sequence
	Exp string // output
}

const (

	// LittleEndian is true for little endian trice data.
	LittleEndian = true

	// BigEndian is the flag value for target endianness.
	BigEndian = false

	// defaultSize is the beginning receive and sync buffer size.
	DefaultSize = 64 * 1024

	// patNextFormatSpecifier is a regex to find next format specifier in a string (exclude %%*) and ignoring %s
	//
	// https://regex101.com/r/BjiD5M/1
	// Language C plus from language Go: %b, %F, %q
	// Partial implemented: %hi, %hu, %ld, %li, %lf, %Lf, %Lu, %lli, %lld
	// Not implemented: %s
	patNextFormatSpecifier = `%([+\-#'0-9\.0-9])*(b|c|d|e|f|g|E|F|G|h|i|l|L|n|o|O|p|q|t|u|x|X)` // assumes no `%%` inside string!

	// patNextFormatUSpecifier is a regex to find next format u specifier in a string
	// It does also match %%u positions!
	patNextFormatUSpecifier = `%[0-9]*u` // assumes no `%%` inside string!

	// patNextFormatISpecifier is a regex to find next format i specifier in a string
	// It does also match %%i positions!
	patNextFormatISpecifier = `%[0-9]*i` // assumes no `%%` inside string!

	// patNextFormatXSpecifier is a regex to find next format x specifier in a string
	// It does also match %%x positions!
	patNextFormatXSpecifier = `%[0-9]*(l|o|O|x|X|b|p|t)` // assumes no `%%` inside string!

	// patNextFormatFSpecifier is a regex to find next format f specifier in a string
	// It does also match %%f positions!
	patNextFormatFSpecifier = `%[(+\-0-9\.0-9#]*(e|E|f|F|g|G)` // assumes no `%%` inside string!

	// patNextFormatBoolSpecifier is a regex to find next format f specifier in a string
	// It does also match %%t positions!
	patNextFormatBoolSpecifier = `%t` // assumes no `%%` inside string!

	// patNextFormatPointerSpecifier is a regex to find next format f specifier in a string
	// It does also match %%t positions!
	patNextFormatPointerSpecifier = `%p` // assumes no `%%` inside string!

	// hints is the help information in case of errors.
	Hints = "att:Hints:Baudrate? Encoding? Interrupt? Overflow? Parameter count? Format specifier? Password? til.json? Version?"

	UnsignedFormatSpecifier = 0 // %u -> %d
	SignedFormatSpecifier   = 1 //
	FloatFormatSpecifier    = 2 // %f and relatives
	BooleanFormatSpecifier  = 3 // a %t (bool) found
	PointerFormatSpecifier  = 4 // a %p (pointer) found

)

var (
	// Verbose gives more information on output if set. The value is injected from main packages.
	Verbose bool

	// ShowID is used as format string for displaying the first trice ID at the start of each line if not "".
	ShowID string

	// decoder.LastTriceID is last decoded ID. It is used for switch -showID.
	LastTriceID id.TriceID

	// TestTableMode is a special option for easy decoder test table generation.
	TestTableMode bool

	// Unsigned if true, forces hex and in values printed as unsigned values.
	Unsigned bool

	matchNextFormatSpecifier        = regexp.MustCompile(patNextFormatSpecifier)
	matchNextFormatUSpecifier       = regexp.MustCompile(patNextFormatUSpecifier)
	matchNextFormatISpecifier       = regexp.MustCompile(patNextFormatISpecifier)
	matchNextFormatXSpecifier       = regexp.MustCompile(patNextFormatXSpecifier)
	matchNextFormatFSpecifier       = regexp.MustCompile(patNextFormatFSpecifier)
	matchNextFormatBoolSpecifier    = regexp.MustCompile(patNextFormatBoolSpecifier)
	matchNextFormatPointerSpecifier = regexp.MustCompile(patNextFormatPointerSpecifier)

	DebugOut                        = false // DebugOut enables debug information.
	DumpLineByteCount               int     // DumpLineByteCount is the bytes per line for the dumpDec decoder.
	InitialCycle                    = true  // InitialCycle is a helper for the cycle counter automatic.
	TargetTimestamp                 uint64  // targetTimestamp contains target specific timestamp value.
	TargetLocation                  uint32  // targetLocation contains 16 bit file id in high and 16 bit line number in low part.
	TargetStamp                     string  // TargetTimeStampUnit is the target timestamps time base for default formatting.
	TargetStamp32                   string  // ShowTargetStamp32 is the format string for target timestamps.
	TargetStamp16                   string  // ShowTargetStamp16 is the format string for target timestamps.
	TargetStamp0                    string  // ShowTargetStamp0 is the format string for target timestamps.
	TargetTimeStampUnitPassed       bool    // TargetTimeStampUnitPassed is true when flag was TargetTimeStampUnit passed.
	ShowTargetStamp32Passed         bool    // ShowTargetStamp32Passed is true when flag was TargetTimeStamp32 passed.
	ShowTargetStamp16Passed         bool    // ShowTargetStamp16Passed is true when flag was TargetTimeStamp16 passed.
	ShowTargetStamp0Passed          bool    // ShowTargetStamp0Passed is true when flag was TargetTimeStamp0 passed.
	LocationInformationFormatString string  // LocationInformationFormatString is the format string for target location: line number and file name.
	TargetTimestampSize             int     // TargetTimestampSize is set in dependence of trice type.
	TargetLocationExists            bool    // TargetLocationExists is set in dependence of p.COBSModeDescriptor. (obsolete)

	PackageFraming  string // Framing is used for packing. Valid values COBS, TCOBS, TCOBSv1 (same as TCOBS)
	IDBits          = 14   // IDBits holds count of bits used for ID (used at least in trexDecoder)
	NewlineIndent   = -1   // Used for trice messages containing several newlines in format string for formatting.
	TriceStatistics bool   // Keep the occured count for each Trice log when Trice is closed.
	IDStat          map[id.TriceID]int
)

func init() {
	IDStat = make(map[id.TriceID]int)
}

// New abstracts the function type for a new decoder.
type New func(out io.Writer, lut id.TriceIDLookUp, m *sync.RWMutex, li id.TriceIDLookUpLI, in io.Reader, endian bool) Decoder

// Decoder is providing a byte reader returning decoded trice's.
// SetInput allows switching the input stream to a different source.
type Decoder interface {
	io.Reader
	SetInput(io.Reader)
}

// DecoderData is the common data struct for all decoders.
type DecoderData struct {
	W           io.Writer          // io.Stdout or the like
	In          io.Reader          // in is the inner reader, which is used to get raw bytes
	InnerBuffer []byte             // avoid repeated allocation (trex)
	IBuf        []byte             // iBuf holds unprocessed (raw) bytes for interpretation.
	B           []byte             // read buffer holds a single decoded TCOBS package, which can contain several trices.
	B0          []byte             // initial value for B
	Endian      bool               // endian is true for LittleEndian and false for BigEndian
	TriceSize   int                // trice head and payload size as number of bytes
	ParamSpace  int                // trice payload size after head
	SLen        int                // string length for TRICE_S
	Lut         id.TriceIDLookUp   // id look-up map for translation
	LutMutex    *sync.RWMutex      // to avoid concurrent map read and map write during map refresh triggered by filewatcher
	Li          id.TriceIDLookUpLI // location information map
	Trice       id.TriceFmt        // id.TriceFmt // received trice
}

// SetInput allows switching the input stream to a different source.
//
// This function is for easier testing with cycle counters.
func (p *DecoderData) SetInput(r io.Reader) {
	p.In = r
}

// ReadU16 returns the 2 b bytes as uint16 according the specified endianness
func (p *DecoderData) ReadU16(b []byte) uint16 {
	if p.Endian {
		return binary.LittleEndian.Uint16(b)
	}
	return binary.BigEndian.Uint16(b)
}

// ReadU32 returns the 4 b bytes as uint32 according the specified endianness
func (p *DecoderData) ReadU32(b []byte) uint32 {
	if p.Endian {
		return binary.LittleEndian.Uint32(b)
	}
	return binary.BigEndian.Uint32(b)
}

// ReadU64 returns the 8 b bytes as uint64 according the specified endianness
func (p *DecoderData) ReadU64(b []byte) uint64 {
	if p.Endian {
		return binary.LittleEndian.Uint64(b)
	}
	return binary.BigEndian.Uint64(b)
}

// UReplaceN checks all format specifier in i and replaces %nu with %nd and returns that result as o.
//
// If a replacement took place on position k u[k] is 1. Afterwards len(u) is amount of found format specifiers.
// Additional, if UnsignedHex is true, for FormatX specifiers u[k] is also 1.
// If a float format specifier was found at position k, u[k] is 2,
// http://www.cplusplus.com/reference/cstdio/printf/
// https://www.codingunit.com/printf-format-specifiers-format-conversions-and-formatted-output
func UReplaceN(i string) (o string, u []int) {
	o = i
	i = strings.ReplaceAll(i, "%%", "__") // this makes regex easier and faster
	var offset int
	for {
		s := i[offset:] // remove processed part
		loc := matchNextFormatSpecifier.FindStringIndex(s)
		if nil == loc { // no (more) fm found
			return
		}
		offset += loc[1] // track position
		fm := s[loc[0]:loc[1]]
		locPointer := matchNextFormatPointerSpecifier.FindStringIndex(fm)
		if nil != locPointer { // a %p found
			// This would require `unsafe.Pointer(uintptr(n))` inside unSignedOrSignedOut.
			// There are false positive windows vet warnings:
			// https://stackoverflow.com/questions/43767898/casting-a-int-to-a-pointer
			// https://github.com/golang/go/issues/41205
			// As workaround replace %p with %x in the format strings.
			// Then trice64( "%p", -1 ) could be a problem when using `trice log -unsigned false`
			// But that we simply ignore right now.
			o = o[:offset-1] + "x" + o[offset:]   // replace %np -> %nx
			u = append(u, PointerFormatSpecifier) // pointer value
			continue
		}
		locBool := matchNextFormatBoolSpecifier.FindStringIndex(fm)
		if nil != locBool { // a %t found
			u = append(u, BooleanFormatSpecifier) // bool value
			continue
		}
		locF := matchNextFormatFSpecifier.FindStringIndex(fm)
		if nil != locF { // a %nf found
			u = append(u, FloatFormatSpecifier) // float value
			continue
		}
		locU := matchNextFormatUSpecifier.FindStringIndex(fm)
		if nil != locU { // a %nu found
			o = o[:offset-1] + "d" + o[offset:]    // replace %nu -> %nd
			u = append(u, UnsignedFormatSpecifier) // no negative values
			continue
		}
		locI := matchNextFormatISpecifier.FindStringIndex(fm)
		if nil != locI { // a %ni found
			o = o[:offset-1] + "d" + o[offset:]  // replace %ni -> %nd
			u = append(u, SignedFormatSpecifier) // also negative values
			continue
		}
		locX := matchNextFormatXSpecifier.FindStringIndex(fm)
		if nil != locX { // a %nx, %nX or, %no, %nO or %nb found
			if Unsigned {
				u = append(u, 0) // no negative values
			} else {
				u = append(u, 1) // also negative values
			}
			continue
		}
		u = append(u, 1) // keep sign in all other cases(also negative values)
	}
}

// Dump prints the byte slice as hex in one line
func Dump(w io.Writer, b []byte) {
	for _, x := range b {
		fmt.Fprintf(w, "%02x ", x)
	}
	fmt.Fprintln(w, "")
}

func RecordForStatistics(tid id.TriceID) {
	if !TriceStatistics && !emitter.AllStatistics {
		return
	}
	count := IDStat[tid]
	count++
	IDStat[tid] = count
}

var (
	IDLUT id.TriceIDLookUp
	LILUT id.TriceIDLookUpLI
)

func PrintTriceStatistics(w io.Writer) {
	if !TriceStatistics && !emitter.AllStatistics {
		return
	}
	var sum int
	fmt.Fprintf(w, "\nTrice Statistics:                  (n: Trice ends with no newline, if 0 ↴)\n\n")
	fmt.Fprintln(w, `   Count |       Location Information       | Line |  ID   |    Type    |n| Format String`)
	fmt.Fprintln(w, " ------- | -------------------------------- | ---- | ----- | ---------- |-|----------------------------------------")
	for tid, count := range IDStat {
		sum += count
		trice := IDLUT[tid]
		li := LILUT[tid]
		var found bool
		trice.Strg, found = strings.CutSuffix(trice.Strg, `\n`)
		newline := ` `
		if !found {
			newline = `0`
		}
		fmt.Fprintf(w, "%8d | %32s |%5d | %5d | %10s |%s| %s\n", count, li.File, li.Line, tid, trice.Type, newline, emitter.Colorize(trice.Strg))
	}
	fmt.Fprintln(w, " ------------------------------------------------------------------------------------------------------------------")
	fmt.Fprintf(w, "%8d Trice messsges\n", sum)
}

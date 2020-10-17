// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package translator converts...
package translator

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/rokath/trice/internal/id"
	"github.com/rokath/trice/internal/receiver"
)

const (
	redBalk = "error:\ne:                                                      \nerror:                                                      \n"
)

var (
	// Verbose gives more information on output if set. This variable is set outside this package.
	Verbose bool
)

// Trice is the bare Trice data type for a Trice atom.
// A Trice starts with zero or several Trice atoms with ID==0 carrying parts of the Trice data payload.
// The last Trice atom of a Trice contains the Trice ID!=0 and the last part of the data payload.
type Trice = receiver.Trice

// TriceAtomsReceiver is the interface a trice receicer has to provide for a trice interpreter.
// The provided channels are read only channels.
type TriceAtomsReceiver interface {
	TriceAtomsChannel() <-chan []Trice
	IgnoredBytesChannel() <-chan []byte
}

// translator is the common data cantainmer with the common methods for BareTranslator, EscTranslator and the like.
type translator struct {
	list     *id.List
	item     id.Item // item is the next trice item ready for output.
	savedErr error
	done     chan int // This channel is used to stop the TriceInterpreter
}

// Translator is the common interface for BareTranslator, EscTranslator and the like.
type Translator interface {
	SavedError() error
	Done() chan int
}

func (p *translator) SavedError() error {
	return p.savedErr
}

func (p *translator) Done() chan int {
	return p.done
}

// ErrorFatal ends in osExit(1) if p.err not nil.
func (p *translator) ErrorFatal() {
	if nil == p.savedErr {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	log.Fatal(p.savedErr, filepath.Base(file), line)
}

// // String is the method for displaying the current TriceTranslator instance.
// func (p *TriceTranslator) String() string {
// 	s := fmt.Sprintf("ID=%d=%x", p.item.ID, p.item.ID)
// 	s += fmt.Sprintf("values in hex: ")
// 	for _, n := range p.values {
// 		s += fmt.Sprintf("%x ", n)
// 	}
// 	return s
// }

// parse lang C formatstring for %u and replace them with %d and extend the
// returned slice with 0 for each %d, %c, %x and 1 for each converted %u
// http://www.cplusplus.com/reference/cstdio/printf/
//
// Changes:        C    Go
// ---------------------------
// int 	           %i   %d replace i with d
// unsigned int    %u   %d replace u with d and uint()
//
// No changes:     C    Go
// ---------------------------
// decimal         %d   %d
// hexadezimal     %x   %x
// HEXADECIMAL     %X   %X
// octal           %o   %o
// float           %f   %f
// char            %c   %c
// (void *)        %p   %p
func langCtoGoFmtStingConverter(f string) (s string, u []bool, err error) {
	s = f // TODO: parse f for %*u, replace with %*d and sett []bool values accordingly
	return
}

/*
// NewEscTrices uses in for reception and sw for writing.
// It collects trice bytes to a complete esc trice message, generates the appropriate string using list and writes it to sw.
// EC LC IH IL ...
func NewEscTrices(sw io.StringWriter, list *id.List, in io.ReadCloser, hardReadError chan bool) *EscTranslator {
	p := new(EscTranslator)
	p.sw = sw
	p.list = list
	p.in = in
	p.syncBuffer = make([]byte, 0, 1000)
	p.done = make(chan int)
	go func() {
		for {
			select {
			case <-p.done:
				if Verbose {
					fmt.Println("end of esc translator life")
				}
				return

			case <-time.After(1 * time.Millisecond): // todo: trigger from fileWatcher
				if io.EOF == p.savedErr {
					p.savedErr = nil
				}
				s := p.readEsc()
				if nil != p.savedErr && io.EOF != p.savedErr {
					fmt.Println("Read error", p.savedErr)
					in.Close()
					hardReadError <- true
					return
				}
				_, p.savedErr = sw.WriteString(s)
			}
		}

	}()
	return p
}

func (p *EscTranslator) readEsc() (s string) {
	p.ErrorFatal()
	var n int
	rb := make([]byte, 100)
	n, p.savedErr = p.in.Read(rb)
	p.syncBuffer = append(p.syncBuffer, rb[:n]...) // merge with leftovers
parse:
	if len(p.syncBuffer) < 4 {
		return
	}
	for _, c := range p.syncBuffer {
		if 0xec != c {
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse                      // no start byte
		}
		if 0xde == p.syncBuffer[1] {
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse                      // no start byte: `0xec 0xde` is no valid esc packet start. Is is an escaped 0xec.
		}

		// p.syncBuffer[0] is esc, p.syncBuffer[1] is length code
		id := (int(p.syncBuffer[2]) << 8) | int(p.syncBuffer[3])
		index := p.list.Index(id)
		if index < 0 { // unknown id
			s = redBalk + fmt.Sprintln("error: unknown id", id, "syncBuffer = ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		}
		p.item = p.list.Item(index)

		switch p.syncBuffer[1] { // length code
		case 0xdf: // no params
			if "TRICE0" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg)
				p.syncBuffer = p.syncBuffer[4:]
				return
			}
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		case 0xe0: // 1 byte param
			if len(p.syncBuffer) < 5 {
				return // wait
			}
			if "TRICE8_1" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg, int8(p.syncBuffer[4]))
				p.syncBuffer = p.syncBuffer[5:]
				return
			}
			s = redBalk + fmt.Sprintln("error: ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		case 0xe1: // 2 bytes param
			if len(p.syncBuffer) < 6 {
				return // wait
			}
			if "TRICE8_2" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]))
				p.syncBuffer = p.syncBuffer[6:]
				return
			}
			if "TRICE16_1" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg, int16(binary.BigEndian.Uint16(p.syncBuffer[4:6])))
				p.syncBuffer = p.syncBuffer[6:]
				return
			}
			s = redBalk + fmt.Sprintln("error: ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		case 0xe2: // 4 bytes param
			if len(p.syncBuffer) < 7 {
				return // wait
			}
			if "TRICE8_3" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]),
					int8(p.syncBuffer[6]))
				p.syncBuffer = p.syncBuffer[7:]
				return
			}
			if len(p.syncBuffer) < 8 {
				return // wait
			}
			if "TRICE8_4" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]),
					int8(p.syncBuffer[6]),
					int8(p.syncBuffer[7]))
				p.syncBuffer = p.syncBuffer[8:]
				return
			}
			if "TRICE16_2" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int16(binary.BigEndian.Uint16(p.syncBuffer[4:6])),
					int16(binary.BigEndian.Uint16(p.syncBuffer[6:8])))
				p.syncBuffer = p.syncBuffer[8:]
				return
			}
			if "TRICE32_1" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg, int32(binary.BigEndian.Uint32(p.syncBuffer[4:8])))
				p.syncBuffer = p.syncBuffer[8:]
				return
			}
			s = redBalk + fmt.Sprintln("error: ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		case 0xe3: // 8 bytes param
			if "TRICE8_5" == p.item.FmtType {
				if len(p.syncBuffer) < 9 {
					return // wait
				}
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]),
					int8(p.syncBuffer[6]),
					int8(p.syncBuffer[7]),
					int8(p.syncBuffer[8]))
				p.syncBuffer = p.syncBuffer[9:]
				return
			}
			if len(p.syncBuffer) < 10 {
				return // wait
			}
			if "TRICE8_6" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]),
					int8(p.syncBuffer[6]),
					int8(p.syncBuffer[7]),
					int8(p.syncBuffer[8]),
					int8(p.syncBuffer[9]))
				p.syncBuffer = p.syncBuffer[10:]
				return
			}
			if "TRICE16_3" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int16(binary.BigEndian.Uint16(p.syncBuffer[4:6])),
					int16(binary.BigEndian.Uint16(p.syncBuffer[6:8])),
					int16(binary.BigEndian.Uint16(p.syncBuffer[8:10])))
				p.syncBuffer = p.syncBuffer[10:]
				return
			}
			if len(p.syncBuffer) < 11 {
				return // wait
			}
			if "TRICE8_7" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]),
					int8(p.syncBuffer[6]),
					int8(p.syncBuffer[7]),
					int8(p.syncBuffer[8]),
					int8(p.syncBuffer[9]),
					int8(p.syncBuffer[10]))
				p.syncBuffer = p.syncBuffer[11:]
				return
			}
			if len(p.syncBuffer) < 12 {
				return // wait
			}
			if "TRICE8_8" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int8(p.syncBuffer[4]),
					int8(p.syncBuffer[5]),
					int8(p.syncBuffer[6]),
					int8(p.syncBuffer[7]),
					int8(p.syncBuffer[8]),
					int8(p.syncBuffer[9]),
					int8(p.syncBuffer[10]),
					int8(p.syncBuffer[11]))
				p.syncBuffer = p.syncBuffer[12:]
				return
			}
			if "TRICE16_4" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int16(binary.BigEndian.Uint16(p.syncBuffer[4:6])),
					int16(binary.BigEndian.Uint16(p.syncBuffer[6:8])),
					int16(binary.BigEndian.Uint16(p.syncBuffer[8:10])),
					int16(binary.BigEndian.Uint16(p.syncBuffer[10:12])))
				p.syncBuffer = p.syncBuffer[12:]
				return
			}
			if "TRICE32_2" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int32(binary.BigEndian.Uint32(p.syncBuffer[4:8])),
					int32(binary.BigEndian.Uint32(p.syncBuffer[8:12])))
				p.syncBuffer = p.syncBuffer[12:]
				return
			}
			if "TRICE64_1" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int64(binary.BigEndian.Uint64(p.syncBuffer[4:12])))
				p.syncBuffer = p.syncBuffer[12:]
				return
			}
			s = redBalk + fmt.Sprintln("error: ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		case 0xe4: // 16 bytes param
			if len(p.syncBuffer) < 16 {
				return // wait
			}
			if "TRICE32_3" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int32(binary.BigEndian.Uint32(p.syncBuffer[4:8])),
					int32(binary.BigEndian.Uint32(p.syncBuffer[8:12])),
					int32(binary.BigEndian.Uint32(p.syncBuffer[12:16])))
				p.syncBuffer = p.syncBuffer[16:]
				return
			}
			if len(p.syncBuffer) < 20 {
				return // wait
			}
			if "TRICE32_4" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int32(binary.BigEndian.Uint32(p.syncBuffer[4:8])),
					int32(binary.BigEndian.Uint32(p.syncBuffer[8:12])),
					int32(binary.BigEndian.Uint32(p.syncBuffer[12:16])),
					int32(binary.BigEndian.Uint32(p.syncBuffer[16:20])))
				p.syncBuffer = p.syncBuffer[20:]
				return
			}
			if "TRICE64_2" == p.item.FmtType {
				s = fmt.Sprintf(p.item.FmtStrg,
					int64(binary.BigEndian.Uint64(p.syncBuffer[4:12])),
					int64(binary.BigEndian.Uint64(p.syncBuffer[12:20])))
				p.syncBuffer = p.syncBuffer[20:]
				return
			}
			s = redBalk + fmt.Sprintln("error: ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:] // remove 1st char
			goto parse
		default:
			s = redBalk + fmt.Sprintln("error: ", p.syncBuffer)
			p.syncBuffer = p.syncBuffer[1:]
			goto parse // invalid length code
		}
	}
	return
}
*/

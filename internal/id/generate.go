// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package id List is responsible for id List managing
package id

// List management

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/rokath/trice/internal/emitter"
	"github.com/rokath/trice/pkg/msg"
	"github.com/spf13/afero"
)

var (
	GenerateHFile           bool
	GenerateCFile           bool
	GenerateCSFile          bool
	WriteAllColors          bool
	IDToFunctionPointerList bool
)

// SubCmdIdGenerate performs sub-command generate, creating support files/output.
func SubCmdGenerate(w io.Writer, fSys *afero.Afero) error {
	if GenerateHFile {
		fmt.Fprintln(w, `CLI "-h" not implemented yet`)
	}
	if GenerateCFile {
		fmt.Fprintln(w, `CLI "-c" not implemented yet`)
	}
	if GenerateCSFile {
		fmt.Fprintln(w, `CLI "-cs" not implemented yet`)
	}
	if IDToFunctionPointerList {
		fmt.Fprintln(w, `CLI "-fpl" not implemented yet`)
	}
	if WriteAllColors {
		emitter.ShowAllColors()
		fmt.Fprintln(w, `Modify ansi.ColorFunc assignments in lineTransformerANSI.go to change Trice colors.`)

	}
	return nil
}

// ToFileLangC generates lang:C helpers for a third party tool.
func (ilu TriceIDLookUp) ToFileLangC(fSys afero.Fs, fn string) (err error) {
	var fC, fH, fCS afero.File
	fnC := fn + ".c"
	fC, err = fSys.Create(fnC)
	msg.FatalOnErr(err)
	fnH := fn + ".h"
	fH, err = fSys.Create(fnH)
	msg.FatalOnErr(err)
	defer func() {
		err = fC.Close()
		msg.FatalOnErr(err)
		err = fH.Close()
		msg.FatalOnErr(err)
		err = fCS.Close()
		msg.FatalOnErr(err)
	}()

	c, e := ilu.toCFmtList(fnC)
	_, err = fC.Write(c)
	msg.FatalOnErr(e)

	h := []byte(`//! \file ` + fnH + `
		//! ///////////////////////////////////////////////////////////////////////////

		//! generated code - do not edit!

		#include <stdint.h>

		typedef struct{
		char* formatString;
		uint16_t id;
		int16_t dataLength;
		uint8_t bitWidth;
		} triceFormatStringList_t;

		extern const triceFormatStringList_t triceFormatStringList[];
		extern const unsigned triceFormatStringListElements;

		`)
	_, e = fH.Write(h)
	msg.FatalOnErr(e)
	return
}

// ToFileCSharp generates C# helpers for a third party tool.
func (ilu TriceIDLookUp) ToFileCSharp(fSys afero.Fs, fn string) (err error) {
	var fCS afero.File

	fnCS := fn + ".cs"
	fCS, err = fSys.Create(fnCS)
	msg.FatalOnErr(err)
	defer func() {
		err = fCS.Close()
		msg.FatalOnErr(err)
	}()

	cs, e := ilu.toCSFmtList(fnCS)
	_, err = fCS.Write(cs)
	msg.FatalOnErr(e)
	return
}

// toCFmtList converts lim into C-source byte slice in human-readable form.
func (ilu TriceIDLookUp) toCFmtList(fileName string) ([]byte, error) {
	fileNameBody := fileNameWithoutSuffix(filepath.Base(fileName))
	c := []byte(`//! \file ` + fileNameBody + `.c
//! ///////////////////////////////////////////////////////////////////////////

//! generated code - do not edit!

#include "` + fileNameBody + `.h"

//! triceFormatStringList contains all trice format strings together with id and parameter information.
const triceFormatStringList_t triceFormatStringList[] = {
	// format-string,                                                                     id, dataLength, bitWidth,
`)
	var s string
	var paramCount int
	var bitWidth int
	var dataLength int
	var add bool

	for id, k := range ilu {
		s = k.Strg
		switch k.Type {

		case "TRICE0", "TRICE", "TRICE32", "TRICE32_1", "TRICE32_2", "TRICE32_3", "TRICE32_4", "TRICE32_5", "TRICE32_6", "TRICE32_7", "TRICE32_8", "TRICE32_9", "TRICE32_10", "TRICE32_11", "TRICE32_12":
			bitWidth = 32
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 4 // use for checks
			add = true
		case "TRICE16", "TRICE16_1", "TRICE16_2", "TRICE16_3", "TRICE16_4", "TRICE16_5", "TRICE16_6", "TRICE16_7", "TRICE16_8", "TRICE16_9", "TRICE16_10", "TRICE16_11", "TRICE16_12":
			bitWidth = 16
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 2 // use for checks
			add = true
		case "TRICE8", "TRICE8_1", "TRICE8_2", "TRICE8_3", "TRICE8_4", "TRICE8_5", "TRICE8_6", "TRICE8_7", "TRICE8_8", "TRICE8_9", "TRICE8_10", "TRICE8_11", "TRICE8_12":
			bitWidth = 8
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 1 // use for checks
			add = true
		case "TRICE64", "TRICE64_1", "TRICE64_2", "TRICE64_3", "TRICE64_4", "TRICE64_5", "TRICE64_6", "TRICE64_7", "TRICE64_8", "TRICE64_9", "TRICE64_10", "TRICE64_11", "TRICE64_12":
			bitWidth = 64
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 8 // use for checks
			add = true
		case "TRICE_S":
			bitWidth = 8
			paramCount = 1
			dataLength = -1 // needs to be derived
			add = true
		case "TRICE_N":
			bitWidth = 8
			paramCount = 1
			dataLength = -2 // needs to be derived, add terminating 0
			add = true
		case "TRICE_B", "TRICE8_B":
			bitWidth = 8
			paramCount = 1
			dataLength = -8 // needs to be derived
			add = false     // provide generated format string
		case "TRICE16_B":
			bitWidth = 16
			paramCount = 1
			dataLength = -16 // needs to be derived
			add = false      // provide generated format string
		case "TRICE32_B":
			bitWidth = 32
			paramCount = 1
			dataLength = -32 // needs to be derived
			add = false      // provide generated format string
		case "TRICE64_B":
			bitWidth = 64
			paramCount = 1
			dataLength = -64 // needs to be derived
			add = false      // provide generated format string
		case "TRICE8_F", "TRICE_F": // example: TRICE_F( Id( 2844), "info:FunctionNameW",  b8,  sizeof(b8) /sizeof(int8_t) );
			bitWidth = 8
			paramCount = 1
			dataLength = -9 // needs to be derived
			add = false     // provide generated format string
		case "TRICE16_F":
			bitWidth = 16
			paramCount = 1
			dataLength = -17 // needs to be derived
			add = false      // provide generated format string
		case "TRICE32_F":
			bitWidth = 32
			paramCount = 1
			dataLength = -33 // needs to be derived
			add = false      // provide generated format string
		case "TRICE64_F":
			bitWidth = 64
			paramCount = 1
			dataLength = -65 // needs to be derived
			add = false      // provide generated format string
		default:
			add = false
		}
		if add {
			c = append(c, []byte(fmt.Sprintf(`    { "%s",%s%5d, %3d, %2d },`+"\n", s, distance(s), id, dataLength, bitWidth))...)
		}
	}

	tail := []byte(`};

//! triceFormatStringListElements holds the compile time computed count of list elements.
const unsigned triceFormatStringListElements = sizeof(triceFormatStringList) / sizeof(triceFormatStringList_t);
`)
	c = append(c, tail...)
	return c, nil
}

// toCSFmtList converts ilu into CS-source byte slice in human-readable form.
func (ilu TriceIDLookUp) toCSFmtList(fileName string) ([]byte, error) {
	c := []byte(`// generated code - do not edit!

// There is still a need to exchange the format specifier from C to C#.
// See https://stackoverflow.com/questions/33432341/how-to-use-c-language-format-specifiers-in-c-sharp
// and https://www.codeproject.com/Articles/19274/A-printf-implementation-in-C for possible help.

namespace TriceIDList;

    public class TilItem
    {
        public TilItem(string strg, int bitWidth, int size)
        {
            Strg = strg;
            BitWidth = bitWidth;
            Size = size;
        }

        public string Strg { get; init; }
        public int BitWidth { get; init; }
        public int Size { get; init; }
    }

    public static class Til
    {
        public static readonly Dictionary<int, TilItem> TilList= new Dictionary<int, TilItem>
        { // id, TilItem ( Strg, bitWidth, dataLength )
`)
	var s string
	var paramCount int
	var bitWidth int
	var dataLength int
	var add bool

	for id, k := range ilu {
		s = k.Strg
		switch k.Type {

		case "TRICE0", "TRICE", "TRICE32", "TRICE32_1", "TRICE32_2", "TRICE32_3", "TRICE32_4", "TRICE32_5", "TRICE32_6", "TRICE32_7", "TRICE32_8", "TRICE32_9", "TRICE32_10", "TRICE32_11", "TRICE32_12":
			bitWidth = 32
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 4 // use for checks
			add = true
		case "TRICE16", "TRICE16_1", "TRICE16_2", "TRICE16_3", "TRICE16_4", "TRICE16_5", "TRICE16_6", "TRICE16_7", "TRICE16_8", "TRICE16_9", "TRICE16_10", "TRICE16_11", "TRICE16_12":
			bitWidth = 16
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 2 // use for checks
			add = true
		case "TRICE8", "TRICE8_1", "TRICE8_2", "TRICE8_3", "TRICE8_4", "TRICE8_5", "TRICE8_6", "TRICE8_7", "TRICE8_8", "TRICE8_9", "TRICE8_10", "TRICE8_11", "TRICE8_12":
			bitWidth = 8
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 1 // use for checks
			add = true
		case "TRICE64", "TRICE64_1", "TRICE64_2", "TRICE64_3", "TRICE64_4", "TRICE64_5", "TRICE64_6", "TRICE64_7", "TRICE64_8", "TRICE64_9", "TRICE64_10", "TRICE64_11", "TRICE64_12":
			bitWidth = 64
			paramCount = formatSpecifierCount(s)
			dataLength = paramCount * 8 // use for checks
			add = true
		case "TRICE_S":
			bitWidth = 8
			paramCount = 1
			dataLength = -1 // needs to be derived
			add = true
		case "TRICE_N":
			bitWidth = 8
			paramCount = 1
			dataLength = -2 // needs to be derived, add terminating 0
			add = true
		case "TRICE_B", "TRICE8_B":
			bitWidth = 8
			paramCount = 1
			dataLength = -8 // needs to be derived
			add = false     // provide generated format string
		case "TRICE16_B":
			bitWidth = 16
			paramCount = 1
			dataLength = -16 // needs to be derived
			add = false      // provide generated format string
		case "TRICE32_B":
			bitWidth = 32
			paramCount = 1
			dataLength = -32 // needs to be derived
			add = false      // provide generated format string
		case "TRICE64_B":
			bitWidth = 64
			paramCount = 1
			dataLength = -64 // needs to be derived
			add = false      // provide generated format string
		case "TRICE8_F", "TRICE_F": // example: TRICE_F( Id( 2844), "info:FunctionNameW",  b8,  sizeof(b8) /sizeof(int8_t) );
			bitWidth = 8
			paramCount = 1
			dataLength = -9 // needs to be derived
			add = false     // provide generated format string
		case "TRICE16_F":
			bitWidth = 16
			paramCount = 1
			dataLength = -17 // needs to be derived
			add = false      // provide generated format string
		case "TRICE32_F":
			bitWidth = 32
			paramCount = 1
			dataLength = -33 // needs to be derived
			add = false      // provide generated format string
		case "TRICE64_F":
			bitWidth = 64
			paramCount = 1
			dataLength = -65 // needs to be derived
			add = false      // provide generated format string
		default:
		}
		if add {
			c = append(c, []byte(fmt.Sprintf(`        { %5d, new TilItem( "%s", %d, %d ) },`+"\n", id, s, bitWidth, dataLength))...)
			add = false
		}
	}

	tail := []byte(`    };
}

`)
	c = append(c, tail...)
	return c, nil
}

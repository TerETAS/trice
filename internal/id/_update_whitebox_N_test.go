// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.
package id

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A wrong parameter count should not be corrected! The compiler will complain and a decision should be made.
func TestDoNotCorrectWrongParamCountSingle(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{
			`Trice8_2( Id(0), "hi %2d",1  );`,
			`Trice8_2( Id(0), "hi %2d",1  );`},
		{
			`TRICE8_2( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8_2( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
	}
	checkTestTable(t, tt, true)
}

// A wrong parameter count should not be corrected! The compiler will complain and a decision should be made.
// todo: emit a warning
func TestDoNotCorrectWrongParamCountSingle2(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{
			`Trice8_2( iD(0), "hi %2d",1  );`,
			`Trice8_2( iD(0), "hi %2d",1  );`},
		{
			`TRICE8_2( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8_2( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertID0NoParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`... TRice ( "hi"); ...`, `... TRice ( iD(0), "hi"); ...`},
		{`... Trice ( "hi"); ...`, `... Trice ( iD(0), "hi"); ...`},
		{`... trice ( "hi"); ...`, `... trice ( iD(0), "hi"); ...`},
	}
	checkTestTable(t, tt, true)
}
func TestInsertParamCountAndIDNoParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`... TRice0 ( "hi"); ...`, `... TRice0 ( iD(0), "hi"); ...`},
		{`... Trice0 ( "hi"); ...`, `... Trice0 ( iD(0), "hi"); ...`},
		{`... trice0 ( "hi"); ...`, `... trice0 ( iD(0), "hi"); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDOneParamN1(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	StampSizeId = " Id(0),"
	tt := []struct{ text, exp string }{
		{`...  TRICE8 ( "hi %03u", 5); ...`, `...  TRICE8_1 ( Id(0), "hi %03u", 5); ...`},
		{`... TRICE16 ( "hi %03u", 5); ...`, `... TRICE16_1 ( Id(0), "hi %03u", 5); ...`},
		{`... TRICE32 ( "hi %03u", 5); ...`, `... TRICE32_1 ( Id(0), "hi %03u", 5); ...`},
		{`... TRICE64 ( "hi %03u", 5); ...`, `... TRICE64_1 ( Id(0), "hi %03u", 5); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDOneParamB(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	StampSizeId = " Id(0),"
	tt := []struct{ text, exp string }{
		{`...  TRICE8 ( "hi %03u", 5); ...`, `...  TRICE8 ( Id(0), "hi %03u", 5); ...`},
		{`... TRICE16 ( "hi %03u", 5); ...`, `... TRICE16 ( Id(0), "hi %03u", 5); ...`},
		{`... TRICE32 ( "hi %03u", 5); ...`, `... TRICE32 ( Id(0), "hi %03u", 5); ...`},
		{`... TRICE64 ( "hi %03u", 5); ...`, `... TRICE64 ( Id(0), "hi %03u", 5); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDOneParamN0(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...  trice8 ( "hi %03u", 5); ...`, `...  trice8_1 ( iD(0), "hi %03u", 5); ...`},
		{`...  Trice8 ( "hi %03u", 5); ...`, `...  Trice8_1 ( iD(0), "hi %03u", 5); ...`},
		{`...  TRice8 ( "hi %03u", 5); ...`, `...  TRice8_1 ( iD(0), "hi %03u", 5); ...`},
		//{`...  TRICE8 ( "hi %03u", 5); ...`, `...  TRICE8_1 ( Id(0), "hi %03u", 5); ...`},
		{`... trice16 ( "hi %03u", 5); ...`, `... trice16_1 ( iD(0), "hi %03u", 5); ...`},
		{`... Trice16 ( "hi %03u", 5); ...`, `... Trice16_1 ( iD(0), "hi %03u", 5); ...`},
		{`... TRice16 ( "hi %03u", 5); ...`, `... TRice16_1 ( iD(0), "hi %03u", 5); ...`},
		//{`... TRICE16 ( "hi %03u", 5); ...`, `... TRICE16_1 ( Id(0), "hi %03u", 5); ...`},
		{`... trice32 ( "hi %03u", 5); ...`, `... trice32_1 ( iD(0), "hi %03u", 5); ...`},
		{`... Trice32 ( "hi %03u", 5); ...`, `... Trice32_1 ( iD(0), "hi %03u", 5); ...`},
		{`... TRice32 ( "hi %03u", 5); ...`, `... TRice32_1 ( iD(0), "hi %03u", 5); ...`},
		//{`... TRICE32 ( "hi %03u", 5); ...`, `... TRICE32_1 ( Id(0), "hi %03u", 5); ...`},
		{`... trice64 ( "hi %03u", 5); ...`, `... trice64_1 ( iD(0), "hi %03u", 5); ...`},
		{`... Trice64 ( "hi %03u", 5); ...`, `... Trice64_1 ( iD(0), "hi %03u", 5); ...`},
		{`... TRice64 ( "hi %03u", 5); ...`, `... TRice64_1 ( iD(0), "hi %03u", 5); ...`},
		//{`... TRICE64 ( "hi %03u", 5); ...`, `... TRICE64_1 ( Id(0), "hi %03u", 5); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDOneParamA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...  trice8 ( "hi %03u", 5); ...`, `...  trice8 ( iD(0), "hi %03u", 5); ...`},
		{`...  Trice8 ( "hi %03u", 5); ...`, `...  Trice8 ( iD(0), "hi %03u", 5); ...`},
		{`...  TRice8 ( "hi %03u", 5); ...`, `...  TRice8 ( iD(0), "hi %03u", 5); ...`},
		//{`...  TRICE8 ( "hi %03u", 5); ...`, `...  TRICE8 ( Id(0), "hi %03u", 5); ...`},
		{`... trice16 ( "hi %03u", 5); ...`, `... trice16 ( iD(0), "hi %03u", 5); ...`},
		{`... Trice16 ( "hi %03u", 5); ...`, `... Trice16 ( iD(0), "hi %03u", 5); ...`},
		{`... TRice16 ( "hi %03u", 5); ...`, `... TRice16 ( iD(0), "hi %03u", 5); ...`},
		//{`... TRICE16 ( "hi %03u", 5); ...`, `... TRICE16 ( Id(0), "hi %03u", 5); ...`},
		{`... trice32 ( "hi %03u", 5); ...`, `... trice32 ( iD(0), "hi %03u", 5); ...`},
		{`... Trice32 ( "hi %03u", 5); ...`, `... Trice32 ( iD(0), "hi %03u", 5); ...`},
		{`... TRice32 ( "hi %03u", 5); ...`, `... TRice32 ( iD(0), "hi %03u", 5); ...`},
		//{`... TRICE32 ( "hi %03u", 5); ...`, `... TRICE32 ( Id(0), "hi %03u", 5); ...`},
		{`... trice64 ( "hi %03u", 5); ...`, `... trice64 ( iD(0), "hi %03u", 5); ...`},
		{`... Trice64 ( "hi %03u", 5); ...`, `... Trice64 ( iD(0), "hi %03u", 5); ...`},
		{`... TRice64 ( "hi %03u", 5); ...`, `... TRice64 ( iD(0), "hi %03u", 5); ...`},
		//{`... TRICE64 ( "hi %03u", 5); ...`, `... TRICE64 ( Id(0), "hi %03u", 5); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDTwoParamNA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d", 5, 7 ); ...`, `...   trice8_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...   Trice8 ( "hi %03u %03d", 5, 7 ); ...`, `...   Trice8_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...   TRice8 ( "hi %03u %03d", 5, 7 ); ...`, `...   TRice8_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d", 5, 7 ); ...`, `...   TRICE8_2 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  trice16 ( "hi %03u %03d", 5, 7 ); ...`, `...  trice16_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  Trice16 ( "hi %03u %03d", 5, 7 ); ...`, `...  Trice16_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  TRice16 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRice16_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRICE16_2 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  trice32 ( "hi %03u %03d", 5, 7 ); ...`, `...  trice32_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  Trice32 ( "hi %03u %03d", 5, 7 ); ...`, `...  Trice32_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  TRice32 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRice32_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRICE32_2 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  trice64 ( "hi %03u %03d", 5, 7 ); ...`, `...  trice64_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  Trice64 ( "hi %03u %03d", 5, 7 ); ...`, `...  Trice64_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  TRice64 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRice64_2 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRICE64_2 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDTwoParamA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d", 5, 7 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...   Trice8 ( "hi %03u %03d", 5, 7 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...   TRice8 ( "hi %03u %03d", 5, 7 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d", 5, 7 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  trice16 ( "hi %03u %03d", 5, 7 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  Trice16 ( "hi %03u %03d", 5, 7 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  TRice16 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  trice32 ( "hi %03u %03d", 5, 7 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  Trice32 ( "hi %03u %03d", 5, 7 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  TRice32 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  trice64 ( "hi %03u %03d", 5, 7 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  Trice64 ( "hi %03u %03d", 5, 7 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		{`...  TRice64 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d", 5, 7 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d", 5, 7 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d", 5, 7 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDThreeParamNA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   trice8_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   Trice8_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   TRice8_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   TRICE8_3 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  trice16_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  Trice16_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRice16_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRICE16_3 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  trice32_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  Trice32_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRice32_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRICE32_3 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  trice64_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  Trice64_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRice64_3 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRICE64_3 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDThreeParamA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b", 5, 7, 9 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d %16b", 5, 7, 9 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDFourParamNA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   trice8_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   Trice8_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   TRice8_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   TRICE8_4 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  trice16_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  Trice16_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRice16_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRICE16_4 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  trice32_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  Trice32_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRice32_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRICE32_4 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  trice64_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  Trice64_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRice64_4 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRICE64_4 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDFourParamA(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d %16b 0x%08x", 5, 7, 9, 3 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDFiveParamN(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   trice8_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   Trice8_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   TRice8_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   TRICE8_5 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  trice16_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  Trice16_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRice16_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRICE16_5 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  trice32_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  Trice32_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRice32_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRICE32_5 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  trice64_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  Trice64_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRice64_5 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRICE64_5 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDFiveParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d %16b 0x%08x %X", 5, 7, 9, 3, 2 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDSixParamN(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   trice8_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   Trice8_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   TRice8_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   TRICE8_6 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  trice16_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  Trice16_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRice16_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRICE16_6 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  trice32_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  Trice32_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRice32_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRICE32_6 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  trice64_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  Trice64_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRice64_6 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRICE64_6 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDSixParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d", 5, 7, 9, 3, 2, 4 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDSevenParamN(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   trice8_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   Trice8_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   TRice8_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   TRICE8_7 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  trice16_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  Trice16_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRice16_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRICE16_7 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  trice32_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  Trice32_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRice32_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRICE32_7 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  trice64_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  Trice64_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRice64_7 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRICE64_7 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDSevenParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u", 5, 7, 9, 3, 2, 4, 6 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDEightParamN(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   trice8_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   Trice8_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   TRice8_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   TRICE8_8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice16_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice16_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice16_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE16_8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice32_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice32_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice32_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE32_8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice64_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice64_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice64_8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE64_8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice_8   ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice_8   ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice_8   ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//  {`...  TRICE   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE_8   ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDEightParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{`...   trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...   Trice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   Trice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...   TRice8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   TRice8 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...   TRICE8 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...   TRICE8 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice16 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE16 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE16 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice32 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE32 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE32 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice64 ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE64 ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE64 ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  trice   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  trice   ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  Trice   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  Trice   ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		{`...  TRice   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRice   ( iD(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
		//{`...  TRICE   ( "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`, `...  TRICE   ( Id(0), "hi %03u %03d %16b 0x%08x %X %17d %99u %04b", 5, 7, 9, 3, 2, 4, 6, 8 ); ...`},
	}
	checkTestTable(t, tt, false)
}

func TestInsertParamCountAndIDAll0A(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	DefaultStampSize = 16
	StampSizeId = " Id(0),"

	tt := []struct{ text, exp string }{
		{
			`... TRICE0 ( "hi"); ...`,
			`... TRICE0 ( Id(0), "hi"); ...`},
		{
			`... TRICE0( "hi"); ...`,
			`... TRICE0( Id(0), "hi"); ...`},
		{
			`... TRICE8( "hi %d", 5); ...`,
			`... TRICE8_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE8( "hi %d, %u", 5, h); ...`,
			`... TRICE8_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... TRICE8( "hi %d, %u, %b", d, u, b); ...`,
			`... TRICE8_3( Id(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... TRICE8( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... TRICE8_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... TRICE8( "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h); ...`,
			`... TRICE8_8( Id(0), "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h); ...`},
		{
			`... TRICE8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h); ...`,
			`... TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h); ...`},
		{
			`... TRICE16( "hi %d", 5); ...`,
			`... TRICE16_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE16( "hi %d, %u", 5, h); ...`,
			`... TRICE16_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... TRICE16( "hi %d, %u, %b", d, u, b); ...`,
			`... TRICE16_3( Id(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... TRICE16( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... TRICE16_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... TRICE32( "hi %d", 5); ...`,
			`... TRICE32_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE32( "hi %d, %u", 5, h); ...`,
			`... TRICE32_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... TRICE32( "hi %d, %u, %b", d, u, b); ...`,
			`... TRICE32_3( Id(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... TRICE32( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... TRICE32_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... TRICE64( "hi %d", 5); ...`,
			`... TRICE64_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE64( "hi %d, %u", 5, h); ...`,
			`... TRICE64_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... trice0 ( "hi"); ...`,
			`... trice0 ( iD(0), "hi"); ...`},
		{
			`... trice0( "hi"); ...`,
			`... trice0( iD(0), "hi"); ...`},
		{
			`... trice8( "hi %d", 5); ...`,
			`... trice8_1( iD(0), "hi %d", 5); ...`},
		{
			`... trice8( "hi %d, %u", 5, h); ...`,
			`... trice8_2( iD(0), "hi %d, %u", 5, h); ...`},
		{
			`... trice8( "hi %d, %u, %b", d, u, b); ...`,
			`... trice8_3( iD(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... trice8( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... trice8_4( iD(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... trice8( "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h); ...`,
			`... trice8_8( iD(0), "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h); ...`},
		{
			`... trice8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h); ...`,
			`... trice8_8( iD(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h); ...`},
		{
			`... trice16( "hi %d", 5); ...`,
			`... trice16_1( iD(0), "hi %d", 5); ...`},
		{
			`... trice16( "hi %d, %u", 5, h); ...`,
			`... trice16_2( iD(0), "hi %d, %u", 5, h); ...`},
		{
			`... trice16( "hi %d, %u, %b", d, u, b); ...`,
			`... trice16_3( iD(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... trice16( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... trice16_4( iD(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... trice32( "hi %d", 5); ...`,
			`... trice32_1( iD(0), "hi %d", 5); ...`},
		{
			`... trice32( "hi %d, %u", 5, h); ...`,
			`... trice32_2( iD(0), "hi %d, %u", 5, h); ...`},
		{
			`... trice32( "hi %d, %u, %b", d, u, b); ...`,
			`... trice32_3( iD(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... trice32( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... trice32_4( iD(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... trice64( "hi %d", 5); ...`,
			`... trice64_1( iD(0), "hi %d", 5); ...`},
		{
			`... trice64( "hi %d, %u", 5, h); ...`,
			`... trice64_2( iD(0), "hi %d, %u", 5, h); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestInsertParamCountAndIDAll0B(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	StampSizeId = " Id(0),"
	tt := []struct{ text, exp string }{
		{
			`... TRICE0 ( "hi"); ...`,
			`... TRICE0 ( Id(0), "hi"); ...`},
		{
			`... TRICE0( "hi"); ...`,
			`... TRICE0( Id(0), "hi"); ...`},
		{
			`... TRICE8( "hi %d", 5); ...`,
			`... TRICE8_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE8( "hi %d, %u", 5, h); ...`,
			`... TRICE8_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... TRICE8( "hi %d, %u, %b", d, u, b); ...`,
			`... TRICE8_3( Id(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... TRICE8( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... TRICE8_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... TRICE8( "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h); ...`,
			`... TRICE8_8( Id(0), "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h); ...`},
		{
			`... TRICE8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h); ...`,
			`... TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h); ...`},
		{
			`... TRICE16( "hi %d", 5); ...`,
			`... TRICE16_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE16( "hi %d, %u", 5, h); ...`,
			`... TRICE16_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... TRICE16( "hi %d, %u, %b", d, u, b); ...`,
			`... TRICE16_3( Id(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... TRICE16( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... TRICE16_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... TRICE32( "hi %d", 5); ...`,
			`... TRICE32_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE32( "hi %d, %u", 5, h); ...`,
			`... TRICE32_2( Id(0), "hi %d, %u", 5, h); ...`},
		{
			`... TRICE32( "hi %d, %u, %b", d, u, b); ...`,
			`... TRICE32_3( Id(0), "hi %d, %u, %b", d, u, b); ...`},
		{
			`... TRICE32( "hi %d, %u, %b, %x", d, u, b, h); ...`,
			`... TRICE32_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h); ...`},
		{
			`... TRICE64( "hi %d", 5); ...`,
			`... TRICE64_1( Id(0), "hi %d", 5); ...`},
		{
			`... TRICE64( "hi %d, %u", 5, h); ...`,
			`... TRICE64_2( Id(0), "hi %d, %u", 5, h); ...`},
	}
	checkTestTable(t, tt, true)
}

func TestOptionallyExtendLenAndInsertID0B(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	StampSizeId = " Id(0),"
	tt := []struct{ text, exp string }{
		{
			`_TRICE8( "hi %d", 5); // to not touch`,
			`_TRICE8( "hi %d", 5); // to not touch`},
		{
			`TRICE8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
		{
			`TRICE8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
		{
			`TRICE8_3( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 ); // do not change wrong count`,
			`TRICE8_3( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 ); // do not change wrong count`},
	}
	checkTestTable(t, tt, true)
}

func checkTestTable(t *testing.T, tt []struct{ text, exp string }, extend bool) {
	for _, x := range tt {
		act, _ := updateParamCountAndID0(os.Stdout, x.text, extend)
		assert.Equal(t, x.exp, act)
	}
}

func TestOptionallyExtendLenAndInsertID0(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	StampSizeId = " Id(0),"
	tt := []struct{ text, exp string }{
		{
			`_TRICE8( "hi %d", 5); // to not touch`,
			`_TRICE8( "hi %d", 5); // to not touch`},
		{
			`TRICE8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
		{
			`TRICE8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
		{
			`TRICE8_3( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 ); // do not change wrong count`,
			`TRICE8_3( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 ); // do not change wrong count`},
		{
			`trice_s  ( "%s\n", rts ) \\ no semicolon`,
			`trice_s  ( iD(0), "%s\n", rts ) \\ no semicolon`},
		{
			`trice_s  ( "%s\n", "rts" );`,
			`trice_s  ( iD(0), "%s\n", "rts" );`},
	}
	checkTestTable(t, tt, true)
}

func TestVariadicInsertId0A(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	StampSizeId = " Id(0),"
	tt := []struct{ text, exp string }{
		{
			`Trice8( "hi %2d",1  );`,
			`Trice8( iD(0), "hi %2d",1  );`},
		{
			`TRICE8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`,
			`TRICE8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x",1,2,3,4,5,6,7,8 );`},
	}
	checkTestTable(t, tt, false)
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// todo

func _legacyUpdate_TestSingleParam(t *testing.T) {
	defer Setup(t)() // This executes Setup(t) and puts the returned function into the defer list.

	tt := []struct{ text, exp string }{
		{
			`Trice8_1( Id(0), "hi %2d",1  );`,
			`Trice8_1( iD(0), "hi %2d",1  );`}, // corrected ID letter case expected
	}
	checkTestTable(t, tt, true)
}

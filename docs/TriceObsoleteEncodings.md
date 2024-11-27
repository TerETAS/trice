# Obsolete *Trice* encodings

> _(Read only you are interested in)_
> This file exists only as reference for some thoughts for just in case an additional *Trice* encoding is considered. > > The **esc** and **flex** encoding worked well but the code is removed now in favor of the COBS encoding. Check out release [v0.32.0](https://github.com/rokath/trice/releases/tag/v0.32.0) or earlier for working code.

<details><summary>Table of Contents</summary><ol><!-- TABLE OF CONTENTS START -->

<!-- 
Table of Contents Generation:
- Install vsCode extension "Markdown TOC" from dumeng 
- Use Shift-Ctrl-P "markdownTOC:generate" to get the automatic numbering.
- replace "<a id=" with "<a id=" 
-->

<!-- vscode-markdown-toc -->
* 1. [General](#general)
* 2. [Quick start recommendation](#quick-start-recommendation)
* 3. [Overview](#overview)
  * 3.1. [`flex` encoding](#`flex`-encoding)
    * 3.1.1. [`flex` short sub-encoding](#`flex`-short-sub-encoding)
    * 3.1.2. [*`flex` medium sub-encoding*](#*`flex`-medium-sub-encoding*)
    * 3.1.3. [*`flex` long sub-encoding*](#*`flex`-long-sub-encoding*)
  * 3.2. [`pack2` & `pacl2L` encoding](#`pack2`-&-`pacl2l`-encoding)
  * 3.3. [Encoding `pack2` & `pack2L` (with cycle counter, 20-bit IDs, runtime strings up to 65535 chars)](#encoding-`pack2`-&-`pack2l`-(with-cycle-counter,-20-bit-ids,-runtime-strings-up-to-65535-chars))
  * 3.4. [`pack` & `packL` encoding](#`pack`-&-`packl`-encoding)
  * 3.5. [Encoding `pack` & `packL` (no cycle counter, 16-bit IDs, runtime strings up to 65535 chars)](#encoding-`pack`-&-`packl`-(no-cycle-counter,-16-bit-ids,-runtime-strings-up-to-65535-chars))
  * 3.6. [`bare` and `bareL` encoding](#`bare`-and-`barel`-encoding)
  * 3.7. [`wrap` and `wrapL` encoding](#`wrap`-and-`wrapl`-encoding)
  * 3.8. [`esc` encoding](#`esc`-encoding)
  * 3.9. [`mini` encoding (idea)](#`mini`-encoding-(idea))
  * 3.10. [`mix` encoding (idea)](#`mix`-encoding-(idea))
  * 3.11. [own encoding](#own-encoding)
* 4. [Encoding `bare` & `bareL`](#encoding-`bare`-&-`barel`)
* 5. [Encoding `wrap` & `wrapL`](#encoding-`wrap`-&-`wrapl`)
* 6. [Encoding `esc` (experimental)](#encoding-`esc`-(experimental))
  * 6.1. [Start byte `EC`](#start-byte-`ec`)
  * 6.2. [Length Code `LC`](#length-code-`lc`)
  * 6.3. [TriceID `IH` and `IL`](#triceid-`ih`-and-`il`)
  * 6.4. [Payload](#payload)
* 7. [Sync packages](#sync-packages)
* 8. [ COBS/R encoding](#-cobs/r-encoding)
* 9. [COBS/R encoding examples](#cobs/r-encoding-examples)
  * 9.1. [COBS/R encoding for 0-byte packages](#cobs/r-encoding-for-0-byte-packages)
  * 9.2. [COBS/R encoding for 1-byte packages](#cobs/r-encoding-for-1-byte-packages)
  * 9.3. [COBS/R encoding for 2-byte packages](#cobs/r-encoding-for-2-byte-packages)
  * 9.4. [COBS/R encoding for n-byte packages](#cobs/r-encoding-for-n-byte-packages)
  * 9.5. [COBS/R encoding](#cobs/r-encoding)
    * 9.5.1. [COBS/R encoding for n-byte packages](#cobs/r-encoding-for-n-byte-packages-1)
    * 9.5.2. [Decoded COBS/R package interpreter](#decoded-cobs/r-package-interpreter)
* 10. [Interpreter for decoded COBS/R package](#interpreter-for-decoded-cobs/r-package)
  * 10.1. [IMPORTANT](#important)
  * 10.2. [Encoding table 0 legend](#encoding-table-0-legend)
  * 10.3. [Encoding table 0 (without cycle counter)](#encoding-table-0-(without-cycle-counter))
  * 10.4. [Encoding table 1 legend](#encoding-table-1-legend)
  * 10.5. [Encoding table 1 (with 4-bit cycle counter)](#encoding-table-1-(with-4-bit-cycle-counter))
  * 10.6. [Encoding table 2 (with 8-bit cycle counter)](#encoding-table-2-(with-8-bit-cycle-counter))
* 11. [Fast TRICE data storing](#fast-trice-data-storing)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

<div id="top"></div></ol></details><!-- TABLE OF CONTENTS END -->

##  1. <a id='general'></a>General

Trice bytes can be encoded in different flavors and it is easy to develop a new encoding. The encoding `esc` is such a trial. It is not as good as the `flex` encoding but kept for reference.

##  2. <a id='quick-start-recommendation'></a>Quick start recommendation

- Use **flexL** encoding if your target processor is a little endian machine, otherwise use **flex**
- The `trice` tool assumes **flexL** per default, so no need for commandline switch `-enc flexL`.

```c
#define TRICE_ENCODING TRICE_FLEX_ENCODING
```

##  3. <a id='overview'></a>Overview

Inside the target project specific triceConfig.h is selectable:

- Encoding in little or big endian.
  - The encoding should match the target processor endiannes.

  ```c
  #define TRICE_TRANSFER_ENDIANNESS TRICE_LITTLE_ENDIANNESS
  ```

- Additionally an encoding can be wrapped with transport information. This setting is done by calling cyclically this function:

```c
triceServeFifoWrappedToBytesBuffer();
```

 As example checkout wrapped bare encoding.

- Also it is possible to use encryption, which is shown as example for the wrapped bare encoding.

Currently these encodings are supported:

###  3.1. <a id='`flex`-encoding'></a>`flex` encoding

The 3 formats **short**, **medium** and **long** are usable parallel. String transfer is done not in the short format.
A format extension is possible by using the 2 reserved patterns in medium format.
In fact the ID contains information how the payload bytes are structured. This is in the first place
the format string, but could also be some (packed) structs and the ID refers to a function taking the payload as parameters.

- `I` = **I**D bit
- `D` = **D**ata bit
- `N` = **N**umber of bytes bit
- `C` = **C**ycle counter bit

- The bit 31 is the mode bit.
  - 0: short encoding
  - 1: medium and long encoding

####  3.1.1. <a id='`flex`-short-sub-encoding'></a>`flex` short sub-encoding

- Maximum payload 2 bytes

This sub-encodig is mainly for _very_ small systems and time critical stuff

- `0IIIIIII IIIIIIII DDDDDDDD DDDDDDDD` : short, implicit count
- An "implicit" count means coded in trice format resulting from trice-ID and therefore no check option.
- Short encoding should be used for time and space critical cases.
  - Configuration is in triceConfig.h possible

```b
0IIIIIII IIIIIIII 00000000 00000000 : short, implicit count=0, Trice0
0IIIIIII IIIIIIII DDDDDDDD DDDDDDDD : short, implicit count=1, Trice8_1, Trice16_1
0IIIIIII IIIIIIII DDDDDDDD DDDDDDDD : short, implicit count=2, Trice8_2 
```

####  3.1.2. <a id='*`flex`-medium-sub-encoding*'></a>*`flex` medium sub-encoding*

- Maximun payload 4 bytes

- `1IIIIIII IIIIIIII IIIIINNN CCCCCCCC` : medium 3-bit count NNN
- The medium 3-bit counts 5 and 6 are reserved for future extensions.

```b
1IIIIIII IIIIIIII IIIII000 CCCCCCCC : medium, count=0, TRICE0

1IIIIIII IIIIIIII IIIII001 CCCCCCCC : medium, count=1, TRICE8_1
DDDDDDDD 00000000 00000000 00000000

1IIIIIII IIIIIIII IIIII010 CCCCCCCC : medium, count=2, TRICE8_2
DDDDDDDD DDDDDDDD 00000000 00000000

1IIIIIII IIIIIIII IIIII011 CCCCCCCC : medium, count=3, TRICE8_3
DDDDDDDD DDDDDDDD DDDDDDDD 00000000

1IIIIIII IIIIIIII IIIII100 CCCCCCCC : medium, count=4, TRICE8_4, TRICE16_2, TRICE32_1, TRICE16_1, TRICE8_1
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

1IIIIIII IIIIIIII IIIII101 CCCCCCCC : reserved
...

1IIIIIII IIIIIIII IIIII110 CCCCCCCC : reserved
...
```

####  3.1.3. <a id='*`flex`-long-sub-encoding*'></a>*`flex` long sub-encoding*

- Maximun payload 65535 bytes

- `1IIIIIII IIIIIIII IIIII111 CCCCCCCC` : long 16-bit count with 16 bit eXor checksum follows
- `NNNNNNNN NNNNNNNN cccccccc cccccccc` : 16-bit count NNNNNNNN NNNNNNNN==^cccccccc cccccccc
- `DDDDDDDD ...`

```b
1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long, TRICE8_5
00000000 00000101 11111111 11111010 = count 5
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD 00000000 00000000 00000000

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long, TRICE 8_6, TRICE 16_3
00000000 00000110 11111111 11111001 = count 6
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD 00000000 00000000

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long, TRICE 8_7
00000000 00000111 11111111 11111000 = count 7
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD 00000000

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long, TRICE8_8 TRICE16_4, TRICE32_2, TRICE64_1
00000000 00001000 11111111 11110111 = count 8
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long, strlen 9
00000000 00001001 11111111 11110110 = count 9
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD 00000000 00000000 00000000

...

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long TRICE32_3,
00000000 00001001 11111111 11110110 = count 12
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

...

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long TRICE32_4, TRICE64_2
00000000 00001001 11111111 11110110 = count 16
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

...

1IIIIIII IIIIIIII IIIII111 CCCCCCCC = long, strlen 65535
11111111 11111111 00000000 00000000 = count 65535
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
...
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD 00000000

```

###  3.2. <a id='`pack2`-&-`pacl2l`-encoding'></a>`pack2` & `pacl2L` encoding

This is the recommended encoding.

- with cycle counter
- 20-bit IDs
- runtime strings up to 65535 chars
- target source: trice/srcTrice.C/intern/tricePack2Encoder.h
- trice tool source: trice/internal/decoder/pack2Decoder.go
- trice tool test file: trice/internal/decoder/pack2Decoder_test.go

###  3.3. <a id='encoding-`pack2`-&-`pack2l`-(with-cycle-counter,-20-bit-ids,-runtime-strings-up-to-65535-chars)'></a>Encoding `pack2` & `pack2L` (with cycle counter, 20-bit IDs, runtime strings up to 65535 chars)

The encoding is similar to `pack` & `packL` encoding with these differences:

`IIIICCCC` is replaced by `IIIIICNN` and  an optional following long count `LLLLcccc`

- `IIIII` = 20-bit ID
- `C` = 4-bit byte count
  - 0...12 = short count (no following long count)
  - 0xd = indicates a following long count
  - 0xe = reserved
  - 0xf = reserved
- `NN` = 8-bit cycle counter
- `LLLL` = 16-bit long count
- `cccc` = bit-inversed LLLL as check sum

Bit pattern:

```b
0IIIIIII IIIIIIII DDDDDDDD DDDDDDDD = mini
1IIIIIII IIIIIIII IIIINNNN CCCCCCCC = pack2

1IIIIIII IIIIIIII IIII0000 CCCCCCCC = pack2 TRICE0

1IIIIIII IIIIIIII IIII0001 CCCCCCCC = pack2 TRICE8_1
DDDDDDDD 00000000 00000000 00000000

1IIIIIII IIIIIIII IIII0010 CCCCCCCC = pack2 TRICE8_2, TRICE16_1
DDDDDDDD DDDDDDDD 00000000 00000000

1IIIIIII IIIIIIII IIII0011 CCCCCCCC = pack2 TRICE8_3
DDDDDDDD DDDDDDDD DDDDDDDD 00000000

1IIIIIII IIIIIIII IIII0100 CCCCCCCC = pack2 TRICE8_4, TRICE16_2, TRICE32_1
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

1IIIIIII IIIIIIII IIII0101 CCCCCCCC = pack2 TRICE8_5
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD 00000000 00000000 00000000

1IIIIIII IIIIIIII IIII0110 CCCCCCCC = pack2 TRICE8_6 TRICE16_3
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD 00000000 00000000

1IIIIIII IIIIIIII IIII0111 CCCCCCCC = pack2 TRICE8_7
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD 00000000

1IIIIIII IIIIIIII IIII1000 CCCCCCCC = pack2 TRICE8_8 TRICE16_4, TRICE32_2, TRICE64_1
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

...

1IIIIIII IIIIIIII IIII1100 CCCCCCCC = pack2 TRICE16_3
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD
DDDDDDDD DDDDDDDD DDDDDDDD DDDDDDDD

1IIIIIII IIIIIIII IIII1101 CCCCCCCC = pack2 long count
1IIIIIII IIIIIIII IIII1110 CCCCCCCC = pack2 reserved
1IIIIIII IIIIIIII IIII1111 CCCCCCCC = pack2 reserved

1IIIIIII IIIIIIII IIII1101 CCCCCCCC = pack2 long count
NNNNNNNN NNNNNNNN nnnnnnnn nnnnnnnn = 16-bit count N and bit invers n
```

###  3.4. <a id='`pack`-&-`packl`-encoding'></a>`pack` & `packL` encoding

This is the pack2 & pack2L predecessor and kept for reference.

- no cycle counter
- 16-bit IDs
- runtime strings up to 65535 chars
- target source: trice/srcTrice.C/intern/tricePackEncoder.h
- trice tool source: trice/internal/decoder/packDecoder.go
- trice tool test file: trice/internal/decoder/packDecoder_test.go

###  3.5. <a id='encoding-`pack`-&-`packl`-(no-cycle-counter,-16-bit-ids,-runtime-strings-up-to-65535-chars)'></a>Encoding `pack` & `packL` (no cycle counter, 16-bit IDs, runtime strings up to 65535 chars)

All values up to 32 bit are combined 32 bit units in big (=network) or little endian order.
64-bit values are in the same byte order.

- `IIII` = 16-bit ID
- `CCCC` = byte count without counting padding bytes
- 0-3 padding 0-bytes fill the last 32-bit unit

```b
byte 3 2 1 0  | macro
--------------|-----------------------------------------------------------------
   0x89abcdef | TRICE16_1( Id(0x89ab), "inf:[ SYNCTRICE 0x89ab%04x ]", 0xcdef );
     IIIICCCC | TRICE0( Id(I), "..." );
     IIIICCCC | TRICE_S( Id(I), "...%s...", "" );
```

```b
byte 3 2 1 0    3 2 1 0      | macro
-----------------------------|-----------------------------------------------------
     IIIICCCC   000000b0     | TRICE8_1( Id(I), "...", b0 );             // cnt = 1
     IIIICCCC   0000b0b1     | TRICE8_2( Id(I), "...", b0, b1 );         // cnt = 2
     IIIICCCC   00b0b1b2     | TRICE8_3( Id(I), "...", b0, b1, b2 );     // cnt = 2
     IIIICCCC   b0b1b2b3     | TRICE8_4( Id(I), "...", b0, b1, b2, b3 ); // cnt = 3
     IIIICCCC   0000w0w0     | TRICE16_1( Id(I), "...", w0 );            // cnt = 2
     IIIICCCC   w0w0w1w1     | TRICE16_2( Id(I), "...", w0, w1 );        // cnt = 4
     IIIICCCC   d0d0d0d0     | TRICE32_1( Id(I), "...", d0 );            // cnt = 4
     IIIICCCC   000000aa     | TRICE_S( Id(I), "...%s...", "a" );        // cnt = 1
     IIIICCCC   0000aabb     | TRICE_S( Id(I), "...%s...", "ab" );       // cnt = 2
     IIIICCCC   00aabbcc     | TRICE_S( Id(I), "...%s...", "abc" );      // cnt = 3
     IIIICCCC   aabbccdd     | TRICE_S( Id(I), "...%s...", "abcd" );     // cnt = 4
```

```b
byte  3 2 1 0   3 2 1 0   3 2 1 0  | macro
-----------------------------------|---------------------------------------------------------------------
     IIIICCCC  b0b1b2b3  000000b4  | TRICE8_5( Id(I), "...", b0, b1, b2, b3, b4 );             // cnt = 5
     IIIICCCC  b0b1b2b3  0000b4b5  | TRICE8_6( Id(I), "...", b0, b1, b2, b3, b4, b5 );         // cnt = 6
     IIIICCCC  b0b1b2b3  00b4b5b6  | TRICE8_7( Id(I), "...", b0, b1, b2, b3, b4, b5, b6 );     // cnt = 7
     IIIICCCC  b0b1b2b3  b4b5b6b7  | TRICE8_8( Id(I), "...", b0, b1, b2, b3, b4, b5, b6, b7 ); // cnt = 8
     IIIICCCC  w0w0w1w1  0000w2w2  | TRICE16_3( Id(I), "...", w0, w1, w2 );                    // cnt = 6
     IIIICCCC  w0w0w1w1  w2w2w3w3  | TRICE16_4( Id(I), "...", w0, w1, w2, w3 );                // cnt = 8
     IIIICCCC  d0d0d0d0  d1d1d1d1  | TRICE32_2( Id(I), "...", d0, d1 );                        // cnt = 8
     IIIICCCC  l0l0l0l0  l0l0l0l0  | TRICE64_1( Id(I), "...", l0 );                            // cnt = 8
     IIIICCCC  aabbccdd  000000ee  | TRICE_S( Id(I), "...%s...", "abcde" );                    // cnt = 5
     IIIICCCC  aabbccdd  0000eeff  | TRICE_S( Id(I), "...%s...", "abcdef" );                   // cnt = 6
     IIIICCCC  aabbccdd  00eeffgg  | TRICE_S( Id(I), "...%s...", "abcdefg" );                  // cnt = 7
     IIIICCCC  aabbccdd  eeffgghh  | TRICE_S( Id(I), "...%s...", "abcdefgh" );                 // cnt = 8
```

and so on...

A [sync package](#sync-packages) can be inserted anytime between 2 trice but not inside a trice.

###  3.6. <a id='`bare`-and-`barel`-encoding'></a>`bare` and `bareL` encoding

This was the first minimal implementation. Could be interstuing for 8-bit and 16-bit processors.

- no cycle counter
- 16-bit IDs
- no runtime strings
- target source: trice/srcTrice.C/intern/triceBareEncoder.h
- trice tool source: trice/internal/decoder/bareDecoder.go
- trice tool test file: trice/internal/decoder/bareDecoder_test.go

###  3.7. <a id='`wrap`-and-`wrapl`-encoding'></a>`wrap` and `wrapL` encoding

This is also `bare` & `bareL` encoding but with additional control bytes.

###  3.8. <a id='`esc`-encoding'></a>`esc` encoding

This is a try-out escape sequence encoding implementation and kept for reference.

- no endianness choice
- no cycle counter
- 16-bit IDs
- runtime strings up to 255 chars
- target source: trice/srcTrice.C/intern/triceEscEncoder.h
- trice tool source: trice/internal/decoder/escDecoder.go
- trice tool test file: trice/internal/decoder/escDecoder_test.go

###  3.9. <a id='`mini`-encoding-(idea)'></a>`mini` encoding (idea)

- Supports only `TRICE0`, `TRICE16_1`, `TRICE8_2`.
- 16-bit IDs and 16-bit data
- Recommended for 8-bit controller
- Minimal memory footprint
- Maximal speed.

###  3.10. <a id='`mix`-encoding-(idea)'></a>`mix` encoding (idea)

- Combines `pack2` and `mini` by using 1 bit from the ID space as mode bit.
- If mode bit is 0, then `mini` encoding with 15-bit IDs (32767 usable)
- If mode bit is 1, then `pack2` encoding with 19-bit IDs (524287 usable)
- The 4-bit count offers 2 reserved values for future extension.

###  3.11. <a id='own-encoding'></a>own encoding

To implement a different encoding:

- Copy trice/srcTrice.C/intern/trice*Any*Encoder.h, to trice/srcTrice.C/intern/trice*Own*Encoder.h.
- Adapt trice/srcTrice.C/intern/trice*Own*Encoder.h and integrate it in trice.h accordingly.
- Create a test project, copy and adapt triceConfig.h in the desired way.
- Copy trice/internal/decoder/*any*Decoder.go to trice/internal/decoder/**own**Decoder.go and adapt it.
- Integrate *own*Decoder.go accordingly.
- Write tests!

##  4. <a id='encoding-`bare`-&-`barel`'></a>Encoding `bare` & `bareL`

- Each trice is coded in one to eight 4-byte trice atoms.
- A trice atom consists of a 2 byte id and 2 bytes data.
- When a trice consists of several trice atoms, only the last one carries the trice id. The others have a trice id 0.

- `IIII` = 16-bit ID

```b
byte  3 2 1 0  | macro
---------------|--------------------------------
     IIII0000  | TRICE0( Id(I), "..." );
     IIII00b0  | TRICE8_1( Id(I), "...", b0 );
     IIIIb0b1  | TRICE8_2( Id(I), "...", b1 );
     IIIIw0w0  | TRICE16_1( Id(I), "...", w0 );
```

```b
byte  3 2 1 0    3 2 1 0  | macro
--------------------------|-----------------------------------------------
     0000b0b1   IIII00b2  | TRICE8_3( Id(I), "...", b0, b1, b2 );
     0000b0b1   IIIIb2b3  | TRICE8_4( Id(I), "...", b0, b1, b2, b3 );
     0000w0w0   IIIIw1w1  | TRICE16_2( Id(I), "...", w0, w1 );
     0000d0d0   IIIId0d0  | TRICE32_1( Id(I), "...", d0 );
```

```b
byte  3 2 1 0    3 2 1 0    3 2 1 0  | macro
-----------------------------------  |------------
     0000b0b1   0000b2b3   IIII00b4  | TRICE8_5( Id(I), "...", b0, b1, b2, b3, b4 );
     0000b0b1   0000b2b3   IIIIb4b5  | TRICE8_6( Id(I), "...", b0, b1, b2, b3, b4, b5 );
     0000w0w0   0000w1w1   IIIIw2w2  | TRICE16_3( Id(I), "...", w0, w1, w2);
```

```b
byte 3 2 1 0    3 2 1 0    3 2 1 0    3 2 1 0  | macro
-----------------------------------------------|------------
    0000b0b1   0000b2b3   0000b4b5   IIII00b6  | TRICE8_7( Id(I), "...", b0, b1, b2, b3, b4, b5, b6 );
    0000b0b1   0000b2b3   0000b4b5   IIIIb6b7  | TRICE8_8( Id(I), "...", b0, b1, b2, b3, b4, b5, b6, b7 );
    0000w0w0   0000w1w1   0000w2w2   IIIIw3w3  | TRICE16_3( Id(I), "...", w0, w1, w2, w3);
    0000d0d0   0000d0d0   0000d1d1   IIIId1d1  | TRICE32_2( Id(I), "...", d0, d1 );
    0000l0l0   0000l0l0   0000l0l0   IIIIl0l0  | TRICE64_1( Id(I), "...", l0 );
```

and so on...

The bare transmit format is exactly the same as the bare internal storage format with these differences:

During runtime normally only the 16-bit ID 12345 (together with the parameters like hour, min, sec) is copied to a buffer.
Execution time for a TRICE16_1 (as example) on a 48 MHz ARM can be about 16 systicks resulting in 250 nanoseconds duration, so you can use `trice` also inside interrupts or the RTOS scheduler to analyze task timings.
The needed buffer space is one 32 bit word per normal trice (for up to 2 data bytes).

If the wrap format is desired as output the buffered 4 byte trice is transmitted as an 8 byte packet allowing start byte, sender and receiver addresses and CRC8 check to be used later in parallel with different software protocols.

The bare output format contains exactly the bare bytes but is enriched with 4 byte [sync packages](#sync-packages) mixed in at 4 byte offsets to achieve syncing. The sync package interval is adjustable.

##  5. <a id='encoding-`wrap`-&-`wrapl`'></a>Encoding `wrap` & `wrapL`

This is the same as bare, but each trice atom is prefixed with a 4 byte wrap information:

- 0xEB = start byte
- 0x80 = source address
- 0x81 = destination address
- crc8 = 8 bit checksum over start byte, source and destination address, and the 4 bare bytes.
--->

##  6. <a id='encoding-`esc`-(experimental)'></a>Encoding `esc` (experimental)

The `esc` encoding uses an escape character for syncing after some data loss. It is extendable.

- All numbers are transmitted in network order (big endian).
- All values are in left-right order - first value comes first.

An `esc` trice transfer packet consists of an 4-byte header followed by an optional payload.

|    Start Byte    | Second Byte      | Third Byte   | Fourth Byte  |
|:----------------:|------------------|--------------|--------------|
| Escape char `EC` | Length Code `LC` | triceID `IH` | triceID `IL` |

###  6.1. <a id='start-byte-`ec`'></a>Start byte `EC`

```c
#define TRICE_ESC  0xEC //!< Escape char is control char to start a package.
```

The ESCaped encoding allowes a syncing to the next trice message on the behalf of the escape character `0xec`. It is
always followed by a length code in the range `0xdf ... 0xe8`. All other bytes are reserved for future usage despite
the `0xde` byte.

```c
#define TRICE_DEL  0xDE //!< Delete char, if follower of TRICE_ESC, deletes the meaning os TRICE_ESC making it an ordinary TRICE_ESC char.
```

This is inserted as not counted value into the bytes stream after an `0xec` to signal that this is an ordinary `0xec`
byte inside the data stream. As byte `0xec` is not used so often is is defined as ESC character:

###  6.2. <a id='length-code-`lc`'></a>Length Code `LC`

The LC is a 1-byte logarithmic length code. This is a copy
from [trice.h lines 44-58](https://github.com/rokath/trice/blob/master/srcTrice.C/trice.h) and shows the length code
meaning:

```c
#define TRICE_P0   0xDF //!< No param char = If follower of TRICE_ESC only a 16 bit ID is inside the payload.
#define TRICE_P1   0xE0 //!< 1 byte param char = If follower of TRICE_ESC a 16 bit ID and 1 byte are inside the payload.
#define TRICE_P2   0xE1 //!< 2 byte param char = If follower of TRICE_ESC a 16 bit ID and 2 byte are inside the payload.
#define TRICE_P4   0xE2 //!< 4 byte param char = If follower of TRICE_ESC a 16 bit ID and 4 byte are inside the payload.
#define TRICE_P8   0xE3 //!< 8 byte param char = If follower of TRICE_ESC a 16 bit ID and 8 byte are inside the payload.
#define TRICE_P16  0xE4 //!< 16 byte param char = If follower of TRICE_ESC a 16 bit ID and 8 byte are inside the payload.
//                 0xE5 // dynamically used for runtime strings with size 17-32
//                 0xE6 // dynamically used for runtime strings with size 33-64
//                 0xE7 // dynamically used for runtime strings with size 63-128
//                 0xE8 // dynamically used for runtime strings with size 127-256
//                 0xE9 // reserved
//                 0xEA // reserved
//                 0xEB // reserved

```

###  6.3. <a id='triceid-`ih`-and-`il`'></a>TriceID `IH` and `IL`

- The third and fourth byte are the 16 bit trice ID: IH & IL.
- The trice ID encodes one of the allowed trice macros, and a format string.
- The format string has some format specifiers accordingly to the trice macro.
- In the case of `TRICE_S` the format string contains one and only one `%s`.

###  6.4. <a id='payload'></a>Payload

A number of bytes according LC is optionally following the header. If within the data to be transmitted an 0xEC occurs
it stays on its place and is followed by a not counted 0xDE byte to signal that this is no start byte.

- Generic description

```b
Code                  |Meaning                   |pad|Remark
----------------------|--------------------------|---|-----------------------------------------------------------------------------
EC LC IH IL ...       |payload = 2^(LC-E0) bytes |   |LC is a length code
EC 00 ...             |reserved                  |   |All packages starting with EC 00 until starting with EC DD are reserved.
EC .. ...             |reserved                  |   |All packages starting with EC 00 until starting with EC DD are reserved.
EC DD ...             |reserved                  |   |All packages starting with EC 00 until starting with EC DD are reserved.
EC DE = EC            |real EC character         |   |If inside the payload occures EC an uncounted DE is injected afterwards.
EC DF IH IL           |16 bit ID no payload      |  0|TRICE0, special case: 2^-1 = 0 byte payload
EC E0 IH IL B0        |16 bit ID   1 byte payload|  0|TRICE8_1, TRICE_S(""): 2^0 = 1 byte payload
EC E1 IH IL B0 B1     |16 bit ID   2 byte payload|  0|TRICE8_2, TRICE16_1, TRICE_S("0")
EC E2 IH IL B0 .. B3  |16 bit ID   4 byte payload|  1|TRICE8_3, TRICE8_4, TRICE16_2, TRICE32_1, TRICE_S("01"), TRICE_S("012")
EC E3 IH IL B0 .. B7  |16 bit ID   8 byte payload|  3|TRICE8_5,...TRICE8_8,TRICE16_3,TRICE16_4,TRICE32_2,TRICE64_1,TRICE_S("0...7")
EC E4 IH IL B0 .. B15 |16 bit ID  16 byte payload|  7|TRICE32_3, TRICE32_4, TRICE64_2, TRICE_S("0...e")
EC E5 IH IL B0 .. B31 |16 bit ID  32 byte payload| 15|TRICE_S("0123456789abcdef"), ... TRICE_S(strlen(31))
EC E6 IH IL B0 .. B63 |16 bit ID  64 byte payload| 31|TRICE_S(strlen(32)), ...,TRICE_S(strlen(63))
EC E7 IH IL B0 .. B127|16 bit ID 128 byte payload| 63|TRICE_S(strlen(64)), ...,TRICE_S(strlen(127))
EC E8 IH IL B0 .. B255|16 bit ID 256 byte payload|127|TRICE_S(strlen(128)), ...,TRICE_S(strlen(255))
EC E9 ...             |reserved                  |   |All packages starting with EC E9 until starting with EC FF are reserved.
EC .. ...             |reserved                  |   |All packages starting with EC E9 until starting with EC FF are reserved.
EC FF ...             |reserved                  |   |All packages starting with EC E9 until starting with EC FF are reserved.
```

- Examples See function `TestEsc` and `TestEscDynStrings` in
  file [decoder_test.go](https://github.com/rokath/trice/blob/master/internal/decoder/decoder_test.go).

##  7. <a id='sync-packages'></a>Sync packages

- The frequency is adjustable and could be every 100ms or 40 bytes.
- The PC `trice` tool removes them silently.
- A sync package is `0x89 0xab 0xcd 0ef`.
- The PC `trice` tool does not use

```c
//! TRICE_SYNC is an optional trice sync message for syncing, when bare transmission is used.
//! The value 35243 (0x89ab) is a reserved pattern used as ID with value DA 0xcdef.
//! The hex notation protects against unwanted automatic changes.
//! The byte sequence of the sync message is 0x89 0xab 0xcd 0xef.
//! It cannot occure in the trice stream in another way due to ID generaion policy.
//! Sync package is IDDA=89abcdef
//!
//! To avoid wrong syncing these ID's are excluded: xx89, abcd, cdef, efxx (514 pieces)
//!
//! Possible:    IH IL DH DL IH IL DH DL IH IL DH DL (1 right)
//!              xx xx xx xx xx 89 ab cd ef xx xx xx -> avoid IL=89, IH=ef
//!
//! Possible:    IH IL DH DL IH IL DH DL IH IL DH DL (2 right)
//!              xx xx xx xx xx xx 89 ab cd ef xx xx -> avoid ID=cdef
//!
//! Possible:    IH IL DH DL IH IL DH DL IH IL DH DL (3 right)
//!              xx xx xx xx xx xx xx 89 ab cd ef xx -> avoid ID=abcd
//!
//! Sync packet: IH IL DH DL IH IL DH DL IH IL DH DL
//!              xx xx xx xx 89 ab cd ef xx xx xx xx -> use ID=89ab with DA=cdef as sync packet
//!
//!  Possible:   IH IL DH DL IH IL DH DL IH IL DH DL (1 left)
//!              xx xx xx 89 ab cd ef xx xx xx xx xx -> avoid ID=abcd
//!
//!  Possible:   IH IL DH DL IH IL DH DL IH IL DH DL (2 left)
//!              xx xx 89 ab cd ef xx xx xx xx xx xx -> avoid ID=cdef
//!
//!  Possible:   IH IL DH DL IH IL DH DL IH IL DH DL (3 left)
//!              xx 89 ab cd ef xx xx xx xx xx xx xx ->  avoid IL=nn89, IH=ef
//!
//! If an ID=89ab with DA!=cdef is detected -> out of sync!
//! If an IH=ef is detected -> out of sync, all 256 IDs starting with 0xef are excluded
//! If an IL=89 is detected -> out of sync, all 256 IDs ending with 0x89 are excluded
//! If an ID=abcd is detected -> out of sync, ID 0xabcd is excluded
//! If an ID=cdef is detected -> out of sync, ID 0xcdef is excluded
//! ID 0x89ab is reserved for this trice sync package.
//! The trice sync message payload must be 0xcdef.
//! You must not change any of the above demands. Otherwise the syncing will not work.
//! The Id(0x89ab) is here as hex value, so it is ignored by ID management.
//! The trice sync string makes the trice sync info invisible just in case,
//! but the trice tool will filter them out anyway. The trice tool automatic id generation
//! follows these rules.
//#define TRICE_SYNC do{ TRICE16_1( Id(0x89ab), "%x\b\b\b\b", 0xcdef ); }while(0)
```

##  8. <a id='-cobs/r-encoding'></a> COBS/R encoding

- Packages are [COBS/R](https://pythonhosted.org/cobs/cobsr-intro.html) encoded.
- Selected separator byte is `00`. That means the COBS/R encoded packages contain no `00` bytes and separated by a `00` byte.
- After a transfer interruption a very easy resync mechanism is usable: simply wait for the next `00` byte.
- The COBS/R encoding usually has the same length as the unencoded data and sometimes has one byte more but an additional `00` is needed for package separation.
- This way the ID bit count is adjustable to the real communication needs.
- One important point is the possibility to embed additional protocols in the data stream.


##  9. <a id='cobs/r-encoding-examples'></a>COBS/R encoding examples

###  9.1. <a id='cobs/r-encoding-for-0-byte-packages'></a>COBS/R encoding for 0-byte packages

- This is simply an empty package. Just the `00` package separator byte is transmitted.
- It is normally used as padding byte to reach a multiple of 8 bytes package length when putting several COBS/R packages into one encryption packet.

###  9.2. <a id='cobs/r-encoding-for-1-byte-packages'></a>COBS/R encoding for 1-byte packages

- One byte COBS/R packages are a 1:1 transformation despite for the values `00` and `01`.

| raw  | COBS/R (all followed by a not shown 00)  | remark
| :--  | :-----                                   | ---------------------------
| `00` |  `01 01`                                 | starting byte 00 prolongs code
| `01` |  `02 01`                                 | starting byte 01 prolongs code
| `02` |  `02`                                    |
| `03` |  `03`                                    |
| `...`|  `...`                                   |
| `fc` |  `fc`                                    |
| `fd` |  `fd`                                    |
| `fe` |  `fe`                                    |
| `ff` |  `ff`                                    |

One byte packages are fast COBS/R codable by simply incrementing the 2 values `00` and `01` and appending a `01`.

###  9.3. <a id='cobs/r-encoding-for-2-byte-packages'></a>COBS/R encoding for 2-byte packages

- Two bytes COBS/R packages are often a 1:1 transformation despite some cases as seen in the following table.

| raw  | COBS/R (all followed by a not shown 00)     | remark
| :--  | :-----                                      | ---------------------------
| `00 00` |  `01 01 01`                              | starting bytes 00, 01 and 02 prolong code usually
| `00 01` |  `02 01 01`                              |
| `00 02` |  `02 02 01`                              |
| `00 03` |  `02 03 01`                              |
| `...`   |  `...`                                   |
| `00 fc` |  `02 fc 01`                              |
| `00 fd` |  `02 fd 01`                              |
| `00 fe` |  `02 fe 01`                              |
| `00 ff` |  `02 ff 01`                              |
| `...`   |  `...`                                   |
| `01 00` |  `01 02 01`                              |
| `01 01` |  `03 01 01`                              |
| `01 02` |  `03 02 01`                              |
| `01 03` |  `03 03 01`                              |
| `...`   |  `...`                                   |
| `01 fc` |  `03 fc 01`                              |
| `01 fd` |  `03 fd 01`                              |
| `01 fe` |  `03 fe 01`                              |
| `01 ff` |  `03 ff 01`                              |
| `...`   |  `...`                                   |
| `02 00` |  `01 02   `                              | special case
| `02 01` |  `03 01 02`                              |
| `02 02` |  `03 02 02`                              |
| `02 03` |  `03 03 02`                              |
| `...`   |  `...`                                   |
| `02 fc` |  `03 fc 02`                              |
| `02 fd` |  `03 fd 02`                              |
| `02 fe` |  `03 fe 02`                              |
| `02 ff` |  `03 ff 02`                              |
| `...`   |  `...`                                   |
| `03 00` |  `01 03`                                 |
| `03 01` |  `03 01`                                 |
| `03 02` |  `03 02`                                 |
| `03 03` |  `03 03`                                 |
| `...`   |  `...`                                   |
| `03 fc` |  `03 fc`                                 |
| `03 fd` |  `03 fd`                                 |
| `03 fe` |  `03 fe`                                 |
| `03 ff` |  `03 ff`                                 |
| `...`   |  `...`                                   |
| `fc 00` |  `01 fc`                                 |
| `fc 01` |  `fc 01`                                 |
| `fc 02` |  `fc 02`                                 |
| `fc 03` |  `fc 03`                                 |
| `...`   |  `...`                                   |
| `fc fc` |  `fc fc`                                 |
| `fc fd` |  `fc fd`                                 |
| `fc fe` |  `fc fe`                                 |
| `fc ff` |  `fc ff`                                 |
| `...`   |  `...`                                   |
| `fd 00` |  `01 fd`                                 |
| `fd 01` |  `fd 01`                                 |
| `fd 02` |  `fd 02`                                 |
| `fd 03` |  `fd 03`                                 |
| `...`   |  `...`                                   |
| `fd fc` |  `fd fc`                                 |
| `fd fd` |  `fd fd`                                 |
| `fd fe` |  `fd fe`                                 |
| `fd ff` |  `fd ff`                                 |
| `...`   |  `...`                                   |
| `fe 00` |  `01 fe`                                 |
| `fe 01` |  `fe 01`                                 |
| `fe 02` |  `fe 02`                                 |
| `fe 03` |  `fe 03`                                 |
| `...`   |  `...`                                   |
| `fe fc` |  `fe fc`                                 |
| `fe fd` |  `fe fd`                                 |
| `fe fe` |  `fe fe`                                 |
| `fe ff` |  `fe ff`                                 |
| `...`   |  `...`                                   |
| `ff 00` |  `01 ff`                                 |
| `ff 01` |  `ff 01`                                 |
| `ff 02` |  `ff 02`                                 |
| `ff 03` |  `ff 03`                                 |
| `...`   |  `...`                                   |
| `ff fc` |  `ff fc`                                 |
| `ff fd` |  `ff fd`                                 |
| `ff fe` |  `ff fe`                                 |
| `ff ff` |  `ff ff`                                 |

- Two byte packages are fast COBS/R codable by simply using an Id subset having no first byte 0, 1, 2 and no 0 in the 2nd byte higher nibble:
- Using `II IC`, where C is a 4 bit cycle counter assumed to sometimes 0 :
  - Id = 0xIII0 = range 4096 
  - Id = 0x00n = range 16 is forbidden
  - Id = 0x01n = range 16 is forbidden // only for trice0
  - Id = 0x02n = range 16 is forbidden // only for trice0
  - Id = 0xnn0 = range 256 is forbidden, but not if cycle counter moves only between 1 and 15.
  - -> 3792 different Ids allowed

###  9.4. <a id='cobs/r-encoding-for-n-byte-packages'></a>COBS/R encoding for n-byte packages

- This looks similar to 1-byte and 2-byte encoding and is not shown here.
- Some super fast code for 3- and 4-byte packet encoding is also possible.
- All *trice* packages are much shorter than 255 bytes so the COBS/R encoding is cheap.

###  9.5. <a id='cobs/r-encoding'></a>COBS/R encoding

Packages are COBS/R encoded (without containing `00` bytes) and separated by a `00` byte. This allows the transfer of n-byte packages without the need to decide the meaning of the payload, means, how many bits are ID and how many bits are value is simply a configuration question. The COBS/R encoding usually has the same length as the unencoded data and sometimes has one byte more but an additional 00 is needed for secure package separation. This way the ID bit count is adjustable to the real communication needs because a data disturbance is easily detectable by just waiting for the next 0.


####  9.5.1. <a id='cobs/r-encoding-for-n-byte-packages-1'></a>COBS/R encoding for n-byte packages

This looks similar to 1-byte and 2-byte encoding and is not shown here.
Some super fast code for 3- and 4-byte packet encoding is also possible.

####  9.5.2. <a id='decoded-cobs/r-package-interpreter'></a>Decoded COBS/R package interpreter

How the packages are to interpret is a question of software configuration. When a decoded COBS/R package is to interpret, the known package length is used to choose the right interpreter. For example all multiple of 8 length packages are possibly XTEA encrypted. Also a fixed-size ID is usable. It is also possible to have several `trice` messages inside a packet. That makes sense to reach a multiple of 8-byte message length good for encryption.

- Examples for ID - value apportionment (these are only thinkable options):

Hint: The value space itself is usable according to ID, for example a 32 bit value space could be a 16-bit and two 8-bit values.

| package byte count |   optional usage                      | apportionment example
| -----------------: |   --------------------------          | :-
|                  2 |   on|off time measurement             | `Ivvvvvvv vvvvvvvv`
|                  2 |   256 different value events          | `IIIIIIII vvvvvvvv`
|                  8 | **XTEA, 256 different 56 bit values** | `IIIIIIII vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv`
|                 10 |   65535 different 64 bit values       | `IIIIIIII IIIIIIII vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv`
|                  6 |   65535 different 32 bit values       | `IIIIIIII IIIIIIII vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv`
|                  4 |   65535 different 16 bit values       | `IIIIIIII IIIIIIII vvvvvvvv vvvvvvvv`
|                  3 |   65535 different  8 bit values       | `IIIIIIII IIIIIIII vvvvvvvv`
|                  2 |   65535 different no-value messages   | `IIIIIIII IIIIIIII`
|                  1 |   8 different  5 bit values           | `IIIvvvvv`
|                  1 |   256 different no-value messages     | `IIIIIIII`
|                  4 |   32768 different 17-bit messages     | `IIIIIIII IIIIIIIx vvvvvvvv vvvvvvvv`
|                  8 |   XTEA, 2 messages in one packet      | `IIIIIIII IIIIIIII vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv IIIIIIII vvvvvvvv`
|                 16 |   XTEA, 5 messages in one packet      | `IIIIIIII IIIIIIII vvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvv IIIIIIII vvvvvvvv IIIIIIII IIIIIIII vvvvvvvv vvvvvvvv IIIIIIII vvvvvvvv IIIIIIII IIIIIIII`

- To keeps things simple:
- The first byte in a package is an ID byte optionally followed by more ID bytes

| ID coding                    | ID bits |         ID range  | remark
| ---------------------------: | ------: | ----------------: | -
| `0IIIIIII`                   |  7      |     0 ...     127 | 
| `10IIIIII IIIIIIII`          | 14      |   128 ...   16383 | 0 ...   127 unused (reserved)
| `100IIIII IIIIIIII IIIIIIII` | 21      | 16384 ... 2097151 | 0 ... 16383 unused (reserved)

- How many value bits are following an ID and how they are to interpret is coded inside the ID information.

- When an COBS/R package was successfully transmitted and decoded its interpretation is a matter of the general configuration.

| ID coding                              | package length    | ID bits |   ID range  |   ID map area   |  remark
| :------------------------------------- | -------------:    | ------: | ----------: |         -:      |  :-
| ``                                     |       0           |    0    |         0   |                 |  reserved, usable as a special very short message consisting of only one 0-byte as COBS/R message|
| `IIIIIIIv`                             |       1           |    7    | 0 ...   127 |     0 ...   127 |  one value bit, avoid IDs 0 and 1 for fast COBS/R encoding
| `1IIIIIII vvvvvvvv`                    |       2           |    7    | 0 ...   127 |   256 ...   383 |  1 value byte
| `00IIIIII IIIIIIII`                    |       2           |   14    | 0 ... 16383 |  1024 ... 16383 |  0 ... 1023 unused, no value bytes
| `01xxxxxx xxxxxxxx`                    |       2           |         |             |        -        |  2^14 packets unused (reserved), usable as a special 2-bytes message
| `1IIIIIII vvvvvvvv vvvvvvvv`           |       3           |    7    | 0 ...   127 |   384 ...   511 |  2 value bytes, avoid IDs 0 and 1 for fast COBS/R encoding
| `00IIIIII IIIIIIII vvvvvvvv`           |       3           |   14    | 0 ... 16383 | 16384 ... 32767 |  1 value byte
| `01xxxxxx xxxxxxxx xxxxxxxx`           |       3           |         |             |        -        |  2^22 packets unused (reserved), usable as a special 3-bytes message
| `1IIIIIII vvvvvvvv vvvvvvvv vvvvvvvv`  |       4           |    7    | 0 ...   127 |   512       639 |  3 value bytes, avoid IDs 0 and 1 for fast COBS/R encoding
| `00IIIIII IIIIIIII vvvvvvvv vvvvvvvv`  |       4           |   14    | 0 ... 16383 | 32768 ... 49151 |  2 value bytes
| `01xxxxxx xxxxxxxx xxxxxxxx xxxxxxxx`  |       4           |         |             |        -        |  2^14 packets unused (reserved), usable as a special 4-bytes message
| 5 ... 7 `xxxxxxxx`                     |     5,6,7         |         |             |        -        |  reserved, usable as a special 5-7-bytes message
| `1IIIIIII` + 7 `vvvvvvvv`              |       8           |    7    | 0 ...   127 |   640 ...   767 |  7 value bytes, avoid IDs 0 and 1 for fast COBS/R encoding
| `00IIIIII IIIIIIII` + 6 `vvvvvvvv`     |       8           |   14    | 0 ... 16383 | 49152 ... 65535 |  6 value bytes
| `01xxxxxx` + 7 `xxxxxxx`               |       8           |         |             |        -        |  2^30 packets unused (reserved), usable as a special 4-bytes message
| 9 ... 15 `xxxxxxxx`                    |9,10,11,12,13,14,15|         |             |        -        |  reserved, usable as a special 9-15-bytes message
| `1IIIIIII` + 15 `vvvvvvvv`             |      16           |    7    | 0 ...   127 |   768 ...   895 | 15 value bytes
| `00IIIIII IIIIIIII` + 14 `vvvvvvvv`    |      16           |   14    | 0 ... 16383 | 65536 ... 81919 | 14 value bytes
| `01xxxxxx` + 15 `xxxxxxxx`             |      16           |         |             |        -        |  2^30 packets unused (reserved), usable as a special 4-bytes message
| n `xxxxxxxx`                           |       n           |         |             |        -        |  n > 16, n mod 8 != 0, reserved, usable as a special n-bytes message
| `1IIIIIII` + (n-1) `vvvvvvvv`          |       n           |    7    | 0 ...   127 |   896 ...  1023 |  n > 16, n mod 8 == 0, n-1 value bytes
| `00IIIIII IIIIIIII` + (n-2) `vvvvvvvv` |       n           |   14    | 0 ... 16383 | 81920 ... 98304 |  n > 16, n mod 8 == 0, n-2 value bytes
| `01xxxxxx` + (n-1) `xxxxxxxx`          |       n           |         |             |        -        |  n > 16, n mod 8 == 0, n-2 value bytes, 2^((n-1)*8+6) packets unused (reserved), usable as a special 4-bytes message

If less value bytes are needed padding bytes are used.

`Trice( DESCRIPTOR, Id(0), "fmtString", ... )`

| Legacy notation                    | COBSR notation                       | ID coding
| :-                                 | :-                                   | :-
| Trice0( id(0), "text" );           | Trice0( IDE(0), "text" );            | `00IIIIII IIIIIIII`
|                                    | TriceB( ID7(0), "text", 1 );         | `IIIIIIIv`
| Trice8( id(0), "text", 255 );      | Trice8( ID7(0), "text", 255 );       | `1IIIIIII vvvvvvvv`
| Trice8( id(0), "text", 255, 255 ); | Trice8( ID7(0), "text", 255, 255 );  | `1IIIIIII vvvvvvvv vvvvvvvv`
| Trice16( id(0), "text", 65535 );   | Trice16( ID7(0), "text", 65535 );    | `1IIIIIII vvvvvvvv vvvvvvvv`
| trice8( Id(0), "text" );           | Trice8( IDE(0), "text", 255 );       | `00IIIIII IIIIIIII vvvvvvvv` 
| trice16( Id(0), "text" );          | Trice16( IDE(0), "text", 65535 );    | `00IIIIII IIIIIIII vvvvvvvv vvvvvvvv`

ID7(n) =  7-bit ID
IDE(n) = 14-bit ID

##  10. <a id='interpreter-for-decoded-cobs/r-package'></a>Interpreter for decoded COBS/R package

###  10.1. <a id='important'></a>IMPORTANT

- After receiving and decoding a COBS/R package, the receiver can decide according to the package length and its starting bits what to do with it:
  - Package lengths 2, 3, 4, 6, 10, 18, 34, 66 starting with four 0-bits are trice logs.
    - Treat as received *trice* message.
  - Multiple of 8 bytes packages are used for XTEA encryption.
    - The decrypted packet is treated again as a COBS/R encoded byte stream and handled recursively the same way.
    - This way several COBS/R encoded data packages can be joint in one package for encryption.
      - Empty COBS/R packages are `00` bytes and used to reach the next multiple of 8-bytes COBS/R sequence.
  - All other packages are useable for other protocols (marked as reserved).
    - Ignore, route forward or call user handler.
    - 1-byte COBS/R packages are not recommended for numerous data. Because of the delimiter byte, are only ~50% bandwidth usable.

###  10.2. <a id='encoding-table-0-legend'></a>Encoding table 0 legend

| Legend | Meaning                                                           |
| :-     | :---------------------------------------------------------------- |
| ...n   | totally n times                                                   |
| I\|iiii| 4 Id-bits (half byte)                                             |
| V\|vvvv| 4 value bits                                                      |
| X\|xxxx| 4 arbitrary bits (any half byte )                                 |
| Y\|yyyy| 4 arbitrary bits, but at least one must be 1 (any half byte != 0) |

###  10.3. <a id='encoding-table-0-(without-cycle-counter)'></a>Encoding table 0 (without cycle counter)

|half bytes      | same as bits                     | bytes|ID bits| ID range    |ID map| remark                                                                      |
| -              | -------------------------------- |:----:| :---: | :------:    |  :-: |     :-                                                                      |
|` `             | ` `                              |    0 |       |             |      | COBS/R padding byte                                                         |
|`0I II`         |`0000iiii iiiiiiii`               |    2 |   12  | 0\-4095     |  0   | `TRICE0`                                                                    |
|`0I II VV`      |`0000iiii iiiiiiii vvvvvvvv`      |    3 |   12  | 4096\- 8191 |  1   | `TRICE8_1`                                                                  |
|`0I II VV VV`   |`0000iiii iiiiiiii vvvvvvvv...2`  |    4 |   12  | 8192\-12287 |  2   | `TRICE8_2`, `TRICE16_1`                                                     |
|`0I II VV...4`  |`0000iiii iiiiiiii vvvvvvvv...4`  |    6 |   12  |12288\-16383 |  3   | `TRICE8_3`, `TRICE8_4`, `TRICE16_1`, `TRICE16_2`,  `TRICE32_1`              |
|`0I II VV...8`  |`0000iiii iiiiiiii vvvvvvvv...8`  |   10 |   12  |16384\-20479 |  4   | `TRICE8_5`...`TRICE8_8`, `TRICE16_3`, `TRICE16_4`, `TRICE32_2`, `TRICE64_1` |
|`0I II VV...16` |`0000iiii iiiiiiii vvvvvvvv...16` |   18 |   12  |20480\-24575 |  5   | `TRICE116_5`...`TRICE16_8`, `TRICE32_3`...`TRICE32_4`, `TRICE64_2`          |
|`0I II VV...32` |`0000iiii iiiiiiii vvvvvvvv...32` |   34 |   12  |24576\-28671 |  6   | `TRICE32_5`...`TRICE32_8`, `TRICE64_3`...`TRICE64_4`                        |
|`0I II VV...64` |`0000iiii iiiiiiii vvvvvvvv...64` |   66 |   12  |28672\-32767 |  7   | `TRICE64_5`...`TRICE64_8`                                                   |
|`YX XX`         |`yyyyxxxx xxxxxxxx`               |    2 |       |             |      | reserved                                                                    |
|`YX XX XX...2^n`|`yyyyxxxx xxxxxxxx xxxxxxxx...2^n`|2+2^n |       |             |      | reserved, n = 0...6                                                         |
|`XX...8*n`      |`xxxxxxxx...8*n`                  |  8*n |       |             |      | encrypted or reserved                                                       |
|`XX...n`        |`xxxxxxxx...n`                    |    n |       |             |      | reserved, n%8 != 0 && n != 2+2^m for m = 0...6                              |

- All packages are as encoded COBS/R sometimes 1 byte longer and always followed by the delimiter `00`byte.
- The ID map number can be deduced from the package length and needs no transmission.
  - So only the 12 lower ID bits are transmitted.

###  10.4. <a id='encoding-table-1-legend'></a>Encoding table 1 legend

| Legend | Meaning                           |
| :-     | :---------------------------------|
| ...n   | totally n times                   |
| I\|iiii| 4 Id-bits (half byte)             |
| V\|vvvv| 4 value bits                      |
| X\|xxxx| 4 arbitrary bits (any half byte ) |
| C\|cccc| 4 cycle counter bits              |

###  10.5. <a id='encoding-table-1-(with-4-bit-cycle-counter)'></a>Encoding table 1 (with 4-bit cycle counter)

|half bytes      | same as bits                     | bytes|ID bits| ID range    |ID map| remark                                                                      |
| -              | -------------------------------- |:----:| :---: | :------:    |  :-: |     :-                                                                      |
|` `             | ` `                              |    0 |       |             |      | COBS/R padding byte                                                         |
|`II IC`         |`iiiiiiii iiiicccc`               |    2 |   12  | 0\-4095     |  0   | `TRICE0`                                                                    |
|`II IC VV`      |`iiiiiiii iiiicccc vvvvvvvv`      |    3 |   12  | 4096\- 8191 |  1   | `TRICE8_1`                                                                  |
|`II IC VV VV`   |`iiiiiiii iiiicccc vvvvvvvv...2`  |    4 |   12  | 8192\-12287 |  2   | `TRICE8_2`, `TRICE16_1`                                                     |
|`II IC VV...4`  |`iiiiiiii iiiicccc vvvvvvvv...4`  |    6 |   12  |12288\-16383 |  3   | `TRICE8_3`, `TRICE8_4`, `TRICE16_1`, `TRICE16_2`,  `TRICE32_1`              |
|`II IC VV...8`  |`iiiiiiii iiiicccc vvvvvvvv...8`  |   10 |   12  |16384\-20479 |  4   | `TRICE8_5`...`TRICE8_8`, `TRICE16_3`, `TRICE16_4`, `TRICE32_2`, `TRICE64_1` |
|`II IC VV...16` |`iiiiiiii iiiicccc vvvvvvvv...16` |   18 |   12  |20480\-24575 |  5   | `TRICE116_5`...`TRICE16_8`, `TRICE32_3`...`TRICE32_4`, `TRICE64_2`          |
|`II IC VV...32` |`iiiiiiii iiiicccc vvvvvvvv...32` |   34 |   12  |24576\-28671 |  6   | `TRICE32_5`...`TRICE32_8`, `TRICE64_3`...`TRICE64_4`                        |
|`II IC VV...64` |`iiiiiiii iiiicccc vvvvvvvv...64` |   66 |   12  |28672\-32767 |  7   | `TRICE64_5`...`TRICE64_8`                                                   |
|`XX...8*n`      |`xxxxxxxx...8*n`                  |  8*n |       |             |      | encrypted                                                                   |
|`XX...n`        |`xxxxxxxx...n`                    |    n |       |             |      | reserved, n%8 != 0 && n != 0, 2, 3, 4, 6, 10, 18, 34, 66                    |

###  10.6. <a id='encoding-table-2-(with-8-bit-cycle-counter)'></a>Encoding table 2 (with 8-bit cycle counter)

|half bytes      | same as bits                              | bytes|ID bits| ID range  |ID map| remark                                                                      |
| -              | --------------------------------          |:----:| :---: | :------:  |  :-: |     :-                                                                      |
|` `             | ` `                                       |    0 |       |           |      | COBS/R padding byte                                                         |
|`II IC`         |`iiiiiiii iiiiiiii cccccccc`               |    3 |   16  | 1 - 65535 |  0   | `TRICE0`                                                                    |
|`II IC VV`      |`iiiiiiii iiiiiiii cccccccc vvvvvvvv`      |    4 |   16  | 1 - 65535 |  1   | `TRICE8_1`                                                                  |
|`II IC VV VV`   |`iiiiiiii iiiiiiii cccccccc vvvvvvvv...2`  |    5 |   16  | 1 - 65535 |  2   | `TRICE8_2`, `TRICE16_1`                                                     |
|`II IC VV...4`  |`iiiiiiii iiiiiiii cccccccc vvvvvvvv...4`  |    7 |   16  | 1 - 65535 |  3   | `TRICE8_3`, `TRICE8_4`, `TRICE16_1`, `TRICE16_2`,  `TRICE32_1`              |
|`II IC VV...8`  |`iiiiiiii iiiiiiii cccccccc vvvvvvvv...8`  |   17 |   16  | 1 - 65535 |  4   | `TRICE8_5`...`TRICE8_8`, `TRICE16_3`, `TRICE16_4`, `TRICE32_2`, `TRICE64_1` |
|`II IC VV...16` |`iiiiiiii iiiiiiii cccccccc vvvvvvvv...16` |   19 |   16  | 1 - 65535 |  5   | `TRICE16_5`...`TRICE16_8`, `TRICE32_3`...`TRICE32_4`, `TRICE64_2`           |
|`II IC VV...32` |`iiiiiiii iiiiiiii cccccccc vvvvvvvv...32` |   35 |   16  | 1 - 65535 |  6   | `TRICE32_5`...`TRICE32_8`, `TRICE64_3`...`TRICE64_4`                        |
|`II IC VV...64` |`iiiiiiii iiiiiiii cccccccc vvvvvvvv...64` |   67 |   16  | 1 - 65535 |  7   | `TRICE64_5`...`TRICE64_8`                                                   |
|`XX...8*n`      |`xxxxxxxx...8*n`                           |  8*n |       |           |      | encrypted                                                                   |
|`XX...n`        |`xxxxxxxx...n`                             |    n |       |           |      | reserved, n%8 != 0 && n != 0, 2, 3, 4, 6, 10, 18, 34, 66                    |

##  11. <a id='fast-trice-data-storing'></a>Fast TRICE data storing

...

<p align="right">(<a href="#top">back to top</a>)</p>

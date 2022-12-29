/*! \file internalTrice8.h
\author thomas.hoehenleitner [at] seerose.net
*******************************************************************************/

#include <stdint.h>


#define TRICE_PUT8_1(  v0 )                                                 TRICE_PUT(                                                   TRICE_BYTE0(v0));

#define TRICE_PUT8_2(  v0, v1                                           )   TRICE_PUT(                                   TRICE_BYTE1(v1) |TRICE_BYTE0(v0));
                                                                            
#define TRICE_PUT8_3(  v0, v1, v2                                       )   TRICE_PUT(                  TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0));
                                                                            
#define TRICE_PUT8_4(  v0, v1, v2, v3                                   )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0));
                                                                            
#define TRICE_PUT8_5(  v0, v1, v2, v3, v4                               )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT(                                                    TRICE_BYTE0(v4)); 
                                                                            
#define TRICE_PUT8_6(  v0, v1, v2, v3, v4, v5                           )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT(                                   TRICE_BYTE1(v5) |TRICE_BYTE0(v4)); 
                                                                            
#define TRICE_PUT8_7(  v0, v1, v2, v3, v4, v5, v6                       )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT(                  TRICE_BYTE2(v6) |TRICE_BYTE1(v5) |TRICE_BYTE0(v4));
                                                                            
#define TRICE_PUT8_8(  v0, v1, v2, v3, v4, v5, v6, v7                   )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT( TRICE_BYTE3(v7) |TRICE_BYTE2(v6) |TRICE_BYTE1(v5) |TRICE_BYTE0(v4));
                                                                            
#define TRICE_PUT8_9(  v0, v1, v2, v3, v4, v5, v6, v7, v8               )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT( TRICE_BYTE3(v7) |TRICE_BYTE2(v6) |TRICE_BYTE1(v5) |TRICE_BYTE0(v4)); \
                                                                            TRICE_PUT(                                                    TRICE_BYTE0(v8));
                                                                            
#define TRICE_PUT8_10( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9           )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT( TRICE_BYTE3(v7) |TRICE_BYTE2(v6) |TRICE_BYTE1(v5) |TRICE_BYTE0(v4)); \
                                                                            TRICE_PUT(                                   TRICE_BYTE1(v9) |TRICE_BYTE0(v8));
                                                                            
#define TRICE_PUT8_11( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10      )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT( TRICE_BYTE3(v7) |TRICE_BYTE2(v6) |TRICE_BYTE1(v5) |TRICE_BYTE0(v4)); \
                                                                            TRICE_PUT(                  TRICE_BYTE2(v10)|TRICE_BYTE1(v9) |TRICE_BYTE0(v8));
                                                                            
#define TRICE_PUT8_12( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 )   TRICE_PUT( TRICE_BYTE3(v3) |TRICE_BYTE2(v2) |TRICE_BYTE1(v1) |TRICE_BYTE0(v0)); \
                                                                            TRICE_PUT( TRICE_BYTE3(v7) |TRICE_BYTE2(v6) |TRICE_BYTE1(v5) |TRICE_BYTE0(v4)); \
                                                                            TRICE_PUT( TRICE_BYTE3(v11)|TRICE_BYTE2(v10)|TRICE_BYTE1(v9) |TRICE_BYTE0(v8));



//! TRICE8_1 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 a 8 bit bit value
#define TRICE8_1( tid, pFmt, v0 ) \
    TRICE_ENTER tid; CNTC(1); \
    TRICE_PUT8_1( v0 )  \
    TRICE_LEAVE

//! TRICE8_2 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v1 are 8 bit bit values
#define TRICE8_2( id, pFmt, v0, v1 ) \
    TRICE_ENTER id; CNTC(2); \
    TRICE_PUT8_2 ( v0, v1 ); \
    TRICE_LEAVE

//! TRICE8_3 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v2 are 8 bit bit values
#define TRICE8_3( id, pFmt, v0, v1, v2 ) \
    TRICE_ENTER id; CNTC(3); \
    TRICE_PUT8_3( v0, v1, v2 ); \
    TRICE_LEAVE

//! TRICE8_4 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v3 are 8 bit bit values
#define TRICE8_4( id, pFmt, v0, v1, v2, v3 ) \
    TRICE_ENTER id; CNTC(4); \
    TRICE_PUT8_4( v0, v1, v2, v3  ); \
    TRICE_LEAVE

//! TRICE8_5 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v4 are 8 bit bit values
#define TRICE8_5( id, pFmt, v0, v1, v2, v3, v4 ) \
    TRICE_ENTER id; CNTC(5); \
    TRICE_PUT8_5( v0, v1, v2, v3, v4  ); \
    TRICE_LEAVE

//! TRICE8_6 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v5 are 8 bit bit values
#define TRICE8_6( id, pFmt, v0, v1, v2, v3, v4, v5 ) \
    TRICE_ENTER id; CNTC(6); \
    TRICE_PUT8_6( v0, v1, v2, v3, v4, v5 ); \
    TRICE_LEAVE

//! TRICE8_8 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v6 are 8 bit bit values
#define TRICE8_7( id, pFmt, v0, v1, v2, v3, v4, v5, v6 ) \
    TRICE_ENTER id; CNTC(7); \
    TRICE_PUT8_7( v0, v1, v2, v3, v4, v5, v6 ); \
    TRICE_LEAVE

//! TRICE8_8 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v7 are 8 bit bit values
#define TRICE8_8( id, pFmt, v0, v1, v2, v3, v4, v5, v6, v7 ) \
    TRICE_ENTER id; CNTC(8); \
    TRICE_PUT8_8( v0, v1, v2, v3, v4, v5, v6, v7 ); \
    TRICE_LEAVE

//! TRICE8_8 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v7 are 8 bit bit values
#define TRICE8_9( id, pFmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) \
    TRICE_ENTER id; CNTC(9); \
    TRICE_PUT8_9( v0, v1, v2, v3, v4, v5, v6, v7, v8 ); \
    TRICE_LEAVE

//! TRICE8_8 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v7 are 8 bit bit values
#define TRICE8_10( id, pFmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) \
    TRICE_ENTER id; CNTC(10); \
    TRICE_PUT8_10( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ); \
    TRICE_LEAVE

//! TRICE8_8 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v7 are 8 bit bit values
#define TRICE8_11( id, pFmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) \
    TRICE_ENTER id; CNTC(11); \
    TRICE_PUT8_11( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ); \
    TRICE_LEAVE

//! TRICE8_12 writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 - v11 are 8 bit bit values
#define TRICE8_12( id, pFmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_ENTER id; CNTC(12); \
    TRICE_PUT8_12( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_LEAVE



#define trice8_1( fmt, v0 ) //!< trice8_1 is an empty macro
#define trice8_2( fmt, v0, v1 ) //!< trice8_2 is an empty macro
#define trice8_3( fmt, v0, v1, v2 ) //!< trice8_3 is an empty macro
#define trice8_4( fmt, v0, v1, v2, v3 ) //!< trice8_4 is an empty macro
#define trice8_5( fmt, v0, v1, v2, v3, v4 ) //!< trice8_5 is an empty macro
#define trice8_6( fmt, v0, v1, v2, v3, v4, v5 ) //!< trice8_6 is an empty macro
#define trice8_7( fmt, v0, v1, v2, v3, v4, v5, v6 ) //!< trice8_7 is an empty macro
#define trice8_8( fmt, v0, v1, v2, v3, v4, v5, v6, v7 ) //!< trice8_8 is an empty macro
#define trice8_9( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) //!< trice8_9 is an empty macro
#define trice8_10( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) //!< trice8_10 is an empty macro
#define trice8_11( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) //!< trice8_11 is an empty macro
#define trice8_12( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) //!< trice8_12 is an empty macro

#define trice8_COUNT(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12, NAME,...) NAME
#define trice8(fmt, ...) trice8_COUNT(__VA_ARGS__,trice8_12,trice8_11,trice8_10,trice8_9,trice8_8,trice8_7,trice8_6,trice8_5,trice8_4,trice8_3,trice8_2,trice8_1)(fmt, __VA_ARGS__)

#define trice8_COUNT(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12, NAME,...) NAME
#define trice8_M(tid,fmt, ...) trice8_COUNT(__VA_ARGS__,trice8_12_M,trice8_11_M,trice8_10_M,trice8_9_M,trice8_8_M,trice8_7_M,trice8_6_M,trice8_5_M,trice8_4_M,trice8_3_M,trice8_2_M,trice8_1_M)(tid,fmt, __VA_ARGS__)

//! trice8_1_m writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 a 8 bit bit value
#define trice8_1_m( tid, v0 ) \
    TRICE_ENTER \
    TRICE_PUT( (1<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_1( v0 ) \
    TRICE_LEAVE

#define trice8_2_m( tid, v0, v1 ) \
    TRICE_ENTER \
    TRICE_PUT( (2<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_2( v0, v1); \
    TRICE_LEAVE

#define trice8_3_m( tid, v0, v1, v2 ) \
    TRICE_ENTER \
    TRICE_PUT( (3<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_3 ( v0, v1, v2 ); \
    TRICE_LEAVE

#define trice8_4_m( tid, v0, v1, v2, v3 ) \
    TRICE_ENTER \
    TRICE_PUT( (4<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_4( v0, v1, v2, v3 ); \
    TRICE_LEAVE

#define trice8_5_m( tid, v0, v1, v2, v3, v4 ) \
    TRICE_ENTER \
    TRICE_PUT( (5<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_5( v0, v1, v2, v3, v4 ); \
    TRICE_LEAVE

#define trice8_6_m( tid, v0, v1, v2, v3, v4, v5 ) \
    TRICE_ENTER \
    TRICE_PUT( (6<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_6( v0, v1, v2, v3, v4, v5 ); \
    TRICE_LEAVE

#define trice8_7_m( tid, v0, v1, v2, v3, v4, v5, v6 ) \
    TRICE_ENTER \
    TRICE_PUT( (7<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_7( v0, v1, v2, v3, v4, v5, v6 ); \
    TRICE_LEAVE

#define trice8_8_m( tid, v0, v1, v2, v3, v4, v5, v6, v7 ) \
    TRICE_ENTER \
    TRICE_PUT( (8<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_8( v0, v1, v2, v3, v4, v5, v6, v7 ); \
    TRICE_LEAVE

#define trice8_9_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) \
    TRICE_ENTER \
    TRICE_PUT( (9<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_9( v0, v1, v2, v3, v4, v5, v6, v7, v8 ); \
    TRICE_LEAVE

#define trice8_10_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) \
    TRICE_ENTER \
    TRICE_PUT( (10<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_10( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ); \
    TRICE_LEAVE

#define trice8_11_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) \
    TRICE_ENTER \
    TRICE_PUT( (11<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_11( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ); \
    TRICE_LEAVE

#define trice8_12_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_ENTER \
    TRICE_PUT( (12<<24) | ((TRICE_CYCLE)<<16) | (0x4000|(tid)) ); \
    TRICE_PUT8_12( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_LEAVE

#define trice8_1_M( tid,  fmt, v0 ) trice8_1_fn( tid,  (uint8_t)(v0) ) //!< trice8_1_M is a macro calling a function to reduce code size.
#define trice8_2_M( tid,  fmt, v0, v1 ) trice8_2_fn( tid,  (uint8_t)(v0), (uint8_t)(v1) ) //!< trice8_2_M is a macro calling a function to reduce code size.
#define trice8_3_M( tid,  fmt, v0, v1, v2 ) trice8_3_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2) ) //!< trice8_3_M is a macro calling a function to reduce code size.
#define trice8_4_M( tid,  fmt, v0, v1, v2, v3 ) trice8_4_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3) ) //!< trice8_4_M is a macro calling a function to reduce code size.
#define trice8_5_M( tid,  fmt, v0, v1, v2, v3, v4 ) trice8_5_fn( tid,  (uint8_t)v0, (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4) ) //!< trice8_5_M is a macro calling a function to reduce code size.
#define trice8_6_M( tid,  fmt, v0, v1, v2, v3, v4, v5 ) trice8_6_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5) ) //!< trice8_6_M is a macro calling a function to reduce code size.
#define trice8_7_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6 ) trice8_7_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6) ) //!< trice8_7_M is a macro calling a function to reduce code size.
#define trice8_8_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6, v7 ) trice8_8_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7) ) //!< trice8_8_M is a macro calling a function to reduce code size.
#define trice8_9_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) trice8_9_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8) ) //!< trice8_9_M is a macro calling a function to reduce code size.
#define trice8_10_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) trice8_10_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9) ) //!< trice8_10_M is a macro calling a function to reduce code size.
#define trice8_11_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) trice8_11_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9), (uint8_t)(v10) ) //!< trice8_11_M is a macro calling a function to reduce code size.
#define trice8_12_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) trice8_12_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9), (uint8_t)(v10), (uint8_t)(v11) ) //!< trice8_12_M is a macro calling a function to reduce code size.

void trice8_1_fn( uint16_t tid,  uint8_t v0 );
void trice8_2_fn( uint16_t tid,  uint8_t v0, uint8_t v1 );
void trice8_3_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2 );
void trice8_4_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3 );
void trice8_5_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4 );
void trice8_6_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5 );
void trice8_7_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6 );
void trice8_8_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7 );
void trice8_9_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8 );
void trice8_10_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9 );
void trice8_11_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9, uint8_t v10 );
void trice8_12_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9, uint8_t v10, uint8_t v11 );

#define Trice8_1( fmt, v0 ) //!< Trice8_1 is an empty macro
#define Trice8_2( fmt, v0, v1 ) //!< Trice8_2 is an empty macro
#define Trice8_3( fmt, v0, v1, v2 ) //!< Trice8_3 is an empty macro
#define Trice8_4( fmt, v0, v1, v2, v3 ) //!< Trice8_4 is an empty macro
#define Trice8_5( fmt, v0, v1, v2, v3, v4 ) //!< Trice8_5 is an empty macro
#define Trice8_6( fmt, v0, v1, v2, v3, v4, v5 ) //!< Trice8_6 is an empty macro
#define Trice8_7( fmt, v0, v1, v2, v3, v4, v5, v6 ) //!< Trice8_7 is an empty macro
#define Trice8_8( fmt, v0, v1, v2, v3, v4, v5, v6, v7 ) //!< Trice8_8 is an empty macro
#define Trice8_9( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) //!< Trice8_9 is an empty macro
#define Trice8_10( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) //!< Trice8_10 is an empty macro
#define Trice8_11( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) //!< Trice8_11 is an empty macro
#define Trice8_12( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) //!< Trice8_12 is an empty macro

#define Trice8_COUNT(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12, NAME,...) NAME
#define Trice8(fmt, ...) Trice8_COUNT(__VA_ARGS__,Trice8_12,Trice8_11,Trice8_10,Trice8_9,Trice8_8,Trice8_7,Trice8_6,Trice8_5,Trice8_4,Trice8_3,Trice8_2,Trice8_1)(fmt, __VA_ARGS__)

#define Trice8_COUNT(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12, NAME,...) NAME
#define Trice8_M(tid,fmt, ...) Trice8_COUNT(__VA_ARGS__,Trice8_12_M,Trice8_11_M,Trice8_10_M,Trice8_9_M,Trice8_8_M,Trice8_7_M,Trice8_6_M,Trice8_5_M,Trice8_4_M,Trice8_3_M,Trice8_2_M,Trice8_1_M)(tid,fmt, __VA_ARGS__)

//! Trice8_1_m writes trice data as fast as possible in a buffer.
//! \param id is a 16 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 a 8 bit bit value
#define Trice8_1_m( tid, v0 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 1<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_1( v0 ) \
    TRICE_LEAVE

#define Trice8_2_m( tid, v0, v1 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 2<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_2( v0, v1); \
    TRICE_LEAVE

#define Trice8_3_m( tid, v0, v1, v2 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 3<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_3 ( v0, v1, v2 ); \
    TRICE_LEAVE

#define Trice8_4_m( tid, v0, v1, v2, v3 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 4<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_4( v0, v1, v2, v3 ); \
    TRICE_LEAVE

#define Trice8_5_m( tid, v0, v1, v2, v3, v4 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 5<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_5( v0, v1, v2, v3, v4 ); \
    TRICE_LEAVE

#define Trice8_6_m( tid, v0, v1, v2, v3, v4, v5 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 6<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_6( v0, v1, v2, v3, v4, v5 ); \
    TRICE_LEAVE

#define Trice8_7_m( tid, v0, v1, v2, v3, v4, v5, v6 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 7<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_7( v0, v1, v2, v3, v4, v5, v6 ); \
    TRICE_LEAVE

#define Trice8_8_m( tid, v0, v1, v2, v3, v4, v5, v6, v7 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 8<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_8( v0, v1, v2, v3, v4, v5, v6, v7 ); \
    TRICE_LEAVE

#define Trice8_9_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 9<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_9( v0, v1, v2, v3, v4, v5, v6, v7, v8 ); \
    TRICE_LEAVE

#define Trice8_10_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 10<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_10( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ); \
    TRICE_LEAVE

#define Trice8_11_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 11<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_11( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ); \
    TRICE_LEAVE

#define Trice8_12_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_ENTER \
    uint16_t ts = TriceStamp16(); \
    TRICE_PUT(0x80008000|(tid<<16)|tid); \
    TRICE_PUT( 12<<24 | (TRICE_CYCLE<<16) | ts ); \
    TRICE_PUT8_12( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_LEAVE

#define Trice8_1_M( tid,  fmt, v0 ) Trice8_1_fn( tid,  (uint8_t)(v0) ) //!< Trice8_1_M is a macro calling a function to reduce code size.
#define Trice8_2_M( tid,  fmt, v0, v1 ) Trice8_2_fn( tid,  (uint8_t)(v0), (uint8_t)(v1) ) //!< Trice8_2_M is a macro calling a function to reduce code size.
#define Trice8_3_M( tid,  fmt, v0, v1, v2 ) Trice8_3_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2) ) //!< Trice8_3_M is a macro calling a function to reduce code size.
#define Trice8_4_M( tid,  fmt, v0, v1, v2, v3 ) Trice8_4_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3) ) //!< Trice8_4_M is a macro calling a function to reduce code size.
#define Trice8_5_M( tid,  fmt, v0, v1, v2, v3, v4 ) Trice8_5_fn( tid,  (uint8_t)v0, (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4) ) //!< Trice8_5_M is a macro calling a function to reduce code size.
#define Trice8_6_M( tid,  fmt, v0, v1, v2, v3, v4, v5 ) Trice8_6_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5) ) //!< Trice8_6_M is a macro calling a function to reduce code size.
#define Trice8_7_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6 ) Trice8_7_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6) ) //!< Trice8_7_M is a macro calling a function to reduce code size.
#define Trice8_8_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6, v7 ) Trice8_8_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7) ) //!< Trice8_8_M is a macro calling a function to reduce code size.
#define Trice8_9_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) Trice8_9_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8) ) //!< Trice8_9_M is a macro calling a function to reduce code size.
#define Trice8_10_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) Trice8_10_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9) ) //!< Trice8_10_M is a macro calling a function to reduce code size.
#define Trice8_11_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) Trice8_11_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9), (uint8_t)(v10) ) //!< Trice8_11_M is a macro calling a function to reduce code size.
#define Trice8_12_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) Trice8_12_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9), (uint8_t)(v10), (uint8_t)(v11) ) //!< Trice8_12_M is a macro calling a function to reduce code size.

void Trice8_1_fn( uint16_t tid,  uint8_t v0 );
void Trice8_2_fn( uint16_t tid,  uint8_t v0, uint8_t v1 );
void Trice8_3_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2 );
void Trice8_4_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3 );
void Trice8_5_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4 );
void Trice8_6_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5 );
void Trice8_7_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6 );
void Trice8_8_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7 );
void Trice8_9_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8 );
void Trice8_10_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9 );
void Trice8_11_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9, uint8_t v10 );
void Trice8_12_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9, uint8_t v10, uint8_t v11 );





#define TRice8_1( fmt, v0 ) //!< TRice8_1 is an empty macro
#define TRice8_2( fmt, v0, v1 ) //!< TRice8_2 is an empty macro
#define TRice8_3( fmt, v0, v1, v2 ) //!< TRice8_3 is an empty macro
#define TRice8_4( fmt, v0, v1, v2, v3 ) //!< TRice8_4 is an empty macro
#define TRice8_5( fmt, v0, v1, v2, v3, v4 ) //!< TRice8_5 is an empty macro
#define TRice8_6( fmt, v0, v1, v2, v3, v4, v5 ) //!< TRice8_6 is an empty macro
#define TRice8_7( fmt, v0, v1, v2, v3, v4, v5, v6 ) //!< TRice8_7 is an empty macro
#define TRice8_8( fmt, v0, v1, v2, v3, v4, v5, v6, v7 ) //!< TRice8_8 is an empty macro
#define TRice8_9( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) //!< TRice8_9 is an empty macro
#define TRice8_10( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) //!< TRice8_10 is an empty macro
#define TRice8_11( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) //!< TRice8_11 is an empty macro
#define TRice8_12( fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) //!< TRice8_12 is an empty macro

#define TRice8_COUNT(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12, NAME,...) NAME
#define TRice8(fmt, ...) TRice8_COUNT(__VA_ARGS__,TRice8_12,TRice8_11,TRice8_10,TRice8_9,TRice8_8,TRice8_7,TRice8_6,TRice8_5,TRice8_4,TRice8_3,TRice8_2,TRice8_1)(fmt, __VA_ARGS__)

#define TRice8_COUNT(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12, NAME,...) NAME
#define TRice8_M(tid,fmt, ...) TRice8_COUNT(__VA_ARGS__,TRice8_12_M,TRice8_11_M,TRice8_10_M,TRice8_9_M,TRice8_8_M,TRice8_7_M,TRice8_6_M,TRice8_5_M,TRice8_4_M,TRice8_3_M,TRice8_2_M,TRice8_1_M)(tid,fmt, __VA_ARGS__)

//! TRice8_1_m writes trice data as fast as possible in a buffer.
//! \param id is a 14 bit Trice id in upper 2 bytes of a 32 bit value
//! \param v0 a 8 bit bit value
#define TRice8_1_m( tid, v0 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 1<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_1( v0 ) \
    TRICE_LEAVE

#define TRice8_2_m( tid, v0, v1 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 2<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_2( v0, v1); \
    TRICE_LEAVE

#define TRice8_3_m( tid, v0, v1, v2 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 3<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_3 ( v0, v1, v2 ); \
    TRICE_LEAVE

#define TRice8_4_m( tid, v0, v1, v2, v3 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 4<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_4( v0, v1, v2, v3 ); \
    TRICE_LEAVE

#define TRice8_5_m( tid, v0, v1, v2, v3, v4 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 5<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_5( v0, v1, v2, v3, v4 ); \
    TRICE_LEAVE

#define TRice8_6_m( tid, v0, v1, v2, v3, v4, v5 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 6<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_6( v0, v1, v2, v3, v4, v5 ); \
    TRICE_LEAVE

#define TRice8_7_m( tid, v0, v1, v2, v3, v4, v5, v6 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 7<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_7( v0, v1, v2, v3, v4, v5, v6 ); \
    TRICE_LEAVE

#define TRice8_8_m( tid, v0, v1, v2, v3, v4, v5, v6, v7 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 8<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_8( v0, v1, v2, v3, v4, v5, v6, v7 ); \
    TRICE_LEAVE

#define TRice8_9_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 9<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_9( v0, v1, v2, v3, v4, v5, v6, v7, v8 ); \
    TRICE_LEAVE

#define TRice8_10_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 10<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_10( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ); \
    TRICE_LEAVE

#define TRice8_11_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 11<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_11( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ); \
    TRICE_LEAVE

#define TRice8_12_m( tid, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_ENTER \
    uint32_t ts = TRICE_HTOTL(TriceStamp32()); \
    TRICE_PUT((ts<<16) | 0xc000 | tid); \
    TRICE_PUT( 12<<24 | (TRICE_CYCLE<<16) | (ts>>16) ); \
    TRICE_PUT8_12( v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) \
    TRICE_LEAVE

#define TRice8_1_M( tid,  fmt, v0 ) TRice8_1_fn( tid,  (uint8_t)(v0) ) //!< TRice8_1_M is a macro calling a function to reduce code size.
#define TRice8_2_M( tid,  fmt, v0, v1 ) TRice8_2_fn( tid,  (uint8_t)(v0), (uint8_t)(v1) ) //!< TRice8_2_M is a macro calling a function to reduce code size.
#define TRice8_3_M( tid,  fmt, v0, v1, v2 ) TRice8_3_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2) ) //!< TRice8_3_M is a macro calling a function to reduce code size.
#define TRice8_4_M( tid,  fmt, v0, v1, v2, v3 ) TRice8_4_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3) ) //!< TRice8_4_M is a macro calling a function to reduce code size.
#define TRice8_5_M( tid,  fmt, v0, v1, v2, v3, v4 ) TRice8_5_fn( tid,  (uint8_t)v0, (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4) ) //!< TRice8_5_M is a macro calling a function to reduce code size.
#define TRice8_6_M( tid,  fmt, v0, v1, v2, v3, v4, v5 ) TRice8_6_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5) ) //!< TRice8_6_M is a macro calling a function to reduce code size.
#define TRice8_7_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6 ) TRice8_7_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6) ) //!< TRice8_7_M is a macro calling a function to reduce code size.
#define TRice8_8_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6, v7 ) TRice8_8_fn( tid,  (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7) ) //!< TRice8_8_M is a macro calling a function to reduce code size.
#define TRice8_9_M( tid,  fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8 ) TRice8_9_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8) ) //!< TRice8_9_M is a macro calling a function to reduce code size.
#define TRice8_10_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9 ) TRice8_10_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9) ) //!< TRice8_10_M is a macro calling a function to reduce code size.
#define TRice8_11_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 ) TRice8_11_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9), (uint8_t)(v10) ) //!< TRice8_11_M is a macro calling a function to reduce code size.
#define TRice8_12_M( tid, fmt, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 ) TRice8_12_fn( tid, (uint8_t)(v0), (uint8_t)(v1), (uint8_t)(v2), (uint8_t)(v3), (uint8_t)(v4), (uint8_t)(v5), (uint8_t)(v6), (uint8_t)(v7), (uint8_t)(v8), (uint8_t)(v9), (uint8_t)(v10), (uint8_t)(v11) ) //!< TRice8_12_M is a macro calling a function to reduce code size.

void TRice8_1_fn( uint16_t tid,  uint8_t v0 );
void TRice8_2_fn( uint16_t tid,  uint8_t v0, uint8_t v1 );
void TRice8_3_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2 );
void TRice8_4_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3 );
void TRice8_5_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4 );
void TRice8_6_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5 );
void TRice8_7_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6 );
void TRice8_8_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7 );
void TRice8_9_fn( uint16_t tid,  uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8 );
void TRice8_10_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9 );
void TRice8_11_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9, uint8_t v10 );
void TRice8_12_fn( uint16_t tid, uint8_t v0, uint8_t v1, uint8_t v2, uint8_t v3, uint8_t v4, uint8_t v5, uint8_t v6, uint8_t v7, uint8_t v8, uint8_t v9, uint8_t v10, uint8_t v11 );





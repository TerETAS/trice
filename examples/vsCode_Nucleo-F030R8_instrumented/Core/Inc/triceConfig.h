/*! \file triceConfig.h
\author Thomas.Hoehenleitner [at] seerose.net
*******************************************************************************/

#ifndef TRICE_CONFIG_H_
#define TRICE_CONFIG_H_

#ifdef __cplusplus
extern "C" {
#endif

#include "stm32f0xx_ll_system.h"

//! ms32 is a 32-bit millisecond counter, counting circular in steps of 1 every ms.
extern uint32_t ms32;

#if 1 // TestDebug

#define TRICE_BUFFER TRICE_RING_BUFFER

// trice l -p jlink -args "-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0" -showID "deb:%04x" -hs off -pw MySecret -pf COBS -d16=true
#define TRICE_DIRECT_OUTPUT 1
#define TRICE_DIRECT_XTEA_ENCRYPT 1
#define TRICE_DIRECT_OUT_FRAMING TRICE_FRAMING_COBS
#define TRICE_DIRECT_SEGGER_RTT_32BIT_WRITE 1

// trice l -p com4 -d16=false
#define TRICE_DEFERRED_OUTPUT 1
#define TRICE_DEFERRED_UARTA 1
#define TRICE_UARTA USART2

//#define TRICE_CGO 1
//#define TRICE_CYCLE_COUNTER 0

#endif


#if 0 // duration investigation

//#define TriceStamp16 (SysTick->VAL) // Counts from 31999 -> 0 in each ms.
//#define TriceStamp32  ms32
#define TRICE_UARTA   USART2

#define TRICE_BUFFER TRICE_RING_BUFFER

// trice l -p jlink -args "-Device STM32F030R8 -if SWD -Speed 4000 -RTTChannel 0" -pf NONE -d16 -showID "deb:%04x"
#define TRICE_DIRECT_OUTPUT 1
#define TRICE_DIRECT_OUT_FRAMING TRICE_FRAMING_COBS
#define TRICE_DIRECT_SEGGER_RTT_8BIT_WRITE 1

// trice l -p com4 -pf COBS -pw MySecret -showID "deb:%04x"
#define TRICE_DEFERRED_OUTPUT 1
//#define TRICE_DEFERRED_XTEA_ENCRYPT 1
#define TRICE_DEFERRED_OUT_FRAMING TRICE_FRAMING_COBS
#define TRICE_DEFERRED_UARTA 1

#endif




#ifdef __cplusplus
}
#endif

#endif /* TRICE_CONFIG_H_ */

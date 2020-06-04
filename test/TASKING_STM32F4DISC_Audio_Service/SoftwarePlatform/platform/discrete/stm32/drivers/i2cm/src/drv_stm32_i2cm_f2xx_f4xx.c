/* ------------------------------------------------------------
**
**  Copyright (c) 2013-2015 Altium Limited
**
**  This software is the proprietary, copyrighted property of
**  Altium Ltd. All Right Reserved.
**
**  SVN revision information:
**  $Rev: 14907 $:
**  $Date: 2015-01-19 13:30:51 +0100 (Mon, 19 Jan 2015) $:
**
** ------------------------------------------------------------
*/

#include <assert.h>

#if defined(STM32F2XX) || defined(STM32F4XX)

#include <stdint.h>
#include <drv_i2cm.h>
#include <time.h>
#include <timing.h>
#include <devices.h>
#include "drv_stm32_i2cm_cfg_instance.h"
#include "stm32_chip_config.h"
#ifdef STM32F2XX
#include "per_stm32f2xx_i2c_cfg_instance.h"
#include "stm32f2xx_i2c.h"
#endif
#ifdef STM32F4XX
#include "per_stm32f4xx_i2c_cfg_instance.h"
#include "stm32f4xx_i2c.h"
#endif
#include "../internal/drv_stm32_i2cm_internal.h"


i2cm_t i2cm_drv_table[DRV_STM32_I2CM_INSTANCE_COUNT];

/**
 * Configures the different clocks & GPIO ports and reset the peripheral and bus
 *
 * param per_cfg pointer to peripheral configuration
 * param drv pointer to driver context
 *
 * return 0 = okay, -1 = clock-line error, -2 = data-line error, -3 = unconnected pin(s), -4 = unknown id.
 */

#ifdef STM32F2XX
static int init_bus_fxxx( const per_stm32f2xx_i2c_cfg_instance_t * per_cfg, i2cm_t * drv)
#else
static int init_bus_fxxx( const per_stm32f4xx_i2c_cfg_instance_t * per_cfg, i2cm_t * drv)
#endif
{
    GPIO_TypeDef *scl_port;
    GPIO_TypeDef *sda_port;
    uint16_t     scl_pin;
    uint16_t     sda_pin;

    switch(per_cfg->instance_id)
    {
    case 0:
#if !(defined(PINCFG_I2C1_SDA) && defined(PINCFG_I2C1_SCL))
        return -3;
#else
        sda_port = PINCFG_I2C1_SDA_PORT;
        sda_pin  = PINCFG_I2C1_SDA_PIN;
        scl_port = PINCFG_I2C1_SCL_PORT;
        scl_pin  = PINCFG_I2C1_SCL_PIN;
        PINCFG_I2C1_SCL_PORT_CLOCK_ENABLE;
        PINCFG_I2C1_SDA_PORT_CLOCK_ENABLE;
        break;
#endif
    case 1:
#if !(defined(PINCFG_I2C2_SDA) && defined(PINCFG_I2C2_SCL))
        return -3;
#else
        sda_port = PINCFG_I2C2_SDA_PORT;
        sda_pin  = PINCFG_I2C2_SDA_PIN;
        scl_port = PINCFG_I2C2_SCL_PORT;
        scl_pin  = PINCFG_I2C2_SCL_PIN;
        PINCFG_I2C2_SCL_PORT_CLOCK_ENABLE;
        PINCFG_I2C2_SDA_PORT_CLOCK_ENABLE;
        break;
#endif
    case 2:
#if !(defined(PINCFG_I2C3_SDA) && defined(PINCFG_I2C3_SCL))
        return -3;
#else
        sda_port = PINCFG_I2C3_SDA_PORT;
        sda_pin  = PINCFG_I2C3_SDA_PIN;
        scl_port = PINCFG_I2C3_SCL_PORT;
        scl_pin  = PINCFG_I2C3_SCL_PIN;
        PINCFG_I2C3_SCL_PORT_CLOCK_ENABLE;
        PINCFG_I2C3_SDA_PORT_CLOCK_ENABLE;
        break;
#endif
    default:
        return -4;
    }

    GPIO_InitTypeDef gpio_init_struct;
    clock_t t;
    const clock_t timeout = CLOCKS_PER_SEC / 65536; // Actual speed should be 16384 kHz?

    // reset bus

    GPIO_WriteBit( scl_port, scl_pin, Bit_SET );
    GPIO_WriteBit( sda_port, sda_pin, Bit_SET );

    gpio_init_struct.GPIO_Speed = GPIO_Speed_50MHz; // 50 MHz
    gpio_init_struct.GPIO_Mode = GPIO_Mode_OUT;
    gpio_init_struct.GPIO_OType = GPIO_OType_OD;
    gpio_init_struct.GPIO_PuPd = GPIO_PuPd_NOPULL;

    gpio_init_struct.GPIO_Pin = scl_pin; // I2C: SCL
    GPIO_Init( scl_port, &gpio_init_struct );
    gpio_init_struct.GPIO_Pin = sda_pin; // I2C: SDA
    GPIO_Init( sda_port, &gpio_init_struct );

    // Generate 9 stop conditions, without start:
    //
    //      ___ ___     ___     ___     ___     ___     ___     ___     ___     ___     __________
    //  SCL ___|   |___| ' |___|   |___|   |___|   |___|   |___|   |___|   |___|   |___|
    //                   '
    //      _________    '___     ___     ___     ___     ___     ___     ___     ___     ________
    //  SDA _________|___|   |___|   |___|   |___|   |___|   |___|   |___|   |___|   |___|
    //                   '
    //                   stop

    for ( int i = 0; i < 9; i++ )
    {
        // Drop clock
        GPIO_WriteBit( scl_port, scl_pin, Bit_RESET );
        for ( t = clock() + timeout; clock() < t; ) __nop();

        // Drop data
        GPIO_WriteBit( sda_port, sda_pin, Bit_RESET );
        for ( t = clock() + timeout; clock() < t; ) __nop();

        // Assert clock
        GPIO_WriteBit( scl_port, scl_pin, Bit_SET );
        for ( t = clock() + timeout; clock() < t; ) __nop();
        if ( GPIO_ReadInputDataBit( scl_port, scl_pin ) == RESET ) return -1;   // Clock is kept low by a slave :-(

        GPIO_WriteBit( sda_port, sda_pin, Bit_SET );  // Assert data (hopefully)
        for ( t = clock() + timeout; clock() < t; ) __nop();

    }

    if ( GPIO_ReadInputDataBit( sda_port, sda_pin ) == RESET ) return -2;   // Did not finish, data kept low by slave

    switch(per_cfg->instance_id)
    {
    case 0:
        RCC_APB1PeriphClockCmd(RCC_APB1Periph_I2C1, ENABLE);
        drv->i2cx = I2C1;
        break;

    case 1:
        RCC_APB1PeriphClockCmd(RCC_APB1Periph_I2C2, ENABLE);
        drv->i2cx = I2C2;
        break;

    case 2:
        RCC_APB1PeriphClockCmd(RCC_APB1Periph_I2C3, ENABLE);
        drv->i2cx = I2C3;
        break;

    }

    per_cfg->pinconfig();

    // Reset the I2C peripheral - it is not happy with unexpected stop conditions
    I2C_SoftwareResetCmd( drv->i2cx, ENABLE );
    I2C_SoftwareResetCmd( drv->i2cx, DISABLE );

    return 0;
}

/**
 * @brief Open the I2CM device driver and initializes the hardware
 *
 * This function opens the device driver for an I2C peripheral on the STM32 core. You should call
 * this function once per instantiation.
 *
 * @param id Instance identifier as generated by the software platform
 *
 * @return NULL on error, pointer to instance context otherwise
 */

i2cm_t * i2cm_open( unsigned int id )
{
    // Assert for validity check on id should be here
//    assert( id < DRV_I2CM_INSTANCE_COUNT );

    i2cm_t * drv;
    const drv_stm32_i2cm_cfg_instance_t * restrict drv_cfg;
#ifdef STM32F2XX
    const per_stm32f2xx_i2c_cfg_instance_t * restrict per_cfg;
    drv_cfg = &drv_stm32_i2cm_instance_table[id];
    per_cfg = &per_stm32f2xx_i2c_instance_table[drv_cfg->per_stm32_i2c];
#endif
#ifdef STM32F4XX
    const per_stm32f4xx_i2c_cfg_instance_t * restrict per_cfg;
    drv_cfg = &drv_stm32_i2cm_instance_table[id];
    per_cfg = &per_stm32f4xx_i2c_instance_table[drv_cfg->per_stm32_i2c];
#endif
    drv = &i2cm_drv_table[id];

    if ( drv->i2cx == NULL )
    {
        drv->i2c_frequency = drv_cfg->i2c_frequency;

        // Initialize the I2C peripheral
        I2C_StructInit( &drv->init );
        drv->init.I2C_Mode = I2C_Mode_I2C;
        drv->init.I2C_DutyCycle = I2C_DutyCycle_2;
        drv->init.I2C_OwnAddress1 = 0x00;
        drv->init.I2C_Ack = I2C_Ack_Enable;
        drv->init.I2C_AcknowledgedAddress = I2C_AcknowledgedAddress_7bit;
        drv->init.I2C_ClockSpeed = drv_cfg->i2c_frequency;

        if ( init_bus_fxxx( per_cfg, drv ) ) return NULL;

        I2C_Cmd( drv->i2cx, ENABLE );
        I2C_Init( drv->i2cx, &drv->init );
    }
    return drv;
}

#endif

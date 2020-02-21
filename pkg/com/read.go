// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package com

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

// https://stackoverflow.com/questions/50088669/golang-reading-from-serial

// export readBytes
func readBytes(s *serial.Port, count int) (int, []byte) {
	b := make([]byte, count) // the buffer size limits the read count
	n, err := s.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return n, b
}

// export readBytes
func readAtLeastBytes(s *serial.Port, count, msTimeout int) ([]byte, error) {
	buf := make([]byte, count) // the buffer size limits the read count
	var b []byte
	var n, ms int
	var err error
	for len(b) < count && ms < msTimeout {
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		b = append(b, buf[:n]...)
		if count == len(b) {
			return b, nil
		}
		buf = buf[n:]
		time.Sleep(50 * time.Millisecond) // has slight influence on cpu load
		ms += 50
	}
	return b, errors.New("read timeout")
}

// FindSerialPorts scans for COM ports between 0 and 100
func FindSerialPorts(config *serial.Config) {
	n := 0
	for n < 100 {
		n++
		config.Name = fmt.Sprintf("COM%d", n)
		stream, err := serial.OpenPort(config)
		if err == nil {
			fmt.Println("found", config.Name)
			stream.Close()
		}
	}
}

package hd44780

import "time"

const (
	// Timing constants
	ePulse = 2 * time.Microsecond
	eDelay = 140 * time.Microsecond

	// LCD 4x20 Mode Setting defaults
	Lines int = 4
	lcdLines int = 4
	Width int = 20 // Maximum characters per line
	lcdWidth int = 20 // Maximum characters per line
	lcdChr bool = true
	lcdCmd bool = false

	// HD44780 LCD Line RAM Physycal Addresses
	lcdLine1 = 0x80
	lcdLine2 = 0xC0
        lcdLine3 = 0x94
        lcdLine4 = 0xD4
)


var (
	BKPin int = 0
        RSPin int = 0 // GOIO 7   --> Raspberry Physical Pin 26  RS Bit/pin
        EPin  int = 0 // GPIO 8   --> Raspberry Physical Pin 24  Enable bit/pin
        D4Pin int = 0 // GPIO 25  --> Raspberry Physical Pin 22  Data4 bit/pin
        D5Pin int = 0 // GPIO 24  --> Raspberry Physical Pin 18  Data5 bit/pin
        D6Pin int = 0 // GPIO 23  --> Raspberry Physical Pin 16  Data6 bit/pin
        D7Pin int = 0 // GPIO 18  --> Raspberry Physical Pin 12  Data7 bit/pin
)

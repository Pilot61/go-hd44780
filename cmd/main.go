	package main

import (
	hd44780 "github.com/talkkonnect/go-hd44780"
)

var (
        BKPin int = 0
        RSPin int = 7  // GOIO 7   --> Raspberry Physical Pin 26  RS Bit/pin
        EPin int  = 8  // GPIO 8   --> Raspberry Physical Pin 24  Enable bit/pin
        D4Pin int = 25 // GPIO 25  --> Raspberry Physical Pin 22  Data4 bit/pin
        D5Pin int = 24 // GPIO 24  --> Raspberry Physical Pin 18  Data5 bit/pin
        D6Pin int = 23 // GPIO 23  --> Raspberry Physical Pin 16  Data6 bit/pin
        D7Pin int = 18 // GPIO 18  --> Raspberry Physical Pin 12  Data7 bit/pin

	LCDInterfaceType string ="i2c"
	LCDI2CAddress byte
)


func main() {
	      //LcdDisplay(lcdtext_show [4]string, PRSPin int, PEPin int, PD4Pin int, PD5Pin int, PD6Pin int, PD7Pin int)
	var lcdtext = [4]string{"1", "2", "3", "4"}

	LCDI2CAddress = 63
	hd44780.LcdDisplay(lcdtext, RSPin, EPin, D4Pin, D5Pin, D6Pin, D7Pin, LCDInterfaceType, LCDI2CAddress)
}


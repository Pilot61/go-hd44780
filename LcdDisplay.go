package hd44780

import (
	"log"
	"sync"
)

// Variable to Show Last Vale Displayed on LCD
var lcdtext_last = [4]string{"", "", "", ""}
var mutex = &sync.Mutex{}

func LcdDisplay(lcdtext_show [4]string, PRSPin int, PEPin int, PD4Pin int, PD5Pin int, PD6Pin int, PD7Pin int, LCDInterfaceType string, LCDI2CAddress byte) {
	RSPin = PRSPin
	EPin = PEPin
	D4Pin = PD4Pin
	D5Pin = PD5Pin
	D6Pin = PD6Pin
	D7Pin = PD7Pin

	mutex.Lock()
	defer mutex.Unlock()

	if lcdtext_show[0] != "nil" && lcdtext_last[0] != lcdtext_show[0] {
		lcdtext_last[0] = lcdtext_show[0]
	}

	if lcdtext_show[1] != "nil" && lcdtext_last[1] != lcdtext_show[1] {
		lcdtext_last[1] = lcdtext_show[1]
	}

	if lcdtext_show[2] != "nil" && lcdtext_last[2] != lcdtext_show[2] {
		lcdtext_last[2] = lcdtext_show[2]
	}

	if lcdtext_show[3] != "nil" && lcdtext_last[3] != lcdtext_show[3] {
		lcdtext_last[3] = lcdtext_show[3]
	}

	switch LCDInterfaceType {
	case "parallel":
		lcd := NewGPIO4bit()
		if err := lcd.Open(); err != nil {
			log.Println("alert: Can't open lcd: " + err.Error())
			return
		}

		lcd.Display(0, lcdtext_last[0])
		lcd.Display(1, lcdtext_last[1])
		lcd.Display(2, lcdtext_last[2])
		lcd.Display(3, lcdtext_last[3])

	case "i2c":
	        lcd := NewI2C4bit(LCDI2CAddress)
		if err := lcd.Open(); err != nil {
			log.Println("alert: Can't open lcd: " + err.Error())
			return
		}

		lcd.Display(0, lcdtext_last[0])
		lcd.Display(1, lcdtext_last[1])
		lcd.Display(2, lcdtext_last[2])
		lcd.Display(3, lcdtext_last[3])

	default:
		log.Println("alert: Error LCD Interfacing Type Not Defined Correctly in XML Config File!")
		return
	}
}

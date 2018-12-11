package hd44780

import (
	"log"
	"sync"
)

// Variable to Show Last Vale Displayed on LCD
var lcdtext_last = [4]string{"", "", "", ""}
var mutex = &sync.Mutex{}

func LcdDisplay(lcdtext_show [4]string, PRSPin int, PEPin int, PD4Pin int, PD5Pin int, PD6Pin int, PD7Pin int) {
	RSPin = PRSPin
	EPin = PEPin
	D4Pin = PD4Pin
	D5Pin = PD5Pin
	D6Pin = PD6Pin
	D7Pin = PD7Pin

	mutex.Lock()
	defer mutex.Unlock()
	lcd := NewGPIO4bit()
	if err := lcd.Open(); err != nil {
		log.Println("warn: Can't open lcd: " + err.Error())
	}

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

	lcd.Display(0, lcdtext_last[0])
	lcd.Display(1, lcdtext_last[1])
	lcd.Display(2, lcdtext_last[2])
	lcd.Display(3, lcdtext_last[3])
}

package main

import (
	"machine"
	"time"
)

func main() {
	// define and configure all LED pins
	leds := [6]machine.Pin{machine.D2, machine.D4, machine.D6, machine.D8, machine.D10, machine.D12}
	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
		led.Low()
	}

	// setup analog-digital-conversion
	machine.InitADC()
	tempPin := machine.ADC{Pin: machine.ADC5}
	tempPin.Configure(machine.ADCConfig{})
	initialTemp := tempPin.Get()

	// main loop
	for {
		// read temp difference and convert to pin index
		idx := int((tempPin.Get() - initialTemp) / 500)
		println(tempPin.Get())

		// turn off all LEDs
		for _, led := range leds {
			led.Low()
		}

		// turn on all LEDs up to calculated index
		if idx >= len(leds) {
			idx = len(leds) - 1
		}
		for n := 0; n <= idx; n++ {
			leds[n].High()
		}

		// wait a bit
		time.Sleep(time.Millisecond * 500)
	}
}

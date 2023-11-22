package main

import (
	"machine"
	"time"
)

func main() {
	// list and configure all available LED pins
	ledPins := [6]machine.Pin{machine.D2, machine.D4, machine.D6, machine.D8, machine.D10, machine.D12}
	for _, pin := range ledPins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		pin.Low()
	}

	// setup analog-digital-conversion for analog reading
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
		for _, led := range ledPins {
			led.Low()
		}

		// turn on all LEDs up to calculated index
		if idx >= len(ledPins) {
			idx = len(ledPins) - 1
		}
		for n := 0; n <= idx; n++ {
			ledPins[n].High()
		}

		// wait a bit
		time.Sleep(time.Millisecond * 500)
	}
}

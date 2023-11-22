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

	// main loop
	for {
		for n := 0; n < len(ledPins); n++ {
			ledPins[n].High()
			time.Sleep(time.Millisecond * 250)
			ledPins[n].Low()
		}
	}
}

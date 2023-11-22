package main

import (
	"machine"
	"time"
)

func main() {
	// list all available LED pins
	ledPins := [6]machine.Pin{machine.D2, machine.D4, machine.D6, machine.D8, machine.D10, machine.D12}
	for _, pin := range ledPins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		pin.Low()
	}

	// main loop
	for {
		// disable all
		for _, pin := range ledPins {
			pin.Low()
		}
		// sleep
		time.Sleep(time.Second)
		// enable all
		for _, pin := range ledPins {
			pin.High()
		}
		// sleep
		time.Sleep(time.Second * 2)
	}
}

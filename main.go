package main

import (
	"machine"
	"time"
)

func main() {
	leds := [6]machine.Pin{machine.D2, machine.D4, machine.D6, machine.D8, machine.D10, machine.D12}
	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
		led.Low()
	}

	machine.InitADC()
	tempPin := machine.ADC{machine.ADC5}
	tempPin.Configure(machine.ADCConfig{})
	initialTemp := tempPin.Get()

	for {
		idx := int((tempPin.Get() - initialTemp) / 500)
		println(tempPin.Get())
		for _, led := range leds {
			led.Low()
		}
		if idx >= len(leds) {
			idx = len(leds) - 1
		}
		for n := 0; n <= idx; n++ {
			leds[n].High()
		}
		time.Sleep(time.Millisecond * 500)
	}

	//for {
	//	led.Low()
	//	time.Sleep(time.Millisecond * 500)
	//
	//	led.High()
	//	time.Sleep(time.Millisecond * 500)
	//}
}

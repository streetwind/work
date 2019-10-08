package main

import (
	rpio "github.com/stianeikeland/go-rpio/v4"
)

func main() {
	rpio.Open()
	def rpio.Close()
	pin := rpio.Pin(10)
	pin.Output()
	pin.High()
}

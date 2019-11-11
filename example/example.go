package main

import (
	"github.com/Cryptkeeper/go-lightorama/pkg/lor"
	"github.com/tarm/serial"
	"log"
	"math/rand"
	"time"
)

func main() {
	var cont = &lor.Controller{
		Id: 0x01,
	}

	// Open the serial port used for communications with the unit
	err := cont.OpenPort(&serial.Config{
		Name: "/dev/tty.usbserial-A603LKCU",
		Baud: 19200,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to LOR unit!")

	for range time.Tick(lor.DefaultHeartbeatRate) {
		// Maintain the connection by consistently sending the heartbeat packet
		if err := cont.SendHeartbeat(); err != nil {
			log.Fatal("Lost connection!")
		}

		// Constantly randomize the brightness of channel 1
		if err := cont.SetBrightness(0, rand.Float64()); err != nil {
			log.Fatal("Failed to set brightness!")
		}
	}
}

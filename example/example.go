package main

import (
	"github.com/Cryptkeeper/go-lightorama/pkg/lor"
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	var cont = &lor.Controller{Id: 0x01}

	// Open the serial port used for communications with the unit
	err := cont.OpenPort(&serial.Config{
		Name: "/dev/tty.usbserial-A603LKCU",
		Baud: 19200,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to LOR unit!")

	// TODO: wait for init signal
	time.Sleep(time.Second * 2)

	cont.Fade(0, 0, 1, time.Second*5)

	for range time.Tick(time.Millisecond * 500) {
		if err := cont.SendHeartbeat(); err != nil {
			log.Fatal("Lost connection!")
		}
	}
}

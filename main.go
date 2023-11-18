package main

import (
	"Adapters/conf"
	"Adapters/infra"

	"fmt"

	"github.com/totoval/framework/helpers/zone"
	"go.bug.st/serial"
)

func main() {
	getCommlist2()
	port, baud := conf.GetSim800c()
	if c, err := infra.NewSim800c(port, baud, 5*zone.Second); err == nil {
		gw := infra.New(c)

		gw.Listen()
	}
}

func getCommlist2() {
	ports, _ := serial.GetPortsList()

	fmt.Printf("%#v", ports)
	for _, port := range ports {
		fmt.Printf("Find serial port: %v\n", port)
	}
}

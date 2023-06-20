package main

import (
	"fmt"
	"os"
	"strings"
)

var filePath = "/sys/devices/pci0000:00/0000:00:14.3/PNP0C09:00/VPC2004:00/conservation_mode"

func main() {
	byteMode, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}

	mode := strings.TrimRight(string(byteMode), "\n")

	if mode == "0" {
		err = os.WriteFile(filePath, []byte("1"), 0644)
		fmt.Println("CONSERVATION MODE IS ON")
	} else {
		err = os.WriteFile(filePath, []byte("0"), 0644)
		fmt.Println("CONSERVATION MODE IS OFF")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	cpu := &CPU{
		Brand:         "intel",
		Name:          "i7",
		NumberCores:   4,
		NumberThreads: 100,
		MinGhz:        1.0,
		MaxGhz:        4.0,
	}
	fmt.Printf("CPU: %v\n", cpu)

	// Write
	out, err := proto.Marshal(cpu)
	if err != nil {
		log.Fatalf("Serialisation error: %s", err.Error())
	}
	if ioutil.WriteFile("cpu.bin", out, 0644); err != nil {
		log.Fatalf("Write File Error: %s", err.Error())
	}
	fmt.Println("Write Success")

	// Read
	in, err := ioutil.ReadFile("cpu.bin")
	if err != nil {
		log.Fatalf("Read File Error: %s", err.Error())
	}
	cpu2 := &CPU{}
	err = proto.Unmarshal(in, cpu2)
	if err != nil {
		log.Fatalf("Deserialisation error: %s", err.Error())
	}

	fmt.Println("Read Success")
	fmt.Printf("CPU2: %v\n", cpu2)
}

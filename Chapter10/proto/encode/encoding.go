package main

import (
	"log"

	"github.com/ibiscum/Hands-On-Systems-Programming-with-Go/Chapter10/proto/gen"
	"google.golang.org/protobuf/proto"
)

func main() {
	var char = gen.Character{
		Name:        "George",
		Surname:     "Gammell Angell",
		YearOfBirth: 1834,
		Job:         "professor emeritus",
	}

	out, err := proto.Marshal(&char)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	log.Printf("%+v", out)
}

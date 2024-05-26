package main

import (
	"log"

	"github.com/ibiscum/Hands-On-Systems-Programming-with-Go/Chapter10/proto/gen"
	pb "google.golang.org/protobuf/proto"
)

func main() {
	var char = gen.Character{
		Name:        "George",
		Surname:     "Gammell Angell",
		YearOfBirth: 1834,
		Job:         "professor emeritus",
	}

	in, err := pb.Marshal(&char)
	if err != nil {
		log.Fatalln(err)
	}

	if err := pb.Unmarshal(in, &char); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	log.Printf("%+v", &char)
}

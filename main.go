package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wackGarcia/protocol_buffers/book"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	now := time.Now()

	// fill message Person
	person := &book.Person{
		Id:       1,
		Name:     "Abel",
		LastName: "Garcia",
		Age:      28,
		Phones: []*book.Person_PhoneNumber{
			{Number: "555-4321", Type: book.Person_HOME},
		},
		Email: "abelgarcia38348@gmail.com",
		LastUpdated: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	// write message Person
	data := writeMessage(person)

	// print message Person on []bytes
	fmt.Println("Bytes data: ", data)

	// get byte array into an object we can modify and use
	printPeople(readMessage(data))
}

func writeMessage(person *book.Person) []byte {
	// apply Marshal to message,  we can use it transfer the data across service
	data, err := proto.Marshal(person)

	if err != nil {
		log.Fatalln("Failed to encode address book: ", err)
	}

	return data
}

func readMessage(data []byte) *book.Person {
	person := &book.Person{}

	// aaply Unmarshal to person message, witten before
	if err := proto.Unmarshal(data, person); err != nil {
		log.Fatalln("Failed to decode address book: ", err)
	}

	return person
}

func printPeople(person *book.Person) {
	fmt.Println("=============================")
	fmt.Println("ID: ", person.GetId())
	fmt.Println("Name: ", person.GetName())
	fmt.Println("Last Name: ", person.GetLastName())
	fmt.Println("Age: ", person.GetAge())
	fmt.Println("Email: ", person.GetEmail())
	fmt.Println("Phones: ", person.GetPhones())
}

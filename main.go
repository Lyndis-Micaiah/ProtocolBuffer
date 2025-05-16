package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	"protocolbuffer/pb"
)

func main() {
	// Tạo một đối tượng Person
	person := &pb.Person{
		Name:   "Alice",
		Age:    18,
		Emails: []string{"alice@example.com", "alice@work.com"},
		Address: &pb.Address{
			Street: "Da hood",
			City: "Royal City",
			Country: "Nohr",
		},
		Gender: pb.Gender_FEMALE,
		Phone: []*pb.PhoneNumber{
			{Number: "0703633182", Phonetype: pb.PhoneType_HOME},
			{Number: "0934140632", Phonetype: pb.PhoneType_WORK},
		},
	}

	// Tuần tự hóa thành dữ liệu nhị phân
	data, err := proto.Marshal(person)
	if err != nil {
		log.Printf("Lỗi khi tuần tự hóa: %v", err)
	}

	fmt.Printf("Dữ liệu nhị phân: %v\n", data)
	fmt.Printf("Dữ liệu nhị phân (hex): %x\n", data)

	// Giải tuần tự hóa dữ liệu
	newPerson := &pb.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("Lỗi khi giải tuần tự hóa:", err)
	}

	if !proto.Equal(person, newPerson) {
    log.Fatal("Dữ liệu sau giải tuần tự hóa không khớp")
}
	//Print details of the new person
	fmt.Printf("Dữ liệu sau khi giải tuần tự hóa:\n")
	fmt.Printf("Name: %s\n", newPerson.Name)
	fmt.Printf("Age: %d\n", newPerson.Age)
	fmt.Printf("Emails: \n")
	for i, email := range newPerson.Emails {
		fmt.Printf("Email %d: %s\n", i+1, email)
	}
	fmt.Printf("Address: %s, %s, %s\n", newPerson.Address.Street, newPerson.Address.City, newPerson.Address.Country)
	fmt.Printf("Gender: %s\n", newPerson.Gender)
	fmt.Printf("Phone Numbers:\n")
	for i, phone := range newPerson.Phone {
		fmt.Printf("Phone %d: %s (%s)\n", i+1, phone.Number, phone.Phonetype)
	}
}
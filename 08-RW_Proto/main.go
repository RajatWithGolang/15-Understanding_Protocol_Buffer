package main

import (
	"fmt"
	pb "github.com/Rajat2019/Protocol_Buffer/08-RW_Proto/proto"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	person := getMessage()
	fmt.Println(person)
	if err := WritetFile("person.text", person); err != nil {
		fmt.Println("error Encountered during writing")
	}
	fmt.Println("Data has been written successfully")
	person1 := &pb.Person{}
	ReadFile("person.text", person1)
	fmt.Println(person1)
}

func getMessage() *pb.Person {
	person1 := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123 - 4567 - 888", Type: pb.Person_WORK},
			{Number: "9856256254", Type: pb.Person_MOBILE},
		},
	}
	return person1
}

func WritetFile(fname string, person proto.Message) error {
	out, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	return nil
}

func ReadFile(fname string, person1 proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	err = proto.Unmarshal(in, person1)
	if err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return err
	}
	return nil
}

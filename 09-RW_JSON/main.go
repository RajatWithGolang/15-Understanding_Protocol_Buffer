package main

import (
	"fmt"
	pb "github.com/Rajat2019/Protocol_Buffer/09-RW_JSON/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"log"
)

func main() {
	person := getMessage()
	fmt.Println(person)
	RWJson(person)
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

func RWJson(person proto.Message){
	personString := WritetoJSON(person)
	fmt.Println(personString)

	person2 := &pb.Person{}
	ReadFromJSON(personString, person2)
	fmt.Println("Successfully created proto struct:", person2)
}

func WritetoJSON(person proto.Message) string{
	marshaler := jsonpb.Marshaler{}
	jsonString, err := marshaler.MarshalToString(person)
	 if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return jsonString
}

func ReadFromJSON(jsonString string,person2 proto.Message){
   err := jsonpb.UnmarshalString(jsonString, person2)
   if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the person struct", err)
	}
}


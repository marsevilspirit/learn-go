package main

import (
	"example/main/proto"
	pb "google.golang.org/protobuf/proto"
)

func main() {
	value := &proto.String{Value: "Hello, World!"}
	println(value.GetValue())
	
	data, err := pb.Marshal(value)
	if err != nil {	
		panic(err)
	}

	newValue := &proto.String{}
	err = pb.Unmarshal(data, newValue)
	if err != nil {
		panic(err)
	}

	println(newValue.GetValue())
}

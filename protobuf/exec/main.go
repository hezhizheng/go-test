package main

import (
	"github.com/golang/protobuf/proto"
	"go-test/protobuf"
	"log"
)

func main() {

	test := &protobuf.Student{
		Name: "hzz",
		Male:  true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &protobuf.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}

	log.Println(test.GetName(), newTest.GetName(),data,string(data))
}

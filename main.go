package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-protobuf-demo/msgproto"
	"log"
)

func main() {

	test := msgproto.Phone{
		Type:   100,
		Number: "100",
	}
	data, err := proto.Marshal(&test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data) // [8 100 18 3 49 48 48]

	newTest := &msgproto.Phone{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Printf("%v\n",newTest) // type:100  number:"100"
	if test.GetNumber() != newTest.GetNumber() {
		log.Fatalf("data mismatch %q != %q", test.GetNumber(), newTest.GetNumber())
	}else {
		fmt.Println("==")
	}
}

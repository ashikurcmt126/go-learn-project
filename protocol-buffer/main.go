package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main()  {
	fmt.Println("Ashikur Rahman Rashid")

	ashikur := &Person{
		Name: "Ashikur",
		Age: 24,
		SocialFollwers: &SocialFollwers{
			Twitter: 1400,
			Youtube: 2500,
		},
	}

	data, err := proto.Marshal(ashikur)
	if err != nil {
		log.Fatalf("Marshalling error", err)
	}

	fmt.Println(data)

	newAshikur := &Person{}

	err = proto.Unmarshal(data, newAshikur)
	if err != nil{
		log.Fatal("unmarshalling error: ", err)
	}

	fmt.Println(newAshikur.GetName())
	fmt.Println(newAshikur.GetAge())
	fmt.Println(newAshikur.SocialFollwers.GetTwitter())
	fmt.Println(newAshikur.SocialFollwers.GetYoutube())
}

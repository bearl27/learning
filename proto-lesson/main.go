package main

import (
	"fmt"
	"grpc/pb"
	"log"

	"github.com/golang/protobuf/jsonpb"
	// "github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@test.com",
		Occupation:  pb.Occupation_DESIGNER,
		PhoneNumber: []string{"080-123-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "Hello",
		},
		Birthday: &pb.Date{
			Year:  2003,
			Month: 7,
			Day:   27,
		},
	}

	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Can't serialize", err)
	// }

	// if err := os.WriteFile("test.bin", binData, 0666); err != nil {
	// 	log.Fatalln("Can't write", err)
	// }

	// in, err := os.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Can't read file", err)
	// }

	// readEmployee := &pb.Employee{}
	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Can't deserialize", err)
	// }

	// fmt.Println(readEmployee)

	// JSONに変換
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	// fmt.Println(out)

	// m := protojson.MarshalOptions{
	// 	Multiline: true,
	// }
	// out, err := m.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Can't marshal to JSON", err)
	// }

	// fmt.Println(string(out))

	// JSONから復元
	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal from JSON", err)
	}
	fmt.Println(readEmployee)
}

package main

import (
	"fmt"
	"go-protocol-buffer3/src/enums"
	"go-protocol-buffer3/src/simple"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	writeAndReadFileFromDisk()

	jsonDemo()

	enumExample()

}

func enumExample() {

	em := &enums.EnumMessage{
		Id:           42,
		DayOfTheWeek: enums.DayOfTheWeek_FRIDAY,
	}

	fmt.Println(em)

}

func jsonDemo() {

	sm := &simple.SimpleMessage{
		Id:         12342342,
		Name:       "Teste Name",
		IsSimple:   true,
		SampleList: []int32{1, 2, 3, 4, 5, 5},
	}

	strJson := toJson(sm)

	sm2 := &simple.SimpleMessage{}

	fromJson(strJson, sm2)

	fmt.Println(sm2)

}

func writeAndReadFileFromDisk() {

	sm := &simple.SimpleMessage{
		Id:         12342342,
		Name:       "Teste Name",
		IsSimple:   true,
		SampleList: []int32{1, 2, 3, 4, 5, 5},
	}

	fmt.Println(sm.String())

	writeFile("SimpleTest.bin", sm)

	sm2 := &simple.SimpleMessage{}

	readFile("SimpleTest.bin", sm2)

	fmt.Println(sm2.String())

	fmt.Println("")

	fmt.Println(toJson(sm2))

}

func writeFile(fname string, pb proto.Message) error {

	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Não serializou!", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Erro ao gravar", err)
		return err
	}

	fmt.Println("Arquivo escrito")

	return nil

}

func readFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("erro ao ler", err)
		return err
	}

	err = proto.Unmarshal(in, pb)

	if err != nil {
		log.Fatalln("Nao conseguiu converter", err)
		return err
	}

	fmt.Println("Leitura concluida")

	return nil

}

func toJson(pb proto.Message) string {

	marshaler := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	out, err := marshaler.Marshal(pb)

	if err != nil {
		log.Fatalln("Erro ao converter", err)
		return ""
	}

	return string(out)

}

func fromJson(json string, pb proto.Message) {

	unmarshaler := protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: true,
	}

	err := unmarshaler.Unmarshal([]byte(json), pb)

	if err != nil {
		log.Fatalln("Erro ao recuperar do json", err)
	}

}

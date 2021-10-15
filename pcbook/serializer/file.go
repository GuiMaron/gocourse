package serializer

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
)

func ReadProtobufFromBinaryFile (filename string, message proto.Message) error {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("cannot read binary file => %w", err)
	}

	err = proto.Unmarshal(data ,message)

	if err != nil {
		return fmt.Errorf("cannot unmarshal (nothing else matters) binary to message => %w", err)
	}

	return nil

}

func WriteProtobufToBinaryFile (message proto.Message, filename string) error {

	data, err := proto.Marshal(message)

	if err != nil {
		return fmt.Errorf("cannot marshal (matters) message to binary => %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0666)

	if err != nil {
		return fmt.Errorf("cannot write binary file => %w", err)
	}

	return nil

}

/*
func WriteProtobufToJSONFile (message proto.Message, filename string) error {

	data, err := ProtobufToJSON(message)

	if err != nil {
		return fmt.Errorf("cannot marshal (matters) message to JSON => %w", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0666)

	if err != nil {
		return fmt.Errorf("cannot write JSON file => %w", err)
	}

	return nil

}
 */

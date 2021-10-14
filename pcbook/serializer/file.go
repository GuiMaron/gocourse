package serializer

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
)

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

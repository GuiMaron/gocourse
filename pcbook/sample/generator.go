package sample

import "github.com/GuiMaron/gocourse/pcbook/pb"

func NewKeyboard() *pb.Keyboard {

	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard

}

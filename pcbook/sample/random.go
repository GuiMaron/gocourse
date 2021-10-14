package sample

import (
	"github.com/GuiMaron/gocourse/pcbook/pb"
	"math/rand"
	"time"
)

/**
 *	General
 */

func getRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomBool() bool {
	return (getRand().Intn(2) % 2) == 0
}

/**
 * Keyboard
 */
func randomKeyboardLayout() *pb.Keyboard {

	switch getRand().Intn(5) {

	case 1:
		return pb.Keyboard_QWERTY

	}

}

package sample

import (
	pcbook "github.com/GuiMaron/gocourse/pcbook/pb"
	"math"
	"math/rand"
	"time"
)

/**
 *	General
 */

func getRand () *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomBool () bool {
	return (getRand().Intn(2) % 2) == 0
}

func randomFloat32 (min float32, max float32) float32 {

	min = float32(math.Min(float64(min),  float64(max)))
	max = float32(math.Max(float64(min),  float64(max)))

	return min + (getRand().Float32() * (max - min))

}

func randomInt (min int, max int) int {

	min = int(math.Min(float64(min),  float64(max)))
	max = int(math.Max(float64(min),  float64(max)))

	return min + getRand().Intn(max - min + 1)

}

func randomStringFromSet (strings ... string) string {

	length := len(strings)

	if 0 == length {
		return ""
	}

	return strings[getRand().Intn(length)]

}

/**
 *	CPU
 */
func randomCPUBrand () string {
	return randomStringFromSet("INTEL", "AMD")
}

func randomCPUName (brand string) string {

	switch brand {

	case "Intel":

		return randomStringFromSet(
			"Xeon E-2286M",
					"Core i9-9980HK",
					"Core i7-9750H",
					"Core i5-9400F",
					"Core i3-1005G1",
				)

	default:

		return randomStringFromSet(
			"Ryzen 7 PRO 2700U",
					"Ryzen 5 PRO 3500U",
					"Ryzen 3 PRO 3200GE",
				)

	}

}

/**
 *	GPU
 */
func randomGPUBrand () string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName (brand string) string {

	switch brand {

	case "NVIDIA":

		return randomStringFromSet(
			"RTX 2060",
					"RTX 2070",
					"GTX 1660-Ti",
					"GTX 1070",
				)

	default:

		return randomStringFromSet(
			"RX 590",
					"RX 580",
					"RX 5700-XT",
					"RX Vega-56",
				)

	}

}

/**
 * 	Keyboard
 */
func randomKeyboardLayout () pcbook.Keyboard_Layout {

	switch getRand().Intn(4) {

	case 1:
		return pcbook.Keyboard_QWERTY
	case 2:
		return pcbook.Keyboard_QWERTZ
	case 3:
		return pcbook.Keyboard_AZERTY
	default:
		return pcbook.Keyboard_DVORAK

	}

}

package sample

import (
	"github.com/GuiMaron/gocourse/pcbook/pb"
	"github.com/google/uuid"
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

func randomFloat64 (min float64, max float64) float64 {

	min = math.Min(min, max)
	max = math.Max(min, max)

	return min + (getRand().Float64() * (max - min))

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
 *	Laptop
 */
func randomId () string {
	return uuid.New().String()
}

func randomLaptopBrand () string {
	return randomStringFromSet("APPLE", "DELL", "LENOVO")
}

func randomLaptopName (brand string) string {

	switch brand {

	case "APPLE":

		return randomStringFromSet(
			"Macbook Air",
					"Macbook Pro",
				)

	case "DELL":

		return randomStringFromSet(
			"Latitude",
					"Vostro",
					"XPS",
					"Alienware",
				)

	default:

		return randomStringFromSet(
			"Thinkpad X1",
					"Thinkpad P1",
					"Thinkpad P53",
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

/**
 *	Screen
 */
func randomScreenPanel () pcbook.Screen_Panel {

	switch getRand().Intn(3) {

	case 1:
		return pcbook.Screen_IPS
	case 2:
		return pcbook.Screen_OLED
	default:
		return pcbook.Screen_LED

	}

}

func randomScreenResolution () *pcbook.Screen_Resolution {

	height	:= randomInt(1080, 4320)
	width	:= height * (16 / 9)

	resolution := &pcbook.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

	return resolution

}

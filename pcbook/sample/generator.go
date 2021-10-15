package sample

import (
	pcbook "github.com/GuiMaron/gocourse/pcbook/pb/proto"
	"github.com/golang/protobuf/ptypes"
)

func NewCPU () *pcbook.CPU {

	brand 		:= randomCPUBrand()
	name  		:= randomCPUName(brand)
	cores		:= randomInt(2, 8)
	threads		:= randomInt(cores, (cores * 2))
	minClock	:= randomFloat32(2.0, 3.5)
	maxClock	:= randomFloat32(minClock, 5.0)

	cpu := &pcbook.CPU {
		Brand:		brand,
		Name:		name,
		Cores:		uint32(cores),
		Threads:	uint32(threads),
		MinClock: 	minClock,
		MaxClock: 	maxClock,
	}

	return cpu

}

func NewGPU () *pcbook.GPU {

	brand 		:= randomCPUBrand()
	name  		:= randomGPUName(brand)
	minClock	:= randomFloat32(1.0, 1.5)
	maxClock	:= randomFloat32(minClock, 2.0)

	memory		:= &pcbook.Memory {
		Value:	uint64(randomInt(2, 6)),
		Unit: 	pcbook.Memory_GIGABYTE,
	}

	gpu := &pcbook.GPU {
		Brand:		brand,
		Name:		name,
		MinClock: 	minClock,
		MaxClock: 	maxClock,
		Memory: 	memory,
	}

	return gpu

}

func NewHDD () *pcbook.Storage {

	hdd	:= &pcbook.Storage {
		Driver: pcbook.Storage_HDD,
		Memory: &pcbook.Memory {
			Value: 	uint64(randomInt(1, 6)),
			Unit:	pcbook.Memory_TERABYTE,
		},
	}

	return hdd

}

func NewKeyboard () *pcbook.Keyboard {

	keyboard := &pcbook.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard

}

func NewLaptop () *pcbook.Laptop {

	brand	:= randomLaptopBrand()
	name	:= randomLaptopName(brand)

	laptop := &pcbook.Laptop {
		Id:				randomId(),
		Brand:			brand,
		Name:			name,
		Cpu:			NewCPU(),
		Ram:			NewRam(),
		Screen: 		NewScreen(),
		Keyboard:		NewKeyboard(),
		Gpus:			[]*pcbook.GPU{ NewGPU() },
		Storages:  		[]*pcbook.Storage{ NewSSD(), NewHDD() },
		Weight:			&pcbook.Laptop_WeightKg {
			WeightKg:	randomFloat32(1.0, 3.0),
		},
		PriceUsd: 		randomFloat32(1500, 3000),
		ReleaseYear:	uint32(randomInt(2015, 2020)),
		UpdatedAt: 		ptypes.TimestampNow(),
	}

	return laptop

}

func NewRam () *pcbook.Memory {

	ram	:= &pcbook.Memory {
		Value: 	uint64(randomInt(4, 64)),
		Unit:	pcbook.Memory_GIGABYTE,
	}

	return ram

}

func NewScreen () *pcbook.Screen {

	screen	:= &pcbook.Screen {
		SizeInches:	randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:		randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen

}

func NewSSD () *pcbook.Storage {

	ssd	:= &pcbook.Storage {
		Driver: pcbook.Storage_SSD,
		Memory: &pcbook.Memory {
			Value: 	uint64(randomInt(128, 1024)),
			Unit:	pcbook.Memory_GIGABYTE,
		},
	}

	return ssd

}

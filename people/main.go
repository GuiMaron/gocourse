package main

import "fmt"

type contactInfo struct {
	phone	string
	zip		int
}

type person struct {
	firstName	string
	lastName  	string
	contact		contactInfo
}

func main () {

	ci 		:= contactInfo	{ phone: "+46 (073) 410-1679", zip: 41136 }
	maron5	:= person 		{ "Guilherme", "Maron", ci }

	maron5_2 := person {
		firstName:	"Guilherme",
		lastName: 	"Maron",
		contact:	ci,
	}

	var maron5_3 person

	maron5_3.firstName	= "Guilherme"
	maron5_3.lastName	= "Maron"
	maron5_3.contact	= ci

	fmt.Println(maron5)
	fmt.Println(maron5_2)
	fmt.Println(maron5_3)

	fmt.Printf("\n%+v\n", maron5)
	fmt.Printf("%+v\n", maron5_2)
	fmt.Printf("%+v\n", maron5_3)

	maron5_4 := person {
		firstName: "Guilherme",
		lastName:	"Maron",
		contact:	contactInfo {
			phone:	"+46 (073) 410-1679",
			zip:	41136,
		},
	}

	fmt.Println(maron5_4)
	fmt.Printf("%+v\n", maron5_4)

	fmt.Println("\nmaron5.print()")
	maron5.print()

	fmt.Println("\nUpdating (FAIL")
	maron5.updateFirstName("Ramon")
	maron5.print()

	fmt.Println("\nUpdating with Pointer")
	maron5Pointer := &maron5
	maron5Pointer.realUpdateFirstName("Ramon")
	maron5.print()

	fmt.Println("\nUpdating with Address (&)")
	(& maron5).realUpdateFirstName("Roman")
	maron5.print()
	fmt.Printf("\n%x", & maron5)

}

func (p person) print () {
	fmt.Printf("%+v\n",p)
}

func (p person) updateFirstName (newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) realUpdateFirstName (newFirstName string) {
	(* p).firstName = newFirstName
}

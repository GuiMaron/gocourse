package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//	The separator
var CARD_SEPARATOR string = ";"

var NUMBERS = []string{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

var SUITS = []string{
	"Clubs",
	"Diamonds",
	"Hearts",
	"Spades",
}

//	Type definition
type deck []string

func deal(d deck, n int) (deck, deck) {

	n = int(math.Max(float64(0), math.Min(float64(n), float64(d.length()))))

	return d[:n], d[n:]

}

func (d deck) equals(o deck) bool {

	if d.length() != o.length() {
		return false
	}

	for i, c := range d {
		if c != o[i] {
			return false
		}
	}

	return true

}

func (d deck) filter(s string) (deck, []int) {

	cards := deck{}
	var indexes []int

	for i, c := range d {

		regExp, e := regexp.MatchString(s, c)

		if regExp && e == nil {
			cards = append(cards, c)
			indexes = append(indexes, i)
		}
	}

	return cards, indexes

}

func (d deck) length() int {
	return len(d)
}

func newDeck() deck {

	cards := deck{}

	for _, value := range NUMBERS {
		for _, suit := range SUITS {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards

}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i+1, c)
	}
}

func readFromFile(fn string) (deck, error) {

	cards := deck{}
	data, e := ioutil.ReadFile(fn)

	if e == nil {
		cards = strings.Split(string(data), CARD_SEPARATOR)
	}

	return cards, e

}

func (d deck) saveToFile(fn string) error {
	return ioutil.WriteFile(fn, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		j := random.Intn(d.length())
		d[i], d[j] = d[j], d[i]
	}

}

func (d deck) toString() string {
	return strings.Join(d, CARD_SEPARATOR)
}

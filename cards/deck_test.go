package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

//	Name of the file cards will be saved and retrieved
var FILENAME string = "test_cards.txt"
// MAximum number of tryes in loops
var MAX_TRYES int = 5

//	Logs what is beeing tested
func logTest(fn string, decorated bool) {

	logText := ""

	if decorated {
		logText = fmt.Sprintf("| Testing %s() |", fn)
	} else {
		logText = fmt.Sprintf("| %s |", fn)
	}

	line := fmt.Sprintf("+%s+", strings.Repeat("-", (len(logText)-2)))

	fmt.Printf("\n\n%s\n%s\n%s\n\n", line, logText, line)

}

//	Tests if the deck is in the original ordering
func isOrderedDeck(cards deck, t *testing.T) bool {

	suits := len(SUITS)
	isOrdered := true

	// All numbers in the right places?
	for numberIndex, number := range NUMBERS {

		_, filteredCardsNumberIndexes := cards.filter(number)

		for fcni, fcn := range filteredCardsNumberIndexes {

			expectedfcn := (numberIndex * suits) + fcni

			if fcn != expectedfcn {
				t.Errorf("Number (%s) out of place => expected:%d | got: %d", number, expectedfcn, fcn)
				isOrdered = false
			}

		}

	}

	// All suits in the right places?
	for suitIndex, suit := range SUITS {

		_, filteredCardsSuitIndexes := cards.filter(suit)

		for fcsi, fcs := range filteredCardsSuitIndexes {

			expectedfcs := (suits * fcsi) + suitIndex

			if fcs != expectedfcs {
				t.Errorf("Suit (%s) out of place => expected: %d | got: %d", suit, expectedfcs, fcs)
				isOrdered = false
			}

		}

	}

	return isOrdered

}

func isValidDeck(cards deck, t *testing.T) bool {

	isValid := true

	//	Deck length
	if cards.length() != 52 {
		t.Errorf("Deck length => expected: 52 | got: %d", cards.length())
		isValid = false
	}

	// 4 of each number
	for _, n := range NUMBERS {

		fcn, _ := cards.filter(n)

		if fcn.length() != len(SUITS) {
			t.Errorf("Wrong number of %s => expected: 4 | got: %d", n, fcn.length())
			isValid = false
		}

	}

	// 13 of each suit
	for _, s := range SUITS {

		fcs, _ := cards.filter(s)

		if fcs.length() != len(NUMBERS) {
			t.Errorf("Wrong number of %s => expected: 13 | got: %d", s, fcs.length())
			isValid = false
		}

	}

	// No repeated cards
	for i, c := range cards {

		fcc, _ := cards[i+1:].filter(c)

		if fcc.length() != 0 {
			t.Errorf("Duplicated card found => %s", c)
			isValid = false
		}

	}

	return isValid

}

func testDeal (	cards deck,
				expectedHand deck,
				expectedRemainingCards deck,
				handSize int,
				t *testing.T,
) (bool, deck, deck) {

	hand := deck{}
	hand, cards = deal(cards, handSize)

	if ! expectedHand.equals(hand) {
		t.Error("Wrong hand dealt, expected:")
		t.Error(expectedHand)
		t.Error("got:")
		t.Error(hand)
		return false, hand, cards
	}

	if ! expectedRemainingCards.equals(cards) {
		t.Error("Wrong remaining cards, expected:")
		t.Error(expectedRemainingCards)
		t.Error("got:")
		t.Error(cards)
		return false, hand, cards
	}

	return true, hand, cards

}

func TestFilter (t *testing.T) {

	logTest("Deck.filter", true)

	cards := deck {
		"Ace of Clubs",
		"Ace of Diamonds",
		"Ace of Hearts",
		"Ace of Spades",
		"Two of Clubs",
		"Two of Diamonds",
		"Two of Hearts",
		"Two of Spades",
	}

	spadesOnly := deck {
		cards[3],
		cards[cards.length() - 1],
	}

	//	Filtering all Aces
	fmt.Println("Filtering all Aces")
	aces, ai := cards.filter("Ace")

	if (aces.length() != 4) || (! cards[0:4].equals(aces)) {
		t.Errorf("Error on filtering all Aces (cards)\nExpected:\t[%s] \nGot:\t\t[%s]", cards[:4].toString(), aces.toString())
	}

	if (ai[0] != 0) || (ai[1] != 1) || (ai[2] != 2) || (ai[3] != 3) {
		t.Error("Error on filtering all Aces (indexes)")
		t.Error("Expected: ")
		t.Error([]int{0, 1, 2, 3})
		t.Error("Got: ")
		t.Error(ai)
	}

	//	Filtering all Spades
	fmt.Println("Filtering all Spades - up to Two")
	spades, si := cards.filter("Spades")

	if (spades.length() != 2) || (! spadesOnly.equals(spades)) {
		t.Errorf("Error on filtering all Spades up to Two (cards)\nExpected:\t[%s] \nGot:\t\t[%s]", spadesOnly.toString(), spades.toString())
	}

	if (si[0] != 3) || (si[1] != (cards.length() - 1))  {
		t.Error("Error on filtering all Spades up to Two (indexes)")
		t.Error("Expected: ")
		t.Error([]int{3, (cards.length() - 1)})
		t.Error("Got: ")
		t.Error(si)
	}

	//	Filtering a card taht is not here
	fmt.Println("Filtering all Fours (not in deck)")
	fours, fi := cards.filter("Four")

	if (fours.length() != 0) || (! deck{}.equals(fours)) {
		t.Errorf("Error on filtering all fours (not in deck)\nExpected:\t[%s] \nGot:\t\t[%s]", deck{}.toString(), fours.toString())
	}

	if len(fi) != 0  {
		t.Error("Error on filtering all Fours (not in deck) (indexes)")
		t.Error("Expected: ")
		t.Error([]int{})
		t.Error("Got: ")
		t.Error(fi)
	}

}

func TestNewDeck(t *testing.T) {

	logTest("newDeck", true)

	cards := newDeck()

	if ! isValidDeck(cards, t) {
		t.Fatalf("Invalid Deck.")
	}

	// The intended order
	for i := range make([]int, MAX_TRYES) {

		fmt.Printf("Testing the order of the deck => try: %d / %d\n", (i + 1), MAX_TRYES)

		if !isOrderedDeck(cards, t) {
			t.Errorf("Deck is not in the intended order.")
		}

		cards = newDeck()

	}

}

func TestLength (t *testing.T) {

	logTest("Deck.length", true)

	cards := newDeck()

	if cards.length() != len(cards) {
		t.Errorf("Deck.length() => %d is different from len(Deck) => %d", cards.length(), len(cards))
	}

}

func TestToString (t *testing.T) {

	logTest("Deck.toString", true)

	aces := deck{
		"Ace of Clubs",
		"Ace of Diamonds",
		"Ace of Hearts",
		"Ace of Spades",
	}

	acesString := strings.Join(aces, CARD_SEPARATOR)

	cards := newDeck()[:4]

	if cards.toString() != acesString {
		t.Errorf("Deck.toString() failed (Aces)\nExpected:\t%s\ngot:\t\t%s", acesString, cards.toString())
	}

}

func TestPrint (t *testing.T) {

	logTest("Deck.print", true)
	cards := newDeck()
	cards.print()

}

func TestEquals(t *testing.T) {

	logTest("Deck.equals", true)

	d1 := newDeck()
	d2 := newDeck()

	if !d1.equals(d2) {
		t.Fatalf("Deck.equals failed")
	}

}

func TestShuffle(t *testing.T) {

	logTest("Deck.shuffle", true)

	cards := newDeck()
	original := newDeck()
	reversed := newDeck()

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	// 5 tryes, it'll fail if it shuffles the deck 5 ways equally, in the original order
	for trial := range make([]int, MAX_TRYES) {

		cards.shuffle()

		// It is still valid?
		if !isValidDeck(cards, t) {
			t.Error("Shuffling invalidated the deck.")
		}

		fmt.Printf("Shuffling => try: %d / %d\n", (trial + 1), MAX_TRYES) // IDK how to loop without warning of unused variable

		// It is different and not only reversed
		if (!reversed.equals(cards)) && (!original.equals(cards)) {
			return
		}

	}

	t.Errorf("Tried shuffling %d times, got same deck.", MAX_TRYES)

}

func TestSaveToFile (t *testing.T)  {

	logTest("Deck.saveToFile", true)
	os.Remove(FILENAME)

	cards := newDeck()

	//	Saved
	we := cards.saveToFile(FILENAME)

	if we != nil {
		t.Error("Error on writing file")
		t.Error(we)
	}

	// Recovering
	data, re := ioutil.ReadFile(FILENAME)

	if re != nil {
		t.Error("Error on reading file")
		t.Error(re)
	}

	newDeck := strings.Split(string(data), CARD_SEPARATOR)

	if ! cards.equals(newDeck) {
		t.Error("Deck read from file is different from the original one (written)")
	}

	os.Remove(FILENAME)

}

func TestReadFromFile (t *testing.T) {

	logTest("Deck.readFromFile", true)
	os.Remove(FILENAME)

	cards := newDeck()

	//	Saved
	we := cards.saveToFile(FILENAME)

	if we != nil {
		t.Error("Error on writing file")
		t.Error(we)
	}

	//	Now test reading
	newDeck, re := readFromFile(FILENAME)

	if re != nil {
		t.Error("Error on reading file")
		t.Error(re)
	}

	if ! cards.equals(newDeck) {
		t.Error("Deck read from file is different from the original one (written)")
	}

	os.Remove(FILENAME)

}

func TestDeal (t *testing.T) {

	logTest("deal", true)

	cards := newDeck()
	dealt := false
	expectedHand := cards[:4]
	expectedCards := cards[4:]
	hand  := deck{}
	message := ""
	newHand := deck{}
	startingHangSize := 4



	//	Dealing all Aces
	message = fmt.Sprintf("Dealing %d cards from %d hand\n", startingHangSize, cards.length())

	fmt.Printf(message)
	dealt, hand, cards = testDeal(cards, expectedHand, expectedCards, startingHangSize, t)

	if ! dealt {
		t.Errorf("Error on %s", message)
		t.Errorf("Hand => expected: %d | got %d \n cards reamining => expected: %d | got: %d", startingHangSize, hand.length(), expectedCards.length(), expectedCards.length())
	}



	//	Dealing 5 cards from 4 hand
	//	Trying do deal more cards than I have (have 4, try do deal 5)
	message = fmt.Sprintf("Dealing %d cards from %d hand\n", (startingHangSize + 1), hand.length())

	fmt.Printf(message)
	dealt, newHand, hand = testDeal(hand, expectedHand, deck{}, (startingHangSize + 1), t)

	if ! dealt {
		t.Errorf("Error on %s", message)
		t.Errorf("Hand => expected: %d | got %d \n cards reamining => expected: %d | got: %d", startingHangSize, newHand.length(), 0, hand.length())
	}



	//	Dealing -1 card from 4 hand
	message = fmt.Sprintf("Dealing %d cards from %d hand\n", -1, newHand.length())

	fmt.Printf(message)
	dealt, newHand, hand = testDeal(newHand, deck{}, newHand, -1, t)

	if ! dealt {
		t.Errorf("Error on %s", message)
		t.Errorf("Hand => expected: %d | got %d \n cards reamining => expected: %d | got: %d", 0, newHand.length(), expectedHand.length(), hand.length())
	}



	//	Dealing 1 card from 0 hand
	message = fmt.Sprintf("Dealing %d cards from %d hand\n", 1, deck{}.length())

	fmt.Printf(message)
	dealt, newHand, hand = testDeal(deck{}, deck{}, deck{}, 1, t)

	if ! dealt {
		t.Errorf("Error on %s", message)
		t.Errorf("Hand => expected: %d | got %d \n cards reamining => expected: %d | got: %d", 0, newHand.length(), 0, hand.length())
	}

}

func Test_ENDOFTEST(t *testing.T) {
	logTest("END OF TEST", false)
}

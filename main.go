package main

import (
	"fmt"
)

type Item struct {
	Quality int
	Name    string
	SellIn  int
}

type funcToTest func(Item) Item

func TestFunction(description string, funcToTest funcToTest, input, expected Item) {
	actual := funcToTest(input)
	passed := actual == expected
	if passed {
		fmt.Printf("Test %s Passed\n", description)
	} else {
		fmt.Printf("Test %s failed\nExpected %v+ to equal %v+", description, actual, expected)
	}
}

func TestExecise(description string, quality int, name string, sellIn int, expectedQuality int) {
	input := Item{
		Quality: quality,
		Name:    name,
		SellIn:  sellIn,
	}
	expected := Item{
		Quality: expectedQuality,
		Name:    name,
		SellIn:  sellIn,
	}
	TestFunction(description, Excercise, input, expected)
}

func qualityIsLessThan50(item Item) bool {
	return item.Quality < 50
}

func sellInLessThanN(item Item, n int) bool {
	return item.SellIn < n
}

func isNameBackstage(item Item) bool {
	return item.Name == "backstage"
}

func bumpQuality(item *Item) {
	item.Quality += 1
}
func Excercise(item Item) Item {
	if qualityIsLessThan50(item) && isNameBackstage(item) {
		if sellInLessThanN(item, 11) {
			bumpQuality(&item)
		}
		if sellInLessThanN(item, 6) && qualityIsLessThan50(item) {
			bumpQuality(&item)
		}
	}
	return item
}

func main() {

	TestExecise("happy Test SellIn less than 6", 48, "backstage", 5, 50)
	TestExecise("happy Test SellIn less than 11", 48, "backstage", 10, 49)
	TestExecise("quality is 49 (bump only once)", 49, "backstage", 5, 50)
	TestExecise("SellIn greater than 11", 48, "backstage", 12, 48)
	TestExecise("greater than 50 Quality input", 51, "backstage", 10, 51)
	TestExecise("name not backstage input", 48, "banana", 10, 48)
}

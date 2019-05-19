package search

import (
	"log"
	"reflect"
)

var dummyFeed = &Feed{}

func ExampleDefault01() {
	dm := new(defaultMatcher)
	log.Println("dm:", reflect.TypeOf(dm))
	dm.SearchType(dummyFeed, "test")
	// Output:
	// ?
}

func ExampleDefault02() {
	var dm defaultMatcher
	log.Println("dm:", reflect.TypeOf(dm))
	dm.SearchType(dummyFeed, "test")
	// Output:
	// ?
}

func ExampleDefault03() {
	dm := new(defaultMatcher)
	var matcher TypeMatcher = dm
	log.Println("dm:", reflect.TypeOf(dm))
	matcher.SearchType(dummyFeed, "test")
	// Output:
	// ?
}

// func ExampleDefault04() {
// 	var dm defaultMatcher
// 	var matcher TypeMatcher = dm
// 	log.Println("dm:", reflect.TypeOf(dm))
// 	matcher.SearchType(dummyFeed, "test")
// 	// Output:
// 	// ?
// }

func ExampleDefault05() {
	var dm defaultMatcher
	var matcher TypeMatcher = &dm
	log.Println("dm:", reflect.TypeOf(dm))
	matcher.SearchType(dummyFeed, "test")
	// Output:
	// ?
}

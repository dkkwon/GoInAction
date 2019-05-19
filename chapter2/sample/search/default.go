package search

import "log"

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	log.Println("DummySearch")
	return nil, nil
}

type TypeMatcher interface {
	SearchType(feed *Feed, searchTerm string) ([]*Result, error)
}

func (m *defaultMatcher) SearchType(feed *Feed, searchTerm string) ([]*Result, error) {
	log.Println("SearchType")
	return nil, nil
}

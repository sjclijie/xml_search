package matchers

import (
	"errors"
	"search"
)

type defaultMatcher struct {
}

func init() {

	var matcher defaultMatcher

	search.Register("default", matcher)
}

func (matcher defaultMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	return nil, errors.New("default matcher")
}

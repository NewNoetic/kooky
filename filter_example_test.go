package kooky_test

import (
	"fmt"
	"regexp"

	"github.com/newnoetic/kooky"
)

// example regex matching base64 strings
var reBase64 = regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$`)

func ExampleFilter_regex() {
	var cookies = []*kooky.Cookie{{Name: `test`, Value: `dGVzdA==`}}

	cookies = kooky.FilterCookies(
		cookies,
		ValueRegexMatch(reBase64), // filter cookies with the regex filter
		// kooky.Debug,            // print cookies after applying the regex filter
	)

	for _, cookie := range cookies {
		fmt.Println(cookie.Value)
		break // only first element
	}

	// Output: dGVzdA==
}

func ValueRegexMatch(re *regexp.Regexp) kooky.Filter {
	return func(cookie *kooky.Cookie) bool {
		return cookie != nil && re != nil && re.Match([]byte(cookie.Value))
	}
}

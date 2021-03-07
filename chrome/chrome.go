package chrome

import (
	"runtime"

	"github.com/newnoetic/kooky"
	"github.com/newnoetic/kooky/internal/chrome"
)

func ReadCookies(filename string, filters ...kooky.Filter) ([]*kooky.Cookie, error) {
	s := &chrome.CookieStore{}
	s.FileNameStr = filename

	// TODO
	switch runtime.GOOS {
	case `windows`, `darwin`:
		s.BrowserStr = `chrome`
	default:
		s.BrowserStr = `chromium`
	}

	defer s.Close()

	return s.ReadCookies(filters...)
}

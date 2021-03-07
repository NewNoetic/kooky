package firefox

import (
	"github.com/newnoetic/kooky"
	"github.com/newnoetic/kooky/internal"
	"github.com/newnoetic/kooky/internal/firefox"
	"github.com/newnoetic/kooky/internal/firefox/find"
)

type firefoxFinder struct{}

var _ kooky.CookieStoreFinder = (*firefoxFinder)(nil)

func init() {
	kooky.RegisterFinder(`firefox`, &firefoxFinder{})
}

func (s *firefoxFinder) FindCookieStores() ([]kooky.CookieStore, error) {
	files, err := find.FindFirefoxCookieStoreFiles()
	if err != nil {
		return nil, err
	}

	var ret []kooky.CookieStore
	for _, file := range files {
		ret = append(
			ret,
			&firefox.CookieStore{
				DefaultCookieStore: internal.DefaultCookieStore{
					BrowserStr:           file.Browser,
					ProfileStr:           file.Profile,
					IsDefaultProfileBool: file.IsDefaultProfile,
					FileNameStr:          file.Path,
				},
			},
		)
	}

	return ret, nil
}

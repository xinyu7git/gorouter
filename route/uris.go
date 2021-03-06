package route

import (
	"errors"
	"strings"
)

type Uri string

func (u Uri) ToLower() Uri {
	return Uri(strings.ToLower(string(u)))
}

func (u Uri) NextWildcard() (Uri, error) {
	uri := strings.TrimPrefix(string(u), "*.")

	i := strings.Index(uri, ".")
	if i == -1 {
		return u, errors.New("no next wildcard available")
	}
	suffix := uri[i+1:]
	return Uri("*." + suffix), nil
}

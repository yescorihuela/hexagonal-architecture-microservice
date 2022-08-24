package entity

import "strings"

type URLImage struct {
	Url string
}

func (u *URLImage) IsValid() bool {
	return strings.TrimSpace(u.Url) != ""
}

package helper

import "github.com/gosimple/slug"

func Slug(s string) string {
	return slug.Make(s)
}

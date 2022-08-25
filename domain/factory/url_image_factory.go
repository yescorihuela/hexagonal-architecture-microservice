package factory

import (
	"regexp"

	"github.com/yescorihuela/agrak/domain/entity"
)

func NewURLImage(url string) *entity.URLImage {
	regex := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	if !regex.MatchString(url) {
		return &entity.URLImage{}
	}
	return &entity.URLImage{
		Url: url,
	}
}

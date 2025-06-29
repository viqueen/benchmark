package helpers

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

type ContentFixture struct {
	Title string
	Body  string
}

func NewValidContentFixture() ContentFixture {
	return ContentFixture{
		Title: sentence(140),
		Body:  faker.Paragraph(),
	}
}

func sentence(maxLength int) string {
	if maxLength <= 0 {
		return ""
	}
	return faker.Paragraph(func(opts *options.Options) {
		opts.RandomStringLength = maxLength
	})
}

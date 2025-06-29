package api

import (
	"api/go-sdk/content/v1/contentV1connect"
)

type service struct{}

func NewContentService() contentV1connect.ContentServiceHandler {
	return &service{}
}

package helpers

import (
	contentV1 "api/go-sdk/content/v1"
	"api/go-sdk/content/v1/contentV1connect"
	"connectrpc.com/connect"
	"net/http"
	"time"
)

func NewContentClient(baseURL string) contentV1connect.ContentServiceClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	return contentV1connect.NewContentServiceClient(httpClient, baseURL)
}

func NewCreateContentRequest(title, body string) *connect.Request[contentV1.CreateContentRequest] {
	return connect.NewRequest(&contentV1.CreateContentRequest{
		Title: title,
		Body:  body,
	})
}

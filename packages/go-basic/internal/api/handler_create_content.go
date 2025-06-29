package api

import (
	contentV1 "api/go-sdk/content/v1"
	"connectrpc.com/connect"
	"context"
)

func (s service) CreateContent(ctx context.Context, c *connect.Request[contentV1.CreateContentRequest]) (*connect.Response[contentV1.CreateContentResponse], error) {
	//TODO implement me
	panic("implement me")
}

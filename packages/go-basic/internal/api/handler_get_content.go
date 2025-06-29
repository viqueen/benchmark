package api

import (
	contentV1 "api/go-sdk/content/v1"
	"connectrpc.com/connect"
	"context"
)

func (s service) GetContent(ctx context.Context, c *connect.Request[contentV1.GetContentRequest]) (*connect.Response[contentV1.GetContentResponse], error) {
	//TODO implement me
	panic("implement me")
}

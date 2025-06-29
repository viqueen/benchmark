package api

import (
	contentV1 "api/go-sdk/content/v1"
	"connectrpc.com/connect"
	"context"
)

func (s service) UpdateContent(ctx context.Context, c *connect.Request[contentV1.UpdateContentRequest]) (*connect.Response[contentV1.UpdateContentResponse], error) {
	//TODO implement me
	panic("implement me")
}

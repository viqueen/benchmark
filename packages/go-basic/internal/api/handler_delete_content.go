package api

import (
	contentV1 "api/go-sdk/content/v1"
	"connectrpc.com/connect"
	"context"
)

func (s service) DeleteContent(ctx context.Context, c *connect.Request[contentV1.DeleteContentRequest]) (*connect.Response[contentV1.DeleteContentResponse], error) {
	//TODO implement me
	panic("implement me")
}

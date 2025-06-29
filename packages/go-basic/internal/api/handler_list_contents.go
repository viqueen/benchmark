package api

import (
	contentV1 "api/go-sdk/content/v1"
	"connectrpc.com/connect"
	"context"
)

func (s service) ListContents(ctx context.Context, c *connect.Request[contentV1.ListContentsRequest]) (*connect.Response[contentV1.ListContentsResponse], error) {
	//TODO implement me
	panic("implement me")
}

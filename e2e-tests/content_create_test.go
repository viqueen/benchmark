package e2e_tests

import (
	"connectrpc.com/connect"
	"context"
	"e2e-tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCreateContent(t *testing.T) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	client := helpers.NewContentClient(baseURL)
	ctx := context.Background()

	t.Run("success: creates content with valid data", func(t *testing.T) {
		fixture := helpers.NewValidContentFixture()
		req := helpers.NewCreateContentRequest(fixture.Title, fixture.Body)
		resp, err := client.CreateContent(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Msg.Content)
		helpers.AssertContentFields(t, resp.Msg.Content, fixture.Title, fixture.Body)
	})

	t.Run("success: creates multiple content items with unique IDs", func(t *testing.T) {
		var createdContentIds []string
		for i := 0; i < 5; i++ {
			fixture := helpers.NewValidContentFixture()
			req := helpers.NewCreateContentRequest(fixture.Title, fixture.Body)
			resp, err := client.CreateContent(ctx, req)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Msg.Content)
			helpers.AssertContentFields(t, resp.Msg.Content, fixture.Title, fixture.Body)

			// Check for unique ID
			assert.NotEmpty(t, resp.Msg.Content.Id, "Content ID should not be empty")
			assert.NotContains(t, createdContentIds, resp.Msg.Content.Id, "Content ID should be unique")

			createdContentIds = append(createdContentIds, resp.Msg.Content.Id)
		}
	})

	t.Run("invalid argument: creates content with empty title", func(t *testing.T) {
		fixture := helpers.NewValidContentFixture()
		req := helpers.NewCreateContentRequest("", fixture.Body)
		resp, err := client.CreateContent(ctx, req)

		require.Nil(t, resp)
		require.Error(t, err, "Expected error when creating content with empty title")
		var connectErr *connect.Error
		require.ErrorAs(t, err, &connectErr, "Error should be a connect error")
		assert.Equal(t, connect.CodeInvalidArgument, connectErr.Code(), "Expected InvalidArgument error code")
	})

	t.Run("invalid argument: creates content with empty body", func(t *testing.T) {
		fixture := helpers.NewValidContentFixture()
		req := helpers.NewCreateContentRequest(fixture.Title, "")
		resp, err := client.CreateContent(ctx, req)

		require.Nil(t, resp)
		require.Error(t, err, "Expected error when creating content with empty body")
		var connectErr *connect.Error
		require.ErrorAs(t, err, &connectErr, "Error should be a connect error")
		assert.Equal(t, connect.CodeInvalidArgument, connectErr.Code(), "Expected InvalidArgument error code")
	})

	t.Run("invalid argument: creates content with empty title and body", func(t *testing.T) {
		req := helpers.NewCreateContentRequest("", "")
		resp, err := client.CreateContent(ctx, req)

		require.Nil(t, resp)
		require.Error(t, err, "Expected error when creating content with empty title and body")
		var connectErr *connect.Error
		require.ErrorAs(t, err, &connectErr, "Error should be a connect error")
		assert.Equal(t, connect.CodeInvalidArgument, connectErr.Code(), "Expected InvalidArgument error code")
	})

	t.Run("invalid argument: creates content with long title", func(t *testing.T) {
		longTitle := "a" + string(make([]byte, 141))
		fixture := helpers.NewValidContentFixture()
		req := helpers.NewCreateContentRequest(longTitle, fixture.Body)
		resp, err := client.CreateContent(ctx, req)

		require.Nil(t, resp)
		require.Error(t, err, "Expected error when creating content with long title")
		var connectErr *connect.Error
		require.ErrorAs(t, err, &connectErr, "Error should be a connect error")
		assert.Equal(t, connect.CodeInvalidArgument, connectErr.Code(), "Expected InvalidArgument error code")
	})
}

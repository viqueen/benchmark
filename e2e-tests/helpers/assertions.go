package helpers

import (
	contentV1 "api/go-sdk/content/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func AssertContentFields(t *testing.T, content *contentV1.Content, expectedTitle, expectedBody string) {
	t.Helper()
	require.NotNil(t, content, "content should not be nil")
	assert.NotEmpty(t, content.Id, "content ID should not be empty")
	assert.Equal(t, expectedTitle, content.Title, "content title should match")
	assert.Equal(t, expectedBody, content.Body, "content body should match")
}

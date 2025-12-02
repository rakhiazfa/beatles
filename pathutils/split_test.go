package pathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	path := "/api/users/:userId/roles/:id"
	segments := Split(path)

	assert.Equal(t, []string{"api", "users", ":userId", "roles", ":id"}, segments)
}

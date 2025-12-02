package pathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	a := "/api/users/"
	b := "/:userId/roles/:id/"

	joined := Join(a, b)

	assert.Equal(t, "/api/users/:userId/roles/:id", joined)
}

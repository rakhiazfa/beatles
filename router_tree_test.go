package beatles

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	router := NewRouterTree()
	router.Use(func(c Context) error {
		return nil
	})

	apiGroup := router.Group("/api")
	apiGroup.Use(func(c Context) error {
		return nil
	})

	userGroup := apiGroup.Group("/users")
	userGroup.Use(func(c Context) error {
		return nil
	})

	userRoleGroup := userGroup.Group("/:userId/roles")
	userRoleGroup.Use(func(c Context) error {
		return nil
	})

	userRoleGroup.Get("/:id", func(c Context) error {
		return nil
	})

	route, err := router.Search(http.MethodGet, "/api/users/018af960-eaae-4768-a7e8-d6d36f55c530/roles/22011fda-8e34-4a79-b7e2-b597fd18bd02")

	require.NoError(t, err)
	require.NotNil(t, route)
	require.NotEmpty(t, route.handlers)
	require.NotEmpty(t, route.parameters)
}

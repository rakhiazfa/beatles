package beatles

import (
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockUsers = []JSON{
	{
		"id":    "b6dd6e16-ea28-4eb7-b67e-21d6cdbccb02",
		"name":  "John Doe",
		"email": "john.doe@example.com",
	},
	{
		"id":    "ec360b14-cec8-484d-b12e-613949ec7425",
		"name":  "Michel",
		"email": "michel@example.com",
	},
}

func TestApplication(t *testing.T) {
	t.Parallel()

	app := New(ApplicationConfig{
		Name: "Beatles",
	})

	router := app.Router()

	apiGroup := router.Group("/api")

	apiGroup.Get("/swagger/**", func(c Context) error {
		return nil
	})

	userGroup := apiGroup.Group("/users")

	userGroup.Get("/", func(c Context) error {
		return c.Response().JSON(JSON{
			"data": mockUsers,
		})
	})
	userGroup.Get("/:id", func(c Context) error {
		id := c.Request().PathVariable("id")

		index := slices.IndexFunc(mockUsers, func(user JSON) bool {
			return user["id"] == id
		})

		if index == -1 {
			return NewHTTPError(http.StatusNotFound, "User not found")
		}

		return c.Response().JSON(JSON{
			"data": mockUsers[index],
		})
	})

	assert.NoError(t, app.Start(":8080"))
}

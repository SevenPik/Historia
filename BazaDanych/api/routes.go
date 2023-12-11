package api

import (
	"net/http"

	"API/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo, h *handlers.Handler) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	people := e.Group("/people")
	people.GET("", h.GetPeople)
	people.POST("", h.CreatePerson)
	people.GET("/:id", h.GetPerson)
	people.PUT("/:id", h.UpdatePerson)
	people.DELETE("/:id", h.DeletePerson)

	people.GET("/:id/ranks", h.GetPersonRanks)
	people.POST("/:id/ranks", h.CreatePersonRank)
	people.DELETE("/:id/ranks/:id", h.DeletePersonRank)

}

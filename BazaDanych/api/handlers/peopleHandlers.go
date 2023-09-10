package handlers

import (
	"context"
	"net/http"
	"strconv"

	"API/db"

	"github.com/labstack/echo/v4"
)

type Person struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (h *DBHandler) GetPeople(c echo.Context) error {
	people, err := h.DB.GetPeople(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, people)
}

func (h *DBHandler) CreatePerson(c echo.Context) error {
	p := new(Person)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.DB.CreatePerson(context.Background(), p.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *DBHandler) GetPerson(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	p, err := h.DB.GetPerson(context.Background(), int32(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

func (h *DBHandler) UpdatePerson(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	p := new(Person)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := db.UpdatePersonParams{
		ID:   int32(id),
		Name: p.Name,
	}

	err = h.DB.UpdatePerson(context.Background(), data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *DBHandler) DeletePerson(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.DB.DeletePerson(context.Background(), int32(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

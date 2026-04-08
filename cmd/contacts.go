package main

import (
	"net/http"
	"strconv"

	"github.com/knadh/listmonk/internal/core"
	"github.com/knadh/listmonk/models"
	"github.com/labstack/echo/v4"
)

func (a *App) GetContacts(c echo.Context) error {
	var (
		orderBy = c.QueryParam("orderBy")
		order   = c.QueryParam("order")
		pg      = a.pg.NewFromURL(c.Request().URL.Query())
	)
	searchParams := core.SearchContactParams{}
	if err := c.Bind(&searchParams); err != nil {
		return err
	}
	contacts, total, err := a.core.QueryContacts(searchParams, orderBy, order, pg.Limit, pg.Offset)
	if err != nil {
		return err
	}
	// No results.
	if len(contacts) == 0 {
		return c.JSON(http.StatusOK, okResp{models.PageResults{Results: []models.Contact{}}})
	}
	out := models.PageResults{
		Results: contacts,
		Total:   total,
		Page:    pg.Page,
		PerPage: pg.PerPage,
	}
	return c.JSON(http.StatusOK, okResp{out})
}

func (a *App) GetContact(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	contact, err := a.core.GetContactByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{contact})
}

func (a *App) CreateContact(c echo.Context) error {
	var contact models.Contact
	if err := c.Bind(&contact); err != nil {
		return err
	}
	out, err := a.core.CreateContact(contact)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{out})
}

func (a *App) UpdateContact(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	var contact models.Contact
	if err := c.Bind(&contact); err != nil {
		return err
	}
	out, err := a.core.UpdateContact(id, contact)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{out})
}

func (a *App) DeleteContact(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	if err := a.core.DeleteContact(id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{nil})
}

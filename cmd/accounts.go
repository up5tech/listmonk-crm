package main

import (
	"net/http"
	"strconv"

	"github.com/knadh/listmonk/internal/core"
	"github.com/knadh/listmonk/models"
	"github.com/labstack/echo/v4"
)

func (a *App) GetAccounts(c echo.Context) error {
	var (
		orderBy = c.QueryParam("orderBy")
		order   = c.QueryParam("order")
		pg      = a.pg.NewFromURL(c.Request().URL.Query())
	)
	searchParams := core.SearchAccountParams{}
	if err := c.Bind(&searchParams); err != nil {
		return err
	}
	accounts, total, err := a.core.QueryAccounts(&searchParams, orderBy, order, pg.Limit, pg.Offset)
	if err != nil {
		return err
	}
	// No results.
	if len(accounts) == 0 {
		return c.JSON(http.StatusOK, okResp{models.PageResults{Results: []models.Account{}}})
	}

	out := models.PageResults{
		Results: accounts,
		Total:   total,
		Page:    pg.Page,
		PerPage: pg.PerPage,
	}

	return c.JSON(http.StatusOK, okResp{out})
}

func (a *App) GetAccount(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	account, err := a.core.GetAccountByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{account})
}

func (a *App) CreateAccount(c echo.Context) error {
	account := models.Account{}
	if err := c.Bind(&account); err != nil {
		return err
	}

	out, err := a.core.CreateAccount(account)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{out})
}

func (a *App) UpdateAccount(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	account := models.Account{}
	if err := c.Bind(&account); err != nil {
		return err
	}
	if err := a.core.UpdateAccount(id, account); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{account})
}

func (a *App) DeleteAccount(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	if err := a.core.DeleteAccount(id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, okResp{nil})
}

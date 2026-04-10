package main

import (
	"database/sql"
	"net/http"

	"github.com/knadh/listmonk/models"
	"github.com/labstack/echo/v4"
)

func (a *App) GetModules(c echo.Context) error {
	var modules = []models.CRMModule{
		{
			Name:  "Accounts",
			Label: "Accounts",
		},
		{
			Name:  "Contacts",
			Label: "Contacts",
		},
		{
			Name:  "Deals",
			Label: "Deals",
		},
		{
			Name:  "Tasks",
			Label: "Tasks",
		},
		{
			Name:  "Notes",
			Label: "Notes",
		},
		{
			Name:  "Products",
			Label: "Products",
		},
		{
			Name:  "Quotes",
			Label: "Quotes",
		},
		{
			Name:  "Orders",
			Label: "Orders",
		},
		{
			Name:  "Invoices",
			Label: "Invoices",
		},
	}
	return c.JSON(http.StatusOK, okResp{modules})
}

func (a *App) GetModuleFields(c echo.Context) error {
	// @TODO
	return nil
}

func (a *App) GetLayout(c echo.Context) error {
	module := c.Param("module")
	layoutType := c.Param("layoutType")
	var layout models.Layout
	if err := a.db.Get(&layout, `
		SELECT id, module, layout_type, meta FROM layouts WHERE module = $1 AND layout_type = $2
	`, module, layoutType); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, okResp{layout})
		}
		return err
	}
	return c.JSON(http.StatusOK, okResp{layout})
}

func (a *App) UpdateLayout(c echo.Context) error {
	layout := models.Layout{}
	if err := c.Bind(&layout); err != nil {
		return err
	}
	// check if layout name already exists
	var existingLayout models.Layout
	if err := a.db.Get(&existingLayout, `
		SELECT id, module, layout_type FROM layouts WHERE module = $1 AND layout_type = $2
	`, layout.Module, layout.LayoutType); err != nil {
		return err
	}
	// insert if not exists
	if existingLayout.ID == 0 {
		err := a.db.Get(&layout, `
			INSERT INTO layouts (module, layout_type, meta)
			VALUES ($1, $2, $3)
			RETURNING id, module, layout_type, meta
		`, layout.Module, layout.LayoutType, layout.Meta)
		if err != nil {
			return err
		}
	} else {
		// update if exists
		err := a.db.Get(&layout, `
			UPDATE layouts SET meta = $1 WHERE id = $2
			RETURNING id, module, layout_type, meta
		`, layout.Meta, existingLayout.ID)
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, okResp{layout})
}

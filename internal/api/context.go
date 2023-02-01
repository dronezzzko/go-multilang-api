package api

import (
	"github.com/dronezzzko/go-multilang-api/internal/localizer/printer"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/message"
)

type LanguageContext struct {
	echo.Context
}

func (c *LanguageContext) Printer() *message.Printer {
	langH := c.Request().Header.Get("Accept-Language")

	p, ok := printer.Make(langH)
	if !ok {
		c.Logger().Errorf("parse Accept-Language header: unsupported language %q", langH)
	}

	return p
}

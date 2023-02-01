package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Just to demonstrate how to deal with pluralization.
// See for details: https://pkg.go.dev/golang.org/x/text/feature/plural
var sungTimes int

func Song(c echo.Context) error {
	defer func() {
		sungTimes++
	}()

	lc := c.(*LanguageContext)
	p := lc.Printer()

	// Using text in source language as a key.
	stat := p.Sprintf("This song has already been sung %d times!", sungTimes)

	// Using ID as a key. More convenient way.
	// Easy to read, search and any text change doesn't require changing the ID.
	song := p.Sprintf("ten-green-bottles-song")

	return c.String(
		http.StatusOK,
		fmt.Sprintf("%s\n\n%s\n", song, stat),
	)
}

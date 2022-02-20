package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/create", func(c echo.Context) error {
		title := c.FormValue("title")
		detail := c.FormValue("detail")
		expireDate := c.FormValue("expire_date")
	return c.String(http.StatusOK, "title:" + title + ", detail:" + detail + ", expire_date:" + expireDate)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
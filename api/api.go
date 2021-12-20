package api

import (
	"fmt"
	"net/http"

	"github.com/MathieuMoalic/amumax/engine"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func dirName(c echo.Context) error {
	res := fmt.Sprintf(`{"message":"%s"}`, engine.OD())
	return c.String(http.StatusOK, res)
}

func Start() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.Static("/home/mat/go/src/test1/frontend/public"))
	e.Static("/", "index.html")
	e.GET("/dir", dirName)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

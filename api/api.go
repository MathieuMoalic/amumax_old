package api

import (
	"fmt"
	"net/http"

	"github.com/MathieuMoalic/amumax/engine"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getDir(c echo.Context) error {
	res := fmt.Sprintf(`{"res":"%s"}`, engine.OD())
	return c.String(http.StatusOK, res)
}

func getTables(c echo.Context) error {
	return c.JSON(http.StatusOK, engine.ZTables)
}
func getImage(c echo.Context) error {
	img := engine.GUI.GetRenderedImg()
	return c.Stream(http.StatusOK, "image/png", img)
}

func Start() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// e.Use(middleware.Static("/home/mat/go/src/test1/frontend/public"))
	// e.Static("/", "index.html")
	e.GET("/dir", getDir)
	e.GET("/tables", getTables)
	e.GET("/image", getImage)

	// Start server
	e.Logger.Fatal(e.Start(":5001"))
}

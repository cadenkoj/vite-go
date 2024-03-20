package web

import (
	"embed"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	controllers_v1 "github.com/cadenkoj/vera/backend/controllers/v1"
	"github.com/cadenkoj/vera/backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	//go:embed all:static
	dist embed.FS
	//go:embed static/index.html
	indexHTML embed.FS

	staticDirFS     = echo.MustSubFS(dist, "static")
	staticIndexHtml = echo.MustSubFS(indexHTML, "static")
)

func RegisterHandlers(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	v1.GET("/profile", controllers_v1.GetProfiles)
	v1.GET("/profile/:slug", controllers_v1.GetProfile)
	v1.POST("/profile", controllers_v1.PostProfile)
	v1.PATCH("/profile/:slug", controllers_v1.PatchProfile)
	v1.DELETE("/profile/:slug", controllers_v1.DeleteProfile)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		e.Logger.Error(err)
		if err := c.JSON(utils.NewError(code)); err != nil {
			e.Logger.Error(err)
		}
	}

	switch mode := os.Getenv("GO_ENV"); mode {
	case "development", "test":
		startDevServer(e)
		return
	}

	e.FileFS("/", "index.html", staticIndexHtml)
	e.StaticFS("/", staticDirFS)
}

func startDevServer(e *echo.Echo) {
	cmd := exec.Command("pnpm", "dev")
	cmd.Dir = "../frontend"
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	url, err := url.Parse("http://localhost:5173")
	if err != nil {
		log.Fatal(err)
	}

	balancer := middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url,
		},
	})

	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: balancer,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/api")
		},
	}))
}

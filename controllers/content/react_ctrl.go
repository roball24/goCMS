package content_ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/goCMS/context"
	"github.com/gocms-io/goCMS/routes"
	"net/http"
)

type ReactController struct {
	routes *routes.Routes
}

func DefaultReactController(routes *routes.Routes) *ReactController {
	rc := &ReactController{
		routes: routes,
	}

	rc.Default()
	return rc
}

func (rc *ReactController) Default() {
	rc.routes.Root.Static("/goCMS", "./content/goCMS")
	rc.routes.Root.GET("/admin", rc.serveReactAdmin)
	rc.routes.Root.GET("/login", rc.serveReactAdmin)
	rc.routes.Root.GET("/admin/*adminPath", rc.serveReactAdmin)
	rc.routes.NoRoute(rc.serveReact)
}

func (rc *ReactController) serveReact(c *gin.Context) {
	c.HTML(http.StatusOK, "react.tmpl", gin.H{
		"Theme":     context.Config.ActiveTheme,
		"AssetBase": context.Config.ActiveThemeAssetsBase,
		"Admin":     false,
	})
}

func (rc *ReactController) serveReactAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "react.tmpl", gin.H{
		"Theme":     context.Config.ActiveTheme,
		"AssetBase": context.Config.ActiveThemeAssetsBase,
		"Admin":     true,
	})
}

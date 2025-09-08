package httpx

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/tim_de/dartlog/v2/dartlog"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(isDebug bool, logger *dartlog.Logger) *Router {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(routesLogger(logger))

	return &Router{
		engine: router,
	}
}

func (router *Router) Group(groupPath string, middleware ...gin.HandlerFunc) *gin.RouterGroup {
	return router.engine.Group(groupPath, middleware...)
}

func (router *Router) CreateFileServer(route string, prefix string, fs string) {
	router.engine.GET(route, gin.WrapH(http.StripPrefix(prefix, http.FileServer(http.FS(os.DirFS(fs))))))
}

func routesLogger(logger *dartlog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		context := map[string]any{
			"host":      ctx.Request.Host,
			"url":       ctx.Request.URL.RequestURI(),
			"client_ip": ctx.ClientIP(),
			"method":    ctx.Request.Method,
			"status":    ctx.Writer.Status(),
			"latency":   time.Since(start),
		}

		if errVal, exists := ctx.Get("error"); exists {
			if err, ok := errVal.(error); ok && err != nil {
				context["error"] = err.Error()
			}
		}

		logger.Log(dartlog.INFO, "routesLogger", "request handled", context)
	}
}

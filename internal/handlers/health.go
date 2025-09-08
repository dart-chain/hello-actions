package handlers

import (
	"net/http"

	"github.com/dart-chain/hello-actions/internal/httpx"
	"github.com/gin-gonic/gin"
)

type HealthAPI struct{}

func NewHealthAPI() *HealthAPI {
	return &HealthAPI{}
}

func (h *HealthAPI) RegisterRoutes(router *gin.RouterGroup, name string) {
	api := router.Group(httpx.API_BASE)
	api.GET(httpx.API_HEALTH, h.HealthCheck(name))
}

// HealthCheck godoc
// @Summary      Health Check
// @Description  Returns the health status of the application
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func (h *HealthAPI) HealthCheck(name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": name + " is healthy"})
	}
}

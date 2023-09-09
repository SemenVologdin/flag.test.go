package lib

import (
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() RequestHandler {
	engine := gin.New()
	return RequestHandler{Gin: engine}
}

func (h *RequestHandler) InitRoutes() {
	h.Gin.GET("")
}

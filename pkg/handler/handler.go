package handler

import (
	"github.com/gin-gonic/gin"
	"notes/pkg/service"
)

type Handler struct {
	Services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitializeRouters() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		notes := api.Group("/notes")
		{
			notes.POST("/", h.CreateNote)
			notes.GET("/", h.GetNotes)
			notes.PUT("/", h.UpdateNote)
			notes.GET("/:id", h.GetNote)
			notes.DELETE("/:id", h.DeleteNote)
		}
	}

	return router
}

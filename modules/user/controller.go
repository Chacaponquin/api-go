package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.services.CreateUser(c.Request.Context(), &u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

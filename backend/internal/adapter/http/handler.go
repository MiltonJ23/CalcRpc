package http_adapter

import (
	"Calculator/backend/internal/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalculationRequest struct {
	Num1 int32 `json:"num1"`
	Num2 int32 `json:"num2"`
}

type HttpHandler struct {
	service service.CalculatorService
}

func NewHttpHandler(s service.CalculatorService) *HttpHandler {
	return &HttpHandler{service: s}
}

func (h *HttpHandler) Add(c *gin.Context) {
	var req CalculationRequest
	bindingError := c.ShouldBindJSON(&req)
	if bindingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingError.Error()})
		return
	}
	result := h.service.Add(req.Num1, req.Num2)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *HttpHandler) Sub(c *gin.Context) {
	var req CalculationRequest
	bindingError := c.ShouldBindJSON(&req)
	if bindingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingError.Error()})
		return
	}
	result := h.service.Sub(req.Num1, req.Num2)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
func (h *HttpHandler) Mul(c *gin.Context) {
	var req CalculationRequest
	bindingError := c.ShouldBindJSON(&req)
	if bindingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingError.Error()})
		return
	}
	result := h.service.Mul(req.Num1, req.Num2)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
func (h *HttpHandler) Div(c *gin.Context) {
	var req CalculationRequest
	bindingError := c.ShouldBindJSON(&req)
	if bindingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingError.Error()})
		return
	}
	result, DivError := h.service.Div(req.Num1, req.Num2)
	if DivError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": DivError.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *HttpHandler) Mod(c *gin.Context) {
	var req CalculationRequest
	bindingError := c.ShouldBindJSON(&req)
	if bindingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingError.Error()})
		return
	}
	result, ModError := h.service.Mod(req.Num1, req.Num2)
	if ModError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ModError.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func RegisterRoutes(r *gin.Engine, h *HttpHandler) {
	r.POST("/add", h.Add)
	r.POST("/sub", h.Sub)
	r.POST("/mul", h.Mul)
	r.POST("/div", h.Div)
	r.POST("mod", h.Mod)
}

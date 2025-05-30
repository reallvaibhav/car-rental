package handlers

import (
	"net/http"

	pb "proto/inventory"

	"github.com/gin-gonic/gin"
)

// CreateCar handles POST /cars
func (h *Handler) CreateCar(c *gin.Context) {
	var req pb.AddCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.clients.InventoryClient.AddCar(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListCars handles GET /cars
func (h *Handler) ListCars(c *gin.Context) {
	// Assuming SearchAvailableCars is the method to list all cars
	res, err := h.clients.InventoryClient.SearchAvailableCars(c.Request.Context(), &pb.SearchRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetCar handles GET /cars/:id
func (h *Handler) GetCar(c *gin.Context) {
	carID := c.Param("id")
	res, err := h.clients.InventoryClient.GetCarByID(c.Request.Context(), &pb.CarIDRequest{CarId: carID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateCar handles PUT /cars/:id
func (h *Handler) UpdateCar(c *gin.Context) {
	carID := c.Param("id")
	var req pb.UpdateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CarId = carID // Assign carID from URL parameter to the struct field

	res, err := h.clients.InventoryClient.UpdateCar(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteCar handles DELETE /cars/:id
func (h *Handler) DeleteCar(c *gin.Context) {
	carID := c.Param("id")
	res, err := h.clients.InventoryClient.DeleteCar(c.Request.Context(), &pb.CarIDRequest{CarId: carID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

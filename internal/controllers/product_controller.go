package controllers

import (
	"net/http"
	"strconv"

	"bumimedika-backend/internal/domain/dto"
	"bumimedika-backend/internal/services/product"
	"bumimedika-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	Manage(c *gin.Context)
	GetByID(c *gin.Context)
	GetList(c *gin.Context)
}

type ProductController struct {
	service product.IProductService
}

func InitProductController(s product.IProductService) IProductController {
	return &ProductController{service: s}
}

// POST /api/products
func (h *ProductController) Manage(c *gin.Context) {
	var product dto.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Manage(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		response := response.ResponseError{}
		response.Error(c, err.Error())
	}
}

// GET /api/products/:id
func (h *ProductController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: product}
		response.Success(c)
	}

	c.JSON(http.StatusOK, product)
}

// GET /api/products?name=para&page=1&limit=10&sort_field=stock&sort_order=desc
func (h *ProductController) GetList(c *gin.Context) {
	name := c.Query("name")
	sortField := c.DefaultQuery("sort_field", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, total, err := h.service.GetList(name, sortField, sortOrder, page, limit)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: products, TotalData: total, TotalPages: (total + int64(limit) - 1) / int64(limit)}
		response.Success(c)
	}

}

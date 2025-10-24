package controllers

import (
	"net/http"
	"strconv"

	"bumimedika-backend/internal/domain/dto"
	"bumimedika-backend/internal/services/customer"
	"bumimedika-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ICustomerController interface {
	Manage(c *gin.Context)
	GetByID(c *gin.Context)
	GetList(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type CustomerController struct {
	service customer.ICustumerService
}

func InitCustomerController(s customer.ICustumerService) ICustomerController {
	return &CustomerController{service: s}
}

// POST /api/customer
func (h *CustomerController) Manage(c *gin.Context) {
	var customer dto.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Manage(&customer)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: customer}
		response.Success(c)
	}
}

// GET /api/customer/:id
func (h *CustomerController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customer, err := h.service.GetByID(id)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: customer}
		response.Success(c)
	}
}

// GET /api/customer?name=para&page=1&limit=10&sort_field=stock&sort_order=desc
func (h *CustomerController) GetList(c *gin.Context) {
	name := c.Query("name")
	sortField := c.DefaultQuery("sort_field", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	customers, total, err := h.service.GetList(name, sortField, sortOrder, page, limit)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{Data: customers, TotalData: total, TotalPages: (total + int64(limit) - 1) / int64(limit)}
		response.Success(c)
	}

}

// GET /api/customer/:id
func (h *CustomerController) DeleteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		response := response.ResponseError{}
		response.Error(c, err.Error())
	} else {
		response := response.Response{}
		response.Success(c)
	}

}

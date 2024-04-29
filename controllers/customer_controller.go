package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"itmxgo/models"
	"itmxgo/services"
	util "itmxgo/utils"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	switch err.(type) {
	case *util.NotFoundError:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomer, err := services.CreateCustomer(customer)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, newCustomer)
}

func GetCustomerById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	customer, err := services.GetCustomerById(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	var customer models.CustomerWithId
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedCustomer, err := services.UpdateCustomer(customer)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedCustomer)
}

func DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = services.DeleteCustomer(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

package main

import (
	"itmxgo/controllers"
	"itmxgo/db"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(err)
	}
	r := gin.Default()

	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers/:id", controllers.GetCustomerById)
	r.PUT("/customers", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)

	r.Run(":8080")
}

package CustomersController

import "C"
import (
	"business/Entities/CustomerEntity"
	"business/Services/CustomerServices"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCustomer(c * gin.Context){
	//checking for validation
	customer := CustomerEntity.Customer{}
	err := CustomerEntity.DecodeJSON(c,&customer)
	if err!=nil{
		//fmt.Println("not bound with JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message":"Incorrect syntax",
		})
		return
	}

	if err = CustomerServices.AddCustomer(&customer) ; err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"message":"Please try again later",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":             customer,
		"response_code":    http.StatusCreated,
		"response_message": "Profile Created successfully",
	})
}



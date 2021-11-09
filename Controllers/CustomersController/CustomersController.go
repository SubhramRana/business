package CustomersController

import "C"
import (
	"business/Entities/CustomerEntity"
	"business/Services/CustomerServices"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCustomer(c * gin.Context){

	//deserializing the input JSON payload
	customer := CustomerEntity.Customer{}
	err := CustomerEntity.DecodeJSON(c,&customer)
	if err!=nil{
		//fmt.Println("not bound with JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message":"Incorrect syntax",
		})
		return
	}

	//checking for validation
	if err = customer.Validate() ; err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message" : "Invalid Input",
		})
		return
	}

	//Try to perform transaction
	if err = CustomerServices.AddCustomer(&customer) ; err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"message":"Please try again later",
		})
		return
	}

	//on [success]
	c.JSON(http.StatusCreated, gin.H{
		"data":             customer,
		"response_code":    http.StatusCreated,
		"response_message": "Profile Created successfully",
	})
}



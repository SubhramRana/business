package OrdersController

import (
	"business/Entities/OrderEntity"
	"business/Services/OrderServices"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func PlaceOrder(c * gin.Context){
	order:=OrderEntity.Order{}
	err:=c.BindJSON(&order)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err = OrderServices.PLaceOrder(&order) ; err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"data":order,
		"status_code" : http.StatusCreated,
		"message" : "order created successfully",
	})
}

func GetOrders(c * gin.Context){
	orders:=[]OrderEntity.Order{}

	//usr id is given in query string
	if CustomerID:=c.Query("CustomerID") ; CustomerID!=""{
		customerID,err := strconv.Atoi(CustomerID)
		if err!=nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid CustomerID",
			})
			return
		}

		//test for valid userId
		if customerID<=0{
			c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
				"message":"Negative CustomerID",
			})
			return
		}

		CustomerID:=uint32(customerID)
		err = OrderServices.GetOrderByUserId(CustomerID,&orders)
		if err!=nil{
			c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
				"message":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"user_id":CustomerID,
			"data":orders,
			"status_code":http.StatusOK,
		})
	}else{//user id is not given
		log.Info("All orders are requested")
		//return all orders
		err:= OrderServices.GetAllOrders(&orders)
		if err!=nil{
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"data":orders,
			"status_code":http.StatusOK,
		})
	}
}

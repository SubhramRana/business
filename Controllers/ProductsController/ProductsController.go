package ProductsController

import (
	"business/Entities/ProductEntity"
	"business/Services/ProductServices"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProducts(c * gin.Context){
	products:=[]ProductEntity.Product{}
	if err := ProductServices.GetProducts(&products) ; err!=nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,products)
}

func AddProduct(c * gin.Context){

	//checking for validation error
	product := ProductEntity.Product{}
	err:=ProductEntity.DecodeJSON(c,&product)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid syntax",
		})
		return
	}

	if err= ProductServices.AddProduct(&product) ; err!=nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":product,
		"status_code":http.StatusOK,
		"message": "Product added successfully",
	})
}
func UpdateProduct(c * gin.Context){
	//checking for invalid productId
	prodId ,err:= strconv.Atoi(c.Param("id"))
	if err!=nil || prodId<=0{
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//checking for wrong syntax in the payload
	product := ProductEntity.Product{}
	err=ProductEntity.DecodeJSON(c,&product)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message":"Invalid input data",
		})
		return
	}

	product.Name=""//can't change the name of a product

	ProdId := uint32(prodId)
	err = ProductServices.UpdateProduct(ProdId,&product)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"message":"Please try again later",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data":product,
		"status_code":http.StatusOK,
		"message":"Product updated successfully",
	})
}


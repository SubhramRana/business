package CustomersController

import "C"
import (
	"business/Entities/CustomerEntity"
	"business/Services/CustomerServices"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

func Login(c * gin.Context){
	// assume that all the users are authenticated without any credentials
	expTime:=time.Now().Add(5*time.Minute)
	token,err:=generateJWTToken(expTime)
	if err!=nil{
		c.AbortWithStatusJSON(401,gin.H{"message":err.Error()})
		return
	}
	c.Header("Autherization",token)
	c.JSON(200,gin.H{"message":"logged in successfully"})
}
type JWTClaims struct{
	jwt.StandardClaims
}
func (c * JWTClaims)Valid()error{
	return nil
}
func generateJWTToken(expTime time.Time)(signedToken string,err error){
	//define JWTclaim

	//initializre a claim
	Claims:=JWTClaims{}
	Claims.ExpiresAt=expTime.Unix()
	Claims.IssuedAt=time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,&Claims)

	key := "secrete_key"
	signedToken,err = token.SignedString([]byte(key))
	return
}



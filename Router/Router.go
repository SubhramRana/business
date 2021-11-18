package Router

import (
	"business/Controllers/CustomersController"
	"business/Controllers/OrdersController"
	"business/Controllers/ProductsController"
	"github.com/gin-gonic/gin"
)

func SetupRouter(server *gin.Engine){

		//login user(customer/retailer)
		server.POST("/login",CustomersController.Login)
		//Add user
		server.POST("/customers",CustomersController.AddCustomer)//done

		//place order
		server.POST("/orders", OrdersController.PlaceOrder)//Done
		//get orders
		server.GET("/orders", OrdersController.GetOrders)//Done

		//get all product details
		server.GET("/products", ProductsController.GetProducts)//Done
		//add a product
		server.POST("/products",ProductsController.AddProduct)//Done
		//update a product's price and quantity
		server.PATCH("/products/:id",ProductsController.UpdateProduct)//Done
			//{
			//	for this testcase
			//	{
			//		"Quantity":50
			//	}
			//	gave "invalid input data"
			//}

}

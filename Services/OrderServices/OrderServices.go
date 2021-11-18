package OrderServices

import (
	"business/Config"
	"business/Entities/CustomerEntity"
	"business/Entities/OrderEntity"
	"business/Entities/ProductEntity"
	"errors"
	"gorm.io/gorm"
	"time"
)

func PLaceOrder(order *OrderEntity.Order)(error) {
	ProductID := order.ProductID
	CustomerID := order.CustomerID
	Quantity := order.Quantity
	CooldownPeriod := 5 * time.Second

	return Config.DB.Transaction(func(tx * gorm.DB)error{
		time.Sleep(5*time.Second)
		//check for existance of customer
		customer := CustomerEntity.Customer{}
		result := tx.First(&customer, CustomerID)
		if result.Error != nil {
			return errors.New("Customer does not exists")
		}

		//check for cool-down period
		if time.Since(customer.LastOrderTime) <= CooldownPeriod {
			return errors.New("Too many order requests. Try after some time")
		}

		//check for product availability
		var product ProductEntity.Product
		result = tx.First(&product, ProductID)
		if result.Error != nil {
			return errors.New("Product does not exists")
		}

		//check against the available quantity
		curQuantity := product.Quantity
		if curQuantity < Quantity {
			return errors.New("Quantity not available")
		}

		//finally try to place the order

			//set the lastOrderTime for the customer
			if err := tx.Model(&CustomerEntity.Customer{}).
				Where("ID=?", CustomerID).
				Updates(CustomerEntity.Customer{LastOrderTime: time.Now()}).Error; err != nil {
				return err
			}
			//Insert a new order record
			if err := tx.Create(order).Error; err != nil {
				return err
			}
			//decrease the quantity of the product
			if err := tx.Model(&ProductEntity.Product{}).
				Where("ID=?", ProductID).
				Update("quantity", curQuantity-Quantity).Error; err != nil {
				return errors.New("try again4")
			}
			return nil

	})
}

func GetAllOrders(orders *[]OrderEntity.Order)(err error){
	Config.DB.Find(orders)
	return
}

func GetOrderByUserId(CustomerID uint32,orders *[]OrderEntity.Order)(err error){
	err=Config.DB.Model(&[]OrderEntity.Order{}).Where("customer_id=?",CustomerID).Find(&orders).Error
	return
}


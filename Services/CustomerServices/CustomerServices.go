package CustomerServices

import (
	"business/Config"
	"business/Entities/CustomerEntity"
)

func AddCustomer(customer *CustomerEntity.Customer)(err error){
	err=Config.DB.Create(customer).Error
	return
}

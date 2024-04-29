package services

import (
	"fmt"
	"itmxgo/models"
	util "itmxgo/utils"

	"itmxgo/db"
)

func CreateCustomer(customer models.Customer) (models.Customer, error) {
	if err := db.DB.Omit("ID").Create(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func UpdateCustomer(customer models.CustomerWithId) (models.Customer, error) {
	var existingCustomer models.Customer

	fmt.Println(customer)
	if err := db.DB.First(&existingCustomer, customer.ID).Error; err != nil {
		return existingCustomer, util.NewCustomerNotFoundError()
	}
	// Update fields
	existingCustomer.Name = customer.Name
	existingCustomer.Age = customer.Age

	if err := db.DB.Save(&existingCustomer).Error; err != nil {
		return existingCustomer, err
	}

	return existingCustomer, nil
}

func DeleteCustomer(id uint) error {
	var customer models.Customer
	if err := db.DB.First(&customer, id).Error; err != nil {
		return util.NewCustomerNotFoundError()
	}
	//soft delete record
	if err := db.DB.Delete(&customer).Error; err != nil {
		return err
	}

	return nil
}

func GetCustomerById(id uint) (models.Customer, error) {
	var customer models.Customer
	if err := db.DB.First(&customer, id).Error; err != nil {
		return customer, util.NewCustomerNotFoundError()
	}
	return customer, nil
}

package tests

import (
	"itmxgo/db"
	"itmxgo/models"
	"itmxgo/services"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	db.SetupTestDB()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateCustomer(t *testing.T) {
	mockCustomer := models.Customer{Name: "Sam Sepiol", Age: 30}

	newCustomer, err := services.CreateCustomer(mockCustomer)

	assert.NoError(t, err)
	assert.NotZero(t, newCustomer.ID)
	assert.Equal(t, mockCustomer.Name, newCustomer.Name)
	assert.Equal(t, mockCustomer.Age, newCustomer.Age)
}

func TestUpdateCustomer(t *testing.T) {
	mockExistingCustomer := models.Customer{Name: "Sam Sepiol", Age: 25}
	db.DB.Create(&mockExistingCustomer)

	mockUpdatedCustomer := models.CustomerWithId{ID: mockExistingCustomer.ID, Name: "Sam Smith", Age: 30}

	updatedCustomer, err := services.UpdateCustomer(mockUpdatedCustomer)

	assert.NoError(t, err)
	assert.Equal(t, mockUpdatedCustomer.Name, updatedCustomer.Name)
	assert.Equal(t, mockUpdatedCustomer.Age, updatedCustomer.Age)
}

func TestDeleteCustomer(t *testing.T) {

	mockCustomer := models.Customer{Name: "Sam Sepiol", Age: 30}
	db.DB.Create(&mockCustomer)

	err := services.DeleteCustomer(mockCustomer.ID)

	assert.NoError(t, err)
}

func TestGetCustomerById(t *testing.T) {

	mockCustomer := models.Customer{Name: "Sam Sepiol", Age: 30}
	db.DB.Create(&mockCustomer)

	customer, err := services.GetCustomerById(mockCustomer.ID)

	assert.NoError(t, err)
	assert.Equal(t, mockCustomer.Name, customer.Name)
	assert.Equal(t, mockCustomer.Age, customer.Age)
}

func TestGetCustomerByIdNotFound(t *testing.T) {

	_, err := services.GetCustomerById(999)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "customer not found")
}

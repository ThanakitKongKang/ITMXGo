# ITMXGo

This is a Golang project  for a job interview assignment at ITMX as Go Developer. It provides functionalities to create, read, update, and delete customer records through a RESTful API built with the Gin web framework.

## Project Structure:
- controllers: Contains handler functions for processing HTTP requests and generating responses.
- models: Defines the data structures representing customers.
- services: Implements business logic for managing customer data.
- utils: Houses utility functions and custom error types for error handling

## Features:
- Create Customer: Allows the addition of new customer records to the system.
- Get Customer by ID: Retrieves customer information based on the provided customer ID.
- Update Customer: Enables updating existing customer records with new information.
- Delete Customer: Removes a customer record from the system.


## Getting Started

*Ensure Go is installed on your system.*

To run this project locally, follow these steps:


1. Clone the Repository: Clone the itmxgo repository from https://github.com/ThanakitKongKang/itmxgo

2. Install Dependencies: Run go mod tidy to install project dependencies.

3. Run the project:

    ```
    go run main.go
    ```

4. Access the API endpoint:

    ```
    http://localhost:8080/customers
    ```

## Usage:
#### Create Customer:
- Endpoint: POST  http://localhost:8080/customers
- Request Body: JSON object representing the new customer.
- Response: JSON object containing the newly created customer details or error message.
#### Get Customer by ID:
- Endpoint: GET  http://localhost:8080/customers/:id
- Parameters: id - ID of the customer to retrieve.
- Response: JSON object containing customer details corresponding to the provided ID or error message.
#### Update Customer:
- Endpoint: PUT  http://localhost:8080/customers
- Request Body: JSON object representing the updated customer.
- Response: JSON object containing the updated customer details or error message.
#### Delete Customer:
- Endpoint: DELETE  http://localhost:8080/customers/:id
- Parameters: id - ID of the customer to delete.
- Response: JSON object indicating success or failure of the deletion operation.

## Running Tests
To run tests for this project, execute the following command:

```
go test ./tests
```
This command will run all tests in the project and display the results.

![image](https://github.com/ThanakitKongKang/ITMXGo/assets/44811853/9012e1de-1a74-41b1-b557-1bba6b58d5eb)


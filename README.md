# Ecommerce-Backend

## Overview

This repository contains the backend code for an e-commerce application, including user authentication, product management, cart functionality, and more. The APIs are documented below, covering various aspects of the system, such as cart management, address handling, and user authentication.

## Prerequisites

Before you start, ensure you have the following installed on your machine:

- [Golang](https://golang.org/doc/install) - Version 1.22.2
- [MongoDB](https://docs.mongodb.com/manual/installation/) - For database management
- [Postman](https://www.postman.com/downloads/) - For API testing

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/jash707/ecommerce-go.git
   ```

2. Navigate to the project directory:

   ```bash
   cd ecommerce-go
   ```

3. Install the required Go modules:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. The server will start on `http://localhost:8000`.

## API Documentation

### 1. Cart Management

- **Add to Cart**

  ```http
  GET /addtocart?productID=<productID>&userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

- **Remove from Cart**

  ```http
  GET /removeitem?productID=<productID>&userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

- **List Cart Items**

  ```http
  GET /listcart?userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

- **Checkout Cart**

  ```http
  GET /cartcheckout?userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

### Address Management

- **Add Address**

  ```http
  POST /addaddress?userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

  **Body:**

  ```json
  {
    "house_name": "white house",
    "street_name": "white street",
    "city_name": "washington",
    "pin_code": "332423432"
  }
  ```

  **Note:**
  The Address array is limited to two values: **Home** and **Work** addresses.

  - The **first** address added will be saved as the **Home** address.
  - The **second** address added will be saved as the **Work** address.
  - Adding more than two addresses is not acceptable.

- **Edit Home Address**

  ```http
  PUT /edithomeaddress?userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

  **Body:**

  ```json
  {
    "house_name": "aangan",
    "street_name": "citylight",
    "city_name": "surat",
    "pin_code": "395007"
  }
  ```

- **Edit Work Address**

  ```http
  PUT /editworkaddress?userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

  **Body:**

  ```json
  {
    "house_name": "white house",
    "street_name": "white street",
    "city_name": "washington",
    "pin_code": "332423432"
  }
  ```

- **Delete Addresses**

  ```http
  GET /deleteaddresses?userID=<userID>
  ```

  **Headers:**

  - `token`: JWT token for authentication

### 3. User Management

- **Sign Up**

  ```http
  POST /users/signup
  ```

  **Body:**

  ```json
  {
    "first_name": "Tony",
    "last_name": "Stark",
    "email": "starkinc@gmail.com",
    "password": "tony1234",
    "phone": "+918569745600"
  }
  ```

- **Login**

  ```http
  POST /users/login
  ```

  **Body:**

  ```json
  {
    "email": "starkinc@gmail.com",
    "password": "tony1234"
  }
  ```

### 4. Product Management

- **Add Product**

  ```http
  POST /admin/addproduct
  ```

  **Body:**

  ```json
  {
    "product_name": "Hp pavillion",
    "price": 4500,
    "rating": 7,
    "image": "hp.jpg"
  }
  ```

- **View Products**

  ```http
  GET /users/productview
  ```

- **Search Product by Query**

  ```http
  GET /users/search?name=<query>
  ```

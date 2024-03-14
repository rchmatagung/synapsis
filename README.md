
# Backend End Challenge

As a BackEnd Engineer you are required to create an online store application, you don't need to create a FrontEnd but focus on the BackEnd (RESTful API) only. The programming language you must use is Go-lang or Java spring boot.
You can develop your app by starting with prioritized features first. The following are the priority features to meet the MVP (minimum viable product) criteria:
- Customer can view product list by product category 
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions
- Login and register customers



## API Reference

#### Register Customer

```http
  POST /api/v1/customer/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required** |
| `password` | `string` | **Required** |
| `email` | `string` | **Required** |

#### Login Customer

```http
  GET /api/v1/customer/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required**. Use Body Form Data |
| `password`      | `string` | **Required**. Use Body Form Data |

#### Get All Category For Filter

```http
  GET /api/v1/category/getall
```

#### Get All Product With Category and Without Category

```http
  GET /api/v1/products/:categoryId
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category_id`      | `string` | **Required**. Use Query Param |

#### Note
if you give category_id = 0 it will display all products without category specifications, if category_id is filled in then it will display products with category specifications

#### Add Product To Shopping Cart

```http
  POST /api/v1/shopping_cart/addproduct
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `status_code`      | `string` | **Required** |
| `payment_total`      | `integer` | **Required** |
| `shopping_cart_product`      | `Array Object` | **Required** |

Parameter Array Object shopping_cart_product

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `product_id`      | `integer` | **Required** |
| `quantity`      | `integer` | **Required** |
| `sub_total`      | `integer` | **Required** |
| `price`      | `integer` | **Required** |

#### Get Shopping Cart

```http
  GET /api/v1/shopping_cart/get
```
#### Note
The following endpoint uses customer_id but has been directly deployed using middleware

#### Delete Product in Shopping Cart

```http
  DELETE /api/v1/shopping_cart/deleteproduct/:shopping_cart_product_id
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `shopping_cart_product_id`      | `integer` | **Required**. Use Query Param |

#### Checkout & Payment Transaction

```http
  PUT /api/v1/shopping_cart/checkout
```

#### Note
The following endpoint uses customer_id but has been directly deployed using middleware




## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd synapsis
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
    go run cmd/api/main.go
```


## Tech Stack

**Server:** Golang, Gorm, Fiber

**Database:** PostgreSQL


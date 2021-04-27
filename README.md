# Microservices
This project is a simple skeleton code for microservice architecture pattern using Golang and Docker.

## Install

1. Clone this git repoistory.

    ```bash
    https://github.com/senowijayanto/microservices
    ```

2. Go to the microservices directory.

    ```bash
    cd microservices
    ```

3. Start the docker composed cluster.

    ```bash
    docker-compose up -d
    ```

3. Running from local.

    ```bash
    http://localhost:8080/api/v1/users
    http://localhost:8080/api/v1/products
    http://localhost:8080/api/v1/orders
    ```

## Config
In this file nginx uses port 80, then it is exposed from the docker-compose file using port 8080. You can change it by editing the docker-compose.yml file, then select the service with the server name and edit the ports.
## Services
This project was decomposed into three cores microservices. All of them are independently deployable applications, organized around certain business domains.

### User Service
Provides serveral API for user account.
| Method | Path              | Description               |
|--------|-------------------|---------------------------|
| POST   | /api/v1/users     | Create new user           |
| GET    | /api/v1/users     | Get all user informations |
| GET    | /api/v1/users/{id}| Get user with ID          |
| PUT    | /api/v1/users/{id}| Edit user with ID         |
| DELETE | /api/v1/users/{id} | Delete user with ID      |

**_Sample POST User_**
```
Path : localhost:8080/api/v1/users
Body :
{
    "email": "sandra@gmail.com"
}
```

### Product Service
Provides several API for product account.
| Method | Path                 | Description                  |
|--------|----------------------|------------------------------|
| POST   | /api/v1/products     | Create new product           |
| GET    | /api/v1/products     | Get all product informations |
| GET    | /api/v1/products/{id}| Get product with ID          |
| PUT    | /api/v1/products/{id}| Edit product with ID         |
| DELETE | /api/v1/products/{id}| Delete product with ID       |

**_Sample POST Product_**
```
Path : localhost:8080/api/v1/products
Body :
{
    "name": "Laptop Lenovo Thinkpad",
    "price": 700000,
    "quantity": 10
}
```

### Order Service
Provides several API for order product.
| Method | Path                 | Description                  |
|--------|----------------------|------------------------------|
| POST   | /api/v1/orders       | Create new order             |
| GET    | /api/v1/orders       | Get all order informations   |

**_Sample POST Order_**
```
Path : localhost:8080/api/v1/orders
Body :
{
    "product_id": 3,
    "user_id": 1,
    "qty": 2
}
```
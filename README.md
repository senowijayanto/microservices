# Microservices
This project is a simple skeleton code for microservice architecture pattern using Golang and Docker.

## Services
This project was decomposed into three cores microservices. All of them are independently deployable applications, organized around certain business domains.

### User Service
Provides serveral API for user account.
| Method | Path              | Description               |
|--------|-------------------|---------------------------|
| POST   | /api/v1/users     | Create new user           |
| GET    | /api/v1/users     | Get all user informations |
| GET    | /api/v1/users/:id | Get user with ID          |
| PUT    | /api/v1/users/:id | Edit user with ID         |
| DELETE | /api/v1/users/:id | Delete user with ID       |


### Product Service
Provides several API for product account.
| Method | Path                 | Description                  |
|--------|----------------------|------------------------------|
| POST   | /api/v1/products     | Create new product           |
| GET    | /api/v1/products     | Get all product informations |
| GET    | /api/v1/products/:id | Get product with ID          |
| PUT    | /api/v1/products/:id | Edit product with ID         |
| DELETE | /api/v1/products/:id | Delete product with ID       |

### Order Service
Provides several API for order product.
| Method | Path                 | Description                  |
|--------|----------------------|------------------------------|
| POST   | /api/v1/orders       | Create new order             |
| GET    | /api/v1/orders       | Get all order informations   |

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
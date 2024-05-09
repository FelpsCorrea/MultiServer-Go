# MultiServer-Go
### This repository is a Server-Client API developed in Golang as a partial assessment for the completion of the Postgraduate Degree in Golang

### In this work, concepts such as Clean Architecture, Migrations, gRPC, GraphQL, Web Application, Dependency Injection, Environment Variables, events with RabbitMQ, etc., are utilized.


## Environment Configuration
Some necessary libraries and programs:

- Go (Remember to run the command `go mod tidy`)

- Evans

    ```
    go install github.com/ktr0731/evans@latest
    ```

- RabbitMQ and MySQL (included in the docker-compose file)

    ```
    docker-compose up -d
    ```

- Golang Migrate

    ```
    curl -fsSL https://packagecloud.io/golang-migrate/migrate/gpgkey | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/migrate.gpg >/dev/null
    ```
    ```
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/migrate.list
    ```
    ```
    sudo apt-get update
    ```
    ```
    sudo apt-get install -y migrate
    ```
    - Once configured, to run the initial migration, use the following command:
        ```
        migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3307)/orders" -verbose up
        ```
        - Adjust the information according to your settings, in the docker-compose file in question, MySQL is using port `3307`.

---
You can access the database through the terminal:

1. Access the MySQL environment from docker:
    ```
    docker-compose exec mysql bash
    ```
2. Access your created schema:
    ```
    mysql -uroot -p orders
    ```
---
Remember to configure the environment variables in `cmd/.env` as desired.

---

## Testing the Application

- Go to the folder where the `main.go` file is located:
    ```
    cd cmd/ordersystem
    ```
- Start the server:
    ```
    go run main.go wire_gen.go
    ```

- Testing gRPC
    - Interact with gRPC using the command (Remember to be at the project root):
    ```
    evans -r repl
    ```

- Testing graphQL
    - Interact using the graphical interface available at `http://localhost:8080` using the following mutations and queries:

    ```
    mutation createOrder {
        createOrder(input: {id: "ccc", Price: 10.0, Tax: 2.0}){
            id
            Price
            FinalPrice
        }
    }

    query orders{
        orders {
            id
            Price
            FinalPrice
        }
    }
    ```

- Testing Web
    - Interact with the test file in `internal/infra/web/test` using the `REST Client` extension of VSCODE



# Voilá! :)
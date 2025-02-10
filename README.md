# TransactionX

## Overview

The application was implemented in Golang, with PostgreSQL chosen as the persistence layer. The orchestration of the containers was performed using Docker Compose, and the service is configured to listen on any interface of the host on port 8080.

Designed to manage financial transactions, the application also supports conversion to any currency defined by the U.S. Treasury, provided there is data available for the currency in the last 6 months.

Transactions must include a description of up to 50 characters, a positive floating-point value, and a date in the format "YYYY-MM-DD HH:MM:SS." All validations are performed at the time of each transaction's registration.

### Supported Functionalities

-   **Register a transaction**: persist on database a financial transaction within the specifications.
-   **List all transactions**: fetch a list of all registered transactions.
-   **Convert a transaction**: convert the purchase amount of a transaction for a currency specified for some country.

## Usage

### 1. Clone the Repository

    git clone https://github.com/paulocesarvasco/transactionx.git
    cd transactionx

### 2. Start the Application

    docker-compose up

### 3. Access the service

    http://127.0.0.1:8080/index.html

An HTML file has been added to enhance the usability of the services. By accessing the service through a browser, the API will route and serve this file. This interface allows users register transactions in an user friendly way.

### 4. Postman Collection

Inside `collections` folder is posible to access a small collection of requests to test the application endpoints.

## Documentation

The API documentation was implemented using the OpenAPI pattern and can be found in the `doc` folder. And it is possible to render the file in:

    https://editor.swagger.io/

## Development workflow

For project development management, a GitHub board was created where tasks were tracked from the documentation phase through TDD and feature development.

A **branch model** was established, where the `main` branch serves as the productive branch, and all new features must follow the naming convention `feature/*` before being merged.

A **CI pipeline** was implemented to enforce the branch model, build the project, and run tests. The pipeline also evaluates test coverage, which must meet a minimum threshold of **70%**, and ensures that the source code follows Golang's standard formatting.

The project can be accessed here:

    https://github.com/users/paulocesarvasco/projects/1/views/1

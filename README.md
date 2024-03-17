# Simple REST ful Fiber
This project is a simple implementation of a RESTful API using Fiber, a fast and flexible 

## Features
- CRUD operations for products
- Input validation
- Error handling

## Installation
Ensure you have Go and MySQL installed.

1. Clone this repository:

```bash
git clone https://github.com/lendral3n/Simple-REST-ful-Fiber.git
```

2. Navigate into the project directory:
```bash
cd Simple-REST-ful-Fiber
```

3. Copy the .env.example file to .env and update the environment variables to match your MySQL configuration.

4. Run the server:
```bash
go run main.go
```
The server should now be running at http://localhost:8000.

## Usage
Here are some of the available endpoints:

- POST /product: Create a new product
- GET /product: Retrieve all products
- DELETE /product/:id: Delete a product by ID
- PUT /product/:id: Update a product by ID

And here is your `.env.example` file:

```plaintext
DBUSER=
DBPASS=
DBHOST=
DBPORT=
DBNAME=
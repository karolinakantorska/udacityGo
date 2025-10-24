# Descripton

A project for a go canguage curse. Implementing CRUD operation for user "database"

# Instalation

In project folder run go ron main.go to serve the project on http://localhost:3000/
Or run *.exe file in project folder.

# Usage

use curl or Postman to:

1. GET http://localhost:3000/customers

display a list of customeres

curl http://localhost:3000/customers


2. GET http://localhost:3000/customers/{id}

get one customer

curl http://localhost:3000/customers/{id}


3. POST http://localhost:3000/customers

create a customer

curl -X POST http://localhost:3000/customers/{id}
    -H "Content-Type: application/json" \
    -d '{
		Name: "New Consumer",
		Role: "user",
		Email:"new.consumer@gmail.com",
		Phone: "0041 76 123 45 99",
		Contacted: true,}'

4. PATCH http://localhost:3000/customers/{id}

update a customer

curl -X PATCH http://localhost:3000/customers/{id}
    -H "Content-Type: application/json" \
    -d '{Id:4,
		Name: "Michael Smith",
		Role: "user",
		Email:"michael.smith@gmail.com",
		Phone: "0041 76 123 45 99",
		Contacted: true,}'

5. DELETE http://localhost:3000/customers/{id}
delete a customer

curl -X DELETE http://localhost:3000/customers/{id}


# development instalation

apt install golang-go

go mod init udacityGo

go get github.com/gorilla/mux

go mod tidy
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

in windows:
curl -Method POST "http://localhost:3000/customers" `
-Headers @{"Content-Type"="application/json"} `
-Body '{
"name":"New Consumer",
"role":"user",
"email":"new.consumer@gmail.com",
"phone":41761234599,
"contacted":true}'


4. PATCH http://localhost:3000/customers/{id}

update a customer

in windows:
curl -Method PATCH "http://localhost:3000/customers/4" `
-Headers @{"Content-Type"="application/json"} `
-Body '{
"id":"4",
"name":"Michael Doe",
"role":"user",
"email":"michael.doe@gmail.com",
"phone":41761234599,
"contacted":true
}'


5. DELETE http://localhost:3000/customers/{id}
delete a customer

in windows:
curl -Method DELETE http://localhost:3000/customers/2


# development instalation

apt install golang-go

go mod init udacityGo

go get github.com/gorilla/mux

go mod tidy


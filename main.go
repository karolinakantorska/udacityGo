package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id        string    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var	c1 = Customer{
		Id:"1",
		Name: "John Doe",
		Role: "user",
		Email:"john.doe@gmail.com",
		Phone: 0041761234567,
		Contacted: false,
	}

var	c2 = Customer{
		Id:"2",
		Name: "John Smyth",
		Role: "admin",
		Email:"john.smyth@gmail.com",
		Phone: 0041761234567,
		Contacted: true,
	}

var	c3 = Customer{
		Id:"3",
		Name: "Joan Roberts",
		Role: "user",
		Email:"joan.roberts@gmail.com",
		Phone: 0041761234567,
		Contacted: true,
	}

var	data_base =[]Customer{c1,c2,c3}

func getCustomers(respWriter http.ResponseWriter, request *http.Request){
	respWriter.Header().Set("Content-Type", "application/json")
	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(data_base)
}

func getCustomer(respWriter http.ResponseWriter, request *http.Request){
	respWriter.Header().Set("Content-Type", "application/json")

	id := mux.Vars(request)["id"]

	customers_ids := []string{}

	for _,customer :=range data_base {
		customers_ids= append(customers_ids, customer.Id)
	}

	if slices.Contains(customers_ids, id ) {
		respWriter.WriteHeader(http.StatusOK)
		response := Customer {}
			for _,customer :=range data_base {
				if id == customer.Id{
					response= customer
				}
			}
		json.NewEncoder(respWriter).Encode(response)
	} else {
		respWriter.WriteHeader(http.StatusNotFound)
		respWriter.Write([]byte("Customer not found"))
	}
}

func showCustomers(respWriter http.ResponseWriter, request *http.Request){
	respWriter.Header().Set("Content-Type", "text/html")
	respWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(respWriter, `
			<fragment>
				<h4>Avaiable Endpoints</h4>
					<ul>
						<li>GET http://localhost:3000/customers</li>
						<li>GET http://localhost:3000/customers/{id}</li>
						<li>POST http://localhost:3000/customers</li>
						<li>PATCH http://localhost:3000/customers/{id}</li>
						<li>DELETE http://localhost:3000/customers/{id}</li>
					</ul>
			</fragment>
			<h4>Customers</h4>
		`,
	)
	for _,value := range data_base {
		fmt.Fprintf(respWriter, `
			<fragment>
				<h6>%v</h6>
				<p>%v</p>
				<p>%v</p>
			</fragment>
		`,
		value.Name,
		value.Email,
		value.Phone,
	)
	}
}

func addCustomer(respWriter http.ResponseWriter, request *http.Request){
	respWriter.Header().Set("Content-Type", "application/json")
	
	reqBody, error := io.ReadAll(request.Body)

	if error != nil || len(reqBody) == 0 {
		respWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	newCustomer := Customer{}
	// parse json body
	if err := json.Unmarshal(reqBody, &newCustomer); err != nil {
		respWriter.WriteHeader(http.StatusBadRequest)
		return
	}

    // Auto-generate ID if missing
    if newCustomer.Id == "" {
        maxID := 0
        for _, c := range data_base {
            // Convert existing string ID to int
            var cid int
            fmt.Sscanf(c.Id, "%d", &cid)
            if cid > maxID {
                maxID = cid
            }
        }
        newCustomer.Id = fmt.Sprintf("%d", maxID+1)
    }

    // Check for duplicate ID
    for _, c := range data_base {
        if c.Id == newCustomer.Id {
            respWriter.WriteHeader(http.StatusConflict)
            return
        }
    }

		data_base = append(data_base, newCustomer)
		respWriter.WriteHeader(http.StatusCreated)


	json.NewEncoder(respWriter).Encode(data_base)

}

func updateCustomer(respWriter http.ResponseWriter, request *http.Request){
	respWriter.Header().Set("Content-Type", "application/json")
	
	id := mux.Vars(request)["id"]

	reqBody, error := io.ReadAll(request.Body)

	if error != nil || len(reqBody) == 0 {
		respWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	newCustomer := Customer{}

	// parse json body
	if err := json.Unmarshal(reqBody, &newCustomer); err != nil {
respWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	customers_ids := []string{}

	for _,customer :=range data_base {
		customers_ids= append(customers_ids, customer.Id)
	}

	if slices.Contains(customers_ids, newCustomer.Id){
		for i,customer :=range data_base{
			if id == customer.Id {
				newCustomer.Id=customer.Id
				data_base[i] = newCustomer
			}
		}
		respWriter.WriteHeader(http.StatusOK)
	} else {
		respWriter.WriteHeader(http.StatusNotFound)
		respWriter.Write([]byte("Customer not found"))
		return
	}
	json.NewEncoder(respWriter).Encode(data_base)
}

func deleteCustomer(respWriter http.ResponseWriter, request *http.Request){
	respWriter.Header().Set("Content-Type", "application/json")

	id := mux.Vars(request)["id"]

	customers_ids := []string{}

	for _,customer :=range data_base {
		customers_ids= append(customers_ids, customer.Id)
	}

	if slices.Contains(customers_ids, id ){
		for i,customer :=range data_base {
			if id == customer.Id{
				data_base = slices.Delete(data_base,i, i+1)
				break
			}
		}
	} else {
		respWriter.WriteHeader(http.StatusNotFound)
		respWriter.Write([]byte("Customer not found"))
		return
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(data_base)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", showCustomers)
	router.HandleFunc("/customers", getCustomers)
	router.HandleFunc("/customers/{id}", getCustomer)
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("server is starting")
	http.ListenAndServe(":3000", router)
}
package controllers

import (
	"api_rest/internal/models"
	dbconfig "api_rest/internal/repositories"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ok
func GetCustomers(rw http.ResponseWriter, rq *http.Request) {
	conn, err := dbconfig.OpenConnection()

	if err != nil {
		return
	}

	rows, err := conn.Query("SELECT id, name, phone, adress FROM costumer")

	if err != nil {
		return
	}

	for rows.Next() {

		var id int
		var name, phone, adress string

		err := rows.Scan(&id, &name, &phone, &adress)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"Id": id, "name": name, "phone": phone, "adress": adress,
		})
	}

	conn.Close()

}

// ok
func GetCustumerById(rw http.ResponseWriter, rq *http.Request) {
	conn, err := dbconfig.OpenConnection()

	if err != nil {
		return
	}

	vars := mux.Vars(rq)
	id := vars["id"]

	rows, err := conn.Query("SELECT id, name, phone, adress FROM costumer WHERE id = $1", id)

	if err != nil {
		return
	}

	for rows.Next() {
		var name, phone, adress string

		err := rows.Scan(&id, &name, &phone, &adress)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"Id": id, "name": name, "phone": phone, "adress": adress,
		})
	}
	conn.Close()
}

// ok
func InsertCustumer(rw http.ResponseWriter, rq *http.Request) {

	conn, err := dbconfig.OpenConnection()

	if err != nil {
		return
	}

	if rq.Method != http.MethodPost {
		http.Error(rw, "erro no metodo de envio da requisição", http.StatusMethodNotAllowed)
		return
	}

	var customer models.Customer

	err = json.NewDecoder(rq.Body).Decode(&customer)

	if err != nil {
		http.Error(rw, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO costumer (name, phone, adress) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err = conn.QueryRow(query, customer.Name, customer.Phone, customer.Adress).Scan(&id)

	if err != nil {
		http.Error(rw, "Erro ao inserir", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"id": id,
	})

	conn.Close()
}

func UpdateCustomer(rw http.ResponseWriter, rq *http.Request) {
	conn, err := dbconfig.OpenConnection()
	if err != nil {
		return
	}
	vars := mux.Vars(rq)
	id_link := vars["id"]

	var customer models.Customer
	err = json.NewDecoder(rq.Body).Decode(&customer)

	query := `UPDATE costumer SET phone = $1, adress= $2 WHERE id = $3`

	err = conn.QueryRow(query, customer.Phone, customer.Adress, id_link).Scan(&customer.Id)

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"status": "Atualização concluida",
	})
	conn.Close()
}

func DeleteCustomer(rw http.ResponseWriter, rq *http.Request) {
	conn, err := dbconfig.OpenConnection()

	if err != nil {
		return
	}

	vars := mux.Vars(rq)
	id := vars["id"]

	conn.QueryRow(`DELETE FROM costumer WHERE id = $1 RETURNING id`, id).Scan(&id)

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"status": " excluido concluida ",
	})
	conn.Close()

}

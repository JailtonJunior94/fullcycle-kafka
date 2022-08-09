package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/models"
	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/repositories"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := repositories.InsertCustomer(customer)
	var res map[string]any

	if err != nil {
		res = map[string]any{
			"Success": false,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	}

	res = map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("Inserido com suceso! ID: %d", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

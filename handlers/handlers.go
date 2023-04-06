package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	types "github.com/pashamakhilkumarreddy/golang-web-server/types"
)

var founders = []types.Founder{
	{
		Id:      1,
		Name:    "Stacy",
		Age:     24,
		Email:   "stacy@ashido.love",
		Company: "Stacy Inc.",
	},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo!!!")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.Response{
		Message:    "pong",
		StatusCode: 200,
	}
	json.NewEncoder(w).Encode(response)
}

func AddFounderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var founder types.Founder

	json.NewDecoder(r.Body).Decode(&founder)

	founder.Name = strings.ToTitle(founder.Name)
	founders = append(founders, founder)

	json.NewEncoder(w).Encode(founders)
}

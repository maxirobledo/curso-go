// API REST simple sin framework utilizando solo standar library de Go (net/http)
// Un solo endpoint GET /health
// Que responde JSON

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Estructura de la respuesta JSON
type HealthResponse struct {
	Status string `json:"status"`
}

// Handler del endpoint GET /health
// w → respuesta HTTP
// r → request entrante
func healthHandler(w http.ResponseWriter, r *http.Request) {

	// Validamos metodo HTTP de la requests entrante
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status: "ok",
	}

	// Indicamos que devolvemos JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convertimos la struct a JSON y la enviamos
	json.NewEncoder(w).Encode(response)

}

func main() {
	// Registramos el endpoint
	http.HandleFunc("/health", healthHandler)

	log.Println("API escuchando en http://localhost:8080")

	// Levantamos el server http
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

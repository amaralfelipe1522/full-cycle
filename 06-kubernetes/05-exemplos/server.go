package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config-map", HelloConfigMap)
	http.HandleFunc("/secret", HelloSecret)
	http.HandleFunc("/healthz", HelloHealthz)
	http.ListenAndServe(":3000", nil)
}

// Exemplo de variáveis de ambiente usando ConfigMap
func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s. I'm %s!!", name, age)
}

// Exemplo de variáveis de ambiente usando ConfigMap como um Volume no Container
func HelloConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/my-family/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}

	fmt.Fprintf(w, "My Family: %s.", data)
}

// Exemplo de variáveis de ambiente usando Secrets
func HelloSecret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s - Password: %s", user, password)
}

// Exemplo de Healthz, onde nossa aplicação estará com problemas mas o Kubernetes não vai saber
func HelloHealthz(w http.ResponseWriter, r *http.Request) {
	// Verificar quanto tempo está em execução
	duration := time.Since(startedAt)

	if duration.Seconds() > 25 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}
}

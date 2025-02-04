package server_replicador

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", reenviarSolicitud)
    log.Println("Servidor replicador escuchando en :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}

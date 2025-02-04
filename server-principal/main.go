package server_principal

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            obtenerUsuarios(w, r)
        case http.MethodPost:
            crearUsuario(w, r)
        default:
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/usuarios/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            actualizarUsuario(w, r)
        case http.MethodDelete:
            eliminarUsuario(w, r)
        default:
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Servidor principal escuchando en :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

package server_principal

import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"
)

var (
    usuarios   = []Usuario{}
    idActual   = 1
    usuariosMu sync.Mutex
)

func crearUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
        http.Error(w, "Solicitud inválida", http.StatusBadRequest)
        return
    }
    usuariosMu.Lock()
    defer usuariosMu.Unlock()
    usuario.Id = idActual
    idActual++
    usuarios = append(usuarios, usuario)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(usuario)
}

func obtenerUsuarios(w http.ResponseWriter, r *http.Request) {
    usuariosMu.Lock()
    defer usuariosMu.Unlock()
    json.NewEncoder(w).Encode(usuarios)
}

func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }
    var usuarioActualizado Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuarioActualizado); err != nil {
        http.Error(w, "Solicitud inválida", http.StatusBadRequest)
        return
    }
    usuariosMu.Lock()
    defer usuariosMu.Unlock()
    for i, usuario := range usuarios {
        if usuario.Id == id {
            usuarios[i].User = usuarioActualizado.User
            usuarios[i].Nombre = usuarioActualizado.Nombre
            json.NewEncoder(w).Encode(usuarios[i])
            return
        }
    }
    http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }
    usuariosMu.Lock()
    defer usuariosMu.Unlock()
    for i, usuario := range usuarios {
        if usuario.Id == id {
            usuarios = append(usuarios[:i], usuarios[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

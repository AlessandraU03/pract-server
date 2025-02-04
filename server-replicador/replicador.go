package server_replicador

import (
    "io"
    "net/http"
)

func reenviarSolicitud(w http.ResponseWriter, r *http.Request) {
    cliente := &http.Client{}
    req, err := http.NewRequest(r.Method, "http://localhost:8080"+r.RequestURI, r.Body)
    if err != nil {
        http.Error(w, "Error al crear la solicitud", http.StatusInternalServerError)
        return
    }
    req.Header = r.Header
    resp, err := cliente.Do(req)
    if err != nil {
        http.Error(w, "Error al reenviar la solicitud", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    for k, v := range resp.Header {
        w.Header()[k] = v
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

package main

import (
    "sync"
    "pract/server_principal"
    "pract/server_replicador"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        server_principal.StartServer1()
    }()

    go func() {
        defer wg.Done()
        server_replicador.StartServer2()
    }()

    wg.Wait()
}
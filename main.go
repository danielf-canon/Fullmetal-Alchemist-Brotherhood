package main

import (
    "backend-alquimia/server"

)

func main() {
    s := server.NewServer()
    s.StartServer()
}
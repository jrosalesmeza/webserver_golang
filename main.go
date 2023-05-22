package main

// Los middlewares son protectores de rutas, que nos permiten verificar condiciones antes de aceptar una petición

import (
	"fmt"
	"web_server/handlers"
	s "web_server/server"
)

func main() {
	var port string = "3000"
	fmt.Printf("Sever corriendo en el puerto: %s \n", port)
	server := s.NewServer(":" + port)
	server.AddHandler("GET", "/", handlers.HandlerRoot)
	server.AddHandler("POST", "/create", handlers.PostRequest)
	server.AddHandler("POST", "/users", handlers.UserPostRequest)

	server.AddHandler("POST", "/api", server.AddMiddleware(handlers.HandlerHome, s.CheckAuth(), s.Logging()))
	server.Listen()
}

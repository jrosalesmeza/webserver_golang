package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"web_server/utils"
)

func CheckAuth() utils.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Check authentication....")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() utils.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start)) // Función anónima
			}()
			f(w, r)
		}
	}
}

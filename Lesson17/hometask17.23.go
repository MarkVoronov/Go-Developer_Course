package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// файл для записи логов
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Не удалось открыть файл для записи логов:", err)
	}
	defer logFile.Close()

	l := log.New(logFile, "", log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: logHandler(
			authHandler(mux),
		),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}
func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}
func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")

		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован",
				http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			l.Println("url:", r.URL) // логируем запрос
			h.ServeHTTP(w, r)

		})
	}
}

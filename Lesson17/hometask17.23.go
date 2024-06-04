package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Обертка для responseWriter, чтобы мы залогировали ответ сервера
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	// Значение по умолчанию - 200 OK
	return &responseWriter{w, http.StatusOK, new(bytes.Buffer)}
}

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
			
			// Создаём обёртку для http.ResponseWriter
			rw := NewResponseWriter(w)

			h.ServeHTTP(rw, r)

			// Логируем ответ
			l.Printf("Response status: %d", rw.statusCode)
			l.Printf("Response body: %s", rw.body.String())

		})
	}
}

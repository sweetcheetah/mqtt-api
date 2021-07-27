package server

import (
	"log"
	"net/http"
)

func (s *Server) routes() {
	s.router.HandleFunc("/bus-1-temp", s.handlerTemp)
	s.router.HandleFunc("/bus-1-humidity", s.handlerHumidity)
	s.router.HandleFunc("/", handlerRoot)

	//swagger
	fs := http.FileServer(http.Dir("./swagger-ui/"))
	s.router.PathPrefix("/swagger-ui").Handler(http.StripPrefix("/swagger-ui", fs))
}

func (s *Server) middleware() {
	s.router.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found."))
}

func (s *Server) handlerTemp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// TODO finish
	w.Write([]byte("72Fake"))
}

func (s *Server) handlerHumidity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// TODO finish
	w.Write([]byte("1%"))
}

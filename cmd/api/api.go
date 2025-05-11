package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiServer struct {
	port string
	pool *pgxpool.Pool
}

func NewApiServer(port string, pool *pgxpool.Pool) *ApiServer {
	return &ApiServer{
		port: port,
		pool: pool,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	// subrouter := router.PathPrefix("/v1/api").Subrouter()

	log.Println("Listening on port", s.port)

	return http.ListenAndServe(s.port, router)
}

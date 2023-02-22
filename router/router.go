package router

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type Router struct {
	db *sqlx.DB
}

func NewRouter(db *sqlx.DB) *Router {
	return &Router{
		db: db,
	}
}

func (r *Router) Ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("pong")); err != nil {
		panic(err)
	}
}

func (r *Router) DBNonce(w http.ResponseWriter, req *http.Request) {
	var (
		code      int
		respBytes = make([]byte, 0)
	)
	dbNonce, err := r.getDBNonce()
	if err != nil {
		code = http.StatusInternalServerError
		log.Printf(" failed to get database noce")
	} else {
		code = http.StatusOK
		resp := map[string]int{
			"nonce": dbNonce,
		}

		respBytes, err = json.Marshal(resp)
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(code)
	if _, err := w.Write(respBytes); err != nil {
		panic(err)
	}
}

func (r *Router) getDBNonce() (int, error) {
	var timestamp int
	err := r.db.Get(&timestamp, "SELECT CAST(EXTRACT(epoch FROM now()) AS INTEGER);")
	return timestamp, err
}

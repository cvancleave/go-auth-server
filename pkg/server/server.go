package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cvancleave/go-auth-server/pkg/utils"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	port      int
	secretKey string
	issuer    string
	audience  string
}

func Start() {

	// private encryption key - usually gotten from a secret store
	secretKey := "ellipsis"

	// set up server
	s := &server{
		port:      4001,
		secretKey: secretKey,
		issuer:    "go-auth-server",
		audience:  "go-auth-server-user",
	}

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", s.port),
		Handler:     s.routes(),
		IdleTimeout: time.Minute,
	}

	fmt.Println("starting go-auth-server on port:", s.port)

	// serve
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *server) routes() *httprouter.Router {

	router := httprouter.New()

	router.POST("/oauth/token", s.handleTokenRequest) // standard oauth endpoint
	router.POST("/validate", s.handleValidateRequest) // example of jwt validation

	// allow cors for non-get methods
	router.GlobalOPTIONS = http.HandlerFunc(options)

	return router
}

func options(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		utils.SetCorsHeaders(w, r)
	}
	w.WriteHeader(http.StatusOK)
}

package main

import (
	"fmt"
	"net/http"
)

func (a *application) ListenAndServe() error {
	host := fmt.Sprintf("%s:%s", a.server.host, a.server.port)

	srv := http.Server{
		Handler: a.routes(),
		Addr:    host,
	}

	a.infoLog.Printf("Server listening on :%s\n", host)

	return srv.ListenAndServe()
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	"ecksbee.com/telefacts/hydratables"
	"ecksbee.com/telefacts/renderer"
	"ecksbee.com/telefacts/sec"
	"ecksbee.com/telefacts/serializables"
	"ecksbee.com/telefacts/server"
	"github.com/gorilla/mux"
)

func main() {
	var ctx = context.Background()
	srv := setupServer()
	go func() {
		fmt.Println("Listening")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	listenForShutdown(ctx, wait, srv)
}

func setupServer() *http.Server {
	scache := server.NewSCache()
	hcache := server.NewHCache()
	seccache := server.NewSECCache()
	serializables.SetGlobalDir(path.Join(".", "taxonomies"))
	serializables.InjectCache(scache)
	hydratables.InjectCache(hcache)
	sec.InjectCache(seccache)
	r := mux.NewRouter()
	r.HandleFunc("/projects", server.Project()).Methods("GET", "POST")
	projectsRoute := r.PathPrefix("/projects").Subrouter()
	projectIDPrefix := projectsRoute.PathPrefix("/{id}")
	projectIDRoute := projectIDPrefix.Subrouter()
	projectIDRoute.HandleFunc("", server.Project()).Methods("GET")
	projectIDRoute.HandleFunc("/serializables", server.ProjectSerializable()).Methods("GET", "POST")
	projectIDRoute.HandleFunc("/renderables", server.ProjectRenderableIndex()).Methods("GET", "POST")
	projectIDRoute.HandleFunc("/renderables/{slug}", server.ProjectRenderable()).Methods("GET", "POST")
	rendererServer := renderer.LoadServer(path.Join(".", "renderer", "public"))
	r.HandleFunc("/renderer", rendererServer.ServeHTTP).Methods("GET")
	//todo serve open api spec at root

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
}

func listenForShutdown(ctx context.Context, grace time.Duration, srv *http.Server) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down")
	ctx, cancel := context.WithTimeout(ctx, grace)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}

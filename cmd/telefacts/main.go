package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"ecksbee.com/telefacts/internal/web"
	"ecksbee.com/telefacts/pkg/cache"
	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/serializables"
)

func main() {
	var ctx = context.Background()
	srv := setupServer()
	go func() {
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
	appCache := cache.NewCache(false)
	dir, err := os.Getwd()
	if err != nil {
		dir = filepath.Join(".")
	}
	wd := os.Getenv("WD")
	if wd == "" {
		wd = dir
	}
	serializables.WorkingDirectoryPath = filepath.Join(wd, "wd")
	gts := os.Getenv("GTS")
	if gts == "" {
		gts = dir
	}
	serializables.GlobalTaxonomySetPath = filepath.Join(gts, "gts")
	hydratables.InjectCache(appCache)
	hydratables.HydrateEntityNames()
	hydratables.HydrateFundamentalSchema()
	hydratables.HydrateUnitTypeRegistry()
	r := web.NewRouter()

	fmt.Println("telefacts<-0.0.0.0:8080")
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

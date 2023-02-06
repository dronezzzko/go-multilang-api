package main

import (
	"errors"
	"log"
	"os"

	"github.com/dronezzzko/go-multilang-api/internal/api"
	_ "github.com/dronezzzko/go-multilang-api/internal/translation"
	"github.com/spf13/pflag"
)

const (
	exitCodeNoErr       = 0
	exitCodeInvalidArgs = 2
)

const defaultPort = 8080

func main() {
	logger := log.New(os.Stderr, "api ", log.LstdFlags)

	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)
	fs.Int("port", defaultPort, "HTTP port to bind to")

	err := fs.Parse(os.Args[1:])

	switch {
	case errors.Is(err, pflag.ErrHelp):
		os.Exit(exitCodeNoErr)
	case err != nil:
		logger.Printf("parse arguments: %s\n", err.Error())
		os.Exit(exitCodeInvalidArgs)
	}

	err = run(fs, logger)
	if err != nil {
		logger.Fatalln(err)
	}
}

func run(fs *pflag.FlagSet, logger *log.Logger) error {
	apiPort, err := fs.GetInt("port")
	if err != nil {
		return err
	}

	server, err := api.NewServer(apiPort, logger)
	if err != nil {
		return err
	}

	return server.ListenAndServe()
}

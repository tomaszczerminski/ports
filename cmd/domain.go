package main

import (
	"context"
	"github.com/alexsasharegan/dotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomaszczerminski/test/pkg/domain"
	"github.com/tomaszczerminski/test/pkg/shared"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger := log.With().Str("service_name", "domain").Logger()
	if err := dotenv.Load(); err != nil {
		logger.Fatal().Err(err).Msg("cannot read env file")
	}
	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return log.With().Str("version", os.Getenv("API_VERSION"))
	})
	ctx := context.TODO()
	ctx = (&logger).WithContext(ctx)
	l, err := net.Listen("tcp", os.Getenv("DOMAIN_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start the TCP server")
	}
	repository := domain.NewRepository()
	server := grpc.NewServer()
	shared.RegisterStreamerServer(server, domain.NewServer(repository))
	if err := server.Serve(l); err != nil {
		log.Fatal().Err(err).Msg("failed to start gRPC server")
	}
}

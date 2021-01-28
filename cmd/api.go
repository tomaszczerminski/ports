package main

import (
	"context"
	"github.com/alexsasharegan/dotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomaszczerminski/test/pkg/api"
	"github.com/tomaszczerminski/test/pkg/api/startup"
	"github.com/tomaszczerminski/test/pkg/shared"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger := log.With().Str("service_name", "api").Logger()
	if err := dotenv.Load(); err != nil {
		logger.Fatal().Err(err).Msg("cannot read env file")
	}
	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return log.With().Str("version", os.Getenv("API_VERSION"))
	})
	ctx := context.TODO()
	ctx = (&logger).WithContext(ctx)
	conn, err := grpc.DialContext(
		ctx,
		os.Getenv("DOMAIN_URL"),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot connect to the domain server")
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Fatal().Err(err).Msg("gRPC connection has not been closed properly")
		}
	}()
	streamer := shared.NewStreamerClient(conn)
	blueprints := startup.FetchPortBlueprints(ctx, os.Getenv("PORTS_PATH"))
	startup.Forward(ctx, blueprints, streamer)
	r := api.Router()
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal().Err(err).Msg("cannot start HTTP server")
	}
}

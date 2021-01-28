package startup

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/tomaszczerminski/test/pkg/shared"
	"io"
	"time"
)

func Forward(ctx context.Context, ports <-chan *shared.PortBlueprint, streamer shared.StreamerClient) {
	logger := log.Ctx(ctx)
	start := time.Now()
	counter := uint64(0)
	defer func() {
		logger.Info().
			TimeDiff("duration", time.Now(), start).
			Uint64("blueprint_count", counter).
			Msg("ports forwarded")
	}()
	stream, err := streamer.Send(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot start a client-side gRPC stream")
	}
	defer func() {
		if err := stream.CloseSend(); err != nil {
			logger.Error().Err(err).Msg("cannot close gRPC client-side stream")
		}
	}()
	// (ASSUMPTION) we should not kill the process if port blueprint cannot be send.
	// it can be retried in a later time or manually
	for port := range ports {
		err := stream.Send(port)
		if err == io.EOF {
			return
		}
		if err != nil {
			logger.Warn().
				Err(err).
				Str("port_name", port.Name).
				Str("port_code", port.Code).
				Msg("cannot send port blueprint")
			continue
		}
		counter++
	}
}

package domain

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/tomaszczerminski/test/pkg/shared"
	"io"
	"time"
)

type Server struct {
	shared.UnimplementedStreamerServer
	countries CountryRepository
}

func (srv *Server) Send(stream shared.Streamer_SendServer) error {
	submitted := uint64(0)
	start := time.Now()
	defer func() {
		log.Info().
			Uint64("blueprint_count", submitted).
			TimeDiff("duration", time.Now(), start).
			Msg("port blueprints processed")
	}()
	ctx := stream.Context()
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&shared.Summary{
				Submitted: submitted,
			})
		}
		if err != nil {
			log.Error().Err(err).Msg("cannot receive message")
			return err
		}
		c, err := srv.countries.FindByName(ctx, b.Country)
		if errors.Is(err, ErrCountryNotFound) {
			c, err := NewCountry(b)
			if err != nil {
				log.Error().Err(err).Msg("cannot create a new country")
				return err
			}
			if err := srv.countries.SaveOrUpdate(ctx, c); err != nil {
				log.Error().Err(err).Msg("cannot upsert a country")
				return err
			}
			submitted++
			continue
		}
		if err != nil {
			log.Error().Err(err).Msg("cannot find country")
			return err
		}
		if _, err := c.AddPort(b); err != nil {
			log.Error().Err(err).Msg("cannot create a new port")
			return err
		}
		if err := srv.countries.SaveOrUpdate(ctx, c); err != nil {
			log.Error().Err(err).Msg("cannot upsert a port")
			return err
		}
		submitted++
	}
}

func NewServer(repository CountryRepository) *Server {
	return &Server{
		countries: repository,
	}
}

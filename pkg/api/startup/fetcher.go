package startup

import (
	"context"
	stream "github.com/bcicen/jstream"
	"github.com/rs/zerolog/log"
	"github.com/tomaszczerminski/test/pkg/shared"
	"os"
	"time"
)

// FetchPortBlueprints reads json file specified by path parameter as a stream (not eagerly).
// Then converts it to PortBlueprint value object
func FetchPortBlueprints(ctx context.Context, path string) <-chan *shared.PortBlueprint {
	logger := log.Ctx(ctx)
	f, err := os.Open(path)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot open blueprints file")
	}
	blueprints := make(chan *shared.PortBlueprint)
	go func() {
		defer close(blueprints)
		start := time.Now()
		counter := uint64(0)
		defer func() {
			logger.Info().
				TimeDiff("duration", time.Now(), start).
				Uint64("blueprint_count", counter).
				Msg("blueprints decoded")
		}()
		decoder := stream.NewDecoder(f, 0)
		for mv := range decoder.Stream() {
			parsed, ok := mv.Value.(map[string]interface{})
			if !ok {
				logger.Fatal().Caller().Msg("cannot decode port blueprint into the map")
			}
			for id, data := range parsed {
				raw, ok := data.(map[string]interface{})
				if !ok {
					logger.Fatal().Msg("cannot decode blueprint fields into the map")
				}
				blueprint := convert(id, raw)
				blueprints <- blueprint
				counter++
			}
		}
	}()
	return blueprints
}

// convert maps raw data to target struct
// There is no information provided whether there are other fields not present in the mapping
// or which of those fields are required so the result represents my current understanding of the problem
func convert(id string, parsed map[string]interface{}) *shared.PortBlueprint {
	name, _ := parsed["name"].(string)
	city, _ := parsed["city"].(string)
	code, _ := parsed["code"].(string)
	var coordinates []float64
	cc, _ := parsed["coordinates"].([]interface{})
	for _, coordinate := range cc {
		coordinates = append(coordinates, coordinate.(float64))
	}
	country, _ := parsed["country"].(string)
	province, _ := parsed["province"].(string)
	timezone, _ := parsed["timezone"].(string)
	return &shared.PortBlueprint{
		Id:          id,
		Name:        name,
		Alias:       nil,
		City:        city,
		Code:        code,
		Coordinates: coordinates,
		Country:     country,
		Province:    province,
		Regions:     nil,
		Timezone:    timezone,
		Unlocs:      nil,
	}
}

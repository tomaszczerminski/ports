package domain

import (
	"github.com/tomaszczerminski/test/pkg/shared"
	"sync"
)

// Country is an aggregate root so the only way
// to interact with all entities within the aggregate is through it
// Aggregate should always be concise, but I do not know the business rules
// so I just leave it without any (as it the only way for all the data to be processed)
type Country struct {
	sync.Mutex
	name     string
	timezone string
	cities   []*City
}

func NewCountry(blueprint *shared.PortBlueprint) (*Country, error) {
	var coordinates Coordinates
	if len(blueprint.Coordinates) == 2 {
		coordinates.x = blueprint.Coordinates[0]
		coordinates.y = blueprint.Coordinates[1]
	}
	return &Country{
		name:     blueprint.Country,
		timezone: blueprint.Timezone,
		cities: []*City{
			{
				name:     blueprint.City,
				province: blueprint.Province,
				ports: []*Port{
					{
						id:          blueprint.Id,
						name:        blueprint.Name,
						alias:       []string{},
						regions:     []string{},
						coordinates: coordinates,
						unlocs:      []string{},
						code:        blueprint.Code,
					},
				},
			},
		},
	}, nil
}

func (c *Country) AddPort(blueprint *shared.PortBlueprint) (*Port, error) {
	c.Lock()
	defer c.Unlock()
	var coordinates Coordinates
	if len(blueprint.Coordinates) == 2 {
		coordinates.x = blueprint.Coordinates[0]
		coordinates.y = blueprint.Coordinates[1]
	}
	port := &Port{
		id:          blueprint.Id,
		name:        blueprint.Name,
		code:        blueprint.Code,
		coordinates: coordinates,
	}
	for _, c := range c.cities {
		if c.name != blueprint.City {
			continue
		}
		var overwritten bool
		for i, p := range c.ports {
			if p.id == blueprint.Id {
				c.ports[i] = port
				overwritten = true
				break
			}
		}
		if !overwritten {
			c.ports = append(c.ports, port)
		}
		return port, nil
	}
	c.cities = append(c.cities, &City{
		name:     blueprint.City,
		province: blueprint.Province,
		ports:    []*Port{port},
	})
	return port, nil
}

type City struct {
	name     string
	province string
	ports    []*Port
}

type Port struct {
	id          string
	name        string
	alias       []string
	regions     []string
	coordinates Coordinates
	unlocs      []string
	code        string
}

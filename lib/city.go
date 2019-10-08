package lib

import (
	"strconv"

	"github.com/umahmood/haversine"
)

// City
type City struct {
	Name       string
	Location   haversine.Coord
	Population uint64
}

// NewCity initializes a new city.
func NewCity(name, lat, long, population string) (*City, error) {
	latF, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil, err
	}

	longF, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil, err
	}

	pop, err := strconv.ParseUint(population, 10, 64)
	if err != nil {
		return nil, err
	}

	return &City{
		Name:       name,
		Location:   haversine.Coord{Lat: latF, Lon: longF},
		Population: pop,
	}, nil

}

func (c *City) Distance(other *City) float64 {
	_, km := haversine.Distance(c.Location, other.Location)
	return km
}

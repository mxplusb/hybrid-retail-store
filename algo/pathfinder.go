package algo

import (
	"context"
	"errors"
	"sync"

	retail "github.com/mxplusb/hybrid-retail-store"
	"googlemaps.github.io/maps"
)

type Cost struct {
	Path          *Path
	CurrentWeight int
	Order         *retail.Order
	mu            sync.Mutex
}

// Determine the weight of an item.
func (c *Cost) GenerateCost(o *retail.Order) {
	c.Order = o
	c.Path = &Path{
		Destination: o.DestinationZipCode,
	}
	c.Path.GenerateItemWeightToDestination(c.Order)

	c.mu.Lock()
	c.CurrentWeight += c.Path.lastMileCost
	c.mu.Unlock()

	// TODO (mxplusb): finish this.
}

// How we determine where to go.
type Path struct {
	SourceFC     int
	Destination  int
	fcOneCost    int
	fcTwoCost    int
	lastMileCost int
	gclient      *maps.Client
}

// TODO (mxplusb): add logging.
// This will generate how much it will cost to ship the item over it's last mile.
func (p *Path) GenerateItemWeightToDestination(o *retail.Order) (bool) {
	if err := p.distanceFromFC(o); err != nil {
		return false
	}
	if p.fcOneCost < p.fcTwoCost {
		p.lastMileCost = p.fcOneCost
	} else {
		p.lastMileCost = p.fcTwoCost
	}
	return false
}

// determine how far away the destination zip code is.
func (p *Path) distanceFromFC(o *retail.Order) error {
	var err error
	p.gclient, err = maps.NewClient(nil)
	if err != nil {
		return err
	}

	resp, err := p.gclient.DistanceMatrix(context.Background(), &maps.DistanceMatrixRequest{
		Origins:      []string{string(retail.FC1Location)},
		Destinations: []string{string(o.DestinationZipCode)},
		Mode:         maps.Mode("driving"),
		Units:        maps.Units("UnitsMetric"),
	})

	// gRPC ftw... :(
	if resp.Rows[0].Elements[0].Status != "OK" {
		return errors.New("bad request to google maps")
	}
	p.fcOneCost = resp.Rows[0].Elements[0].Distance.Meters

	resp, err = p.gclient.DistanceMatrix(context.Background(), &maps.DistanceMatrixRequest{
		Origins:      []string{string(retail.FC1Location)},
		Destinations: []string{string(o.DestinationZipCode)},
		Mode:         maps.Mode("driving"),
		Units:        maps.Units("UnitsMetric"),
	})

	if resp.Rows[0].Elements[0].Status != "OK" {
		return errors.New("bad request to google maps")
	}
	p.fcTwoCost = resp.Rows[0].Elements[0].Distance.Meters
	return nil
}

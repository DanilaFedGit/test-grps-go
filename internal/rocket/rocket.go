//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/DanilaFedGit/git b/internal/rocket Store
package rocket

import "context"

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rocket Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{Store: store}
}

func (service Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := service.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}
func (service Service) InsertRocket(ctx context.Context, rocket Rocket) (Rocket, error) {
	rkt, err := service.Store.InsertRocket(rocket)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, err
}
func (service Service) DeleteRocket(id string) error {
	err := service.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}

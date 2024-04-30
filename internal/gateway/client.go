package gateway

import "github/silva4dev/golang-ms-wallet-project/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

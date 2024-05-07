package gateway

import "github/silva4dev/golang-clean-api-project/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

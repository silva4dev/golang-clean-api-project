package gateway

import "github/silva4dev/golang-ms-wallet-project/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

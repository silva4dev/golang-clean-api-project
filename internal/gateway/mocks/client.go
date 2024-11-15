package mocks

import (
	"github/silva4dev/golang-clean-api-project/internal/entity"

	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func NewClientGatewayMock() *ClientGatewayMock {
	return &ClientGatewayMock{}
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

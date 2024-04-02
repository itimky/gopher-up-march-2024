package ordersdb_test

import (
	"context"
	"testing"

	ordersdb "github.com/itimky/gopher-up-march-2024/pkg/orders/db"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
	"github.com/stretchr/testify/suite"
)

type InMemDBSuite struct {
	suite.Suite
}

func (s *InMemDBSuite) Test_AddOrder() {
	testArgs := []struct {
		name        string
		params      ordersdomain.AddOrderParams
		expectedErr error
	}{
		{
			name: "ok",
			params: ordersdomain.AddOrderParams{
				ProductID: "1",
				UserID:    "2",
				OrderID:   "3",
				Quantity:  4,
			},
		},
	}

	for _, tt := range testArgs {
		s.Run(tt.name, func() {
			inMemDB := ordersdb.NewInMemDB()

			err := inMemDB.AddOrder(context.Background(), tt.params)

			s.ErrorIs(tt.expectedErr, err)
		})
	}
}

func TestInMemDBSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(InMemDBSuite))
}

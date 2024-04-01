package ordersdomain_test

import (
	"context"
	"fmt"
	"testing"

	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
	"github.com/itimky/gopher-up-march-2024/pkg/sys"
	mocks "github.com/itimky/gopher-up-march-2024/test/pkg/orders/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type OrdersSuite struct {
	suite.Suite

	dbMock *mocks.Mockdb
	orders *ordersdomain.Orders
}

func (s *OrdersSuite) SetupSuite() {
	s.dbMock = mocks.NewMockdb(s.T())

	s.orders = ordersdomain.NewOrders(s.dbMock)
}

func (s *OrdersSuite) Test_CreateOrder() {
	testErr := fmt.Errorf("test error")
	orderID := "3"

	testArgs := []struct {
		name           string
		params         ordersdomain.CreateOrderParams
		expectedResult *ordersdomain.CreateOrderResult
		expectedErr    error
		dbParams       *ordersdomain.AddOrderParams
		dbErr          error
	}{
		{
			name: "err: db error",
			params: ordersdomain.CreateOrderParams{
				ProductID: "1",
				UserID:    "2",
				Quantity:  4,
			},
			expectedErr: testErr,
			dbParams: &ordersdomain.AddOrderParams{
				ProductID: "1",
				UserID:    "2",
				OrderID:   "3",
				Quantity:  4,
			},
			dbErr: testErr,
		},
		{
			name: "ok",
			params: ordersdomain.CreateOrderParams{
				ProductID: "1",
				UserID:    "2",
				Quantity:  4,
			},
			expectedResult: &ordersdomain.CreateOrderResult{
				OrderID: "3",
			},
			dbParams: &ordersdomain.AddOrderParams{
				ProductID: "1",
				UserID:    "2",
				OrderID:   "3",
				Quantity:  4,
			},
		},
	}

	for _, tt := range testArgs {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.dbParams != nil {
				s.dbMock.EXPECT().AddOrder(mock.Anything, *tt.dbParams).Return(tt.dbErr).Once()
			}

			orders := ordersdomain.NewOrders(s.dbMock)

			ctx := sys.RandomIDToCtx(context.Background(), orderID)
			result, err := orders.CreateOrder(ctx, tt.params)

			s.ErrorIs(err, tt.expectedErr)
			s.EqualValues(tt.expectedResult, result)
		})
	}
}

func TestOrdersSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(OrdersSuite))
}

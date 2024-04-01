package ordershttp_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	contractshttp "github.com/itimky/gopher-up-march-2024/pkg/contracts/http"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
	ordershttp "github.com/itimky/gopher-up-march-2024/pkg/orders/http"
	"github.com/itimky/gopher-up-march-2024/test"
	mocks "github.com/itimky/gopher-up-march-2024/test/pkg/orders/http"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite

	ordersMock *mocks.Mockorders
	handler    *ordershttp.Handler
}

func (s *HandlerSuite) SetupTest() {
	s.ordersMock = mocks.NewMockorders(s.T())
	s.handler = ordershttp.NewHandler(s.ordersMock)
}

func (s *HandlerSuite) Test_CreateOrderV1() {
	testErr := fmt.Errorf("test")

	testArgs := []struct {
		name             string
		params           []byte
		expectedCode     int
		expectedResponse []byte
		ordersParams     *ordersdomain.CreateOrderParams
		ordersResult     *ordersdomain.CreateOrderResult
		ordersErr        error
	}{
		{
			name:             "err: bad request",
			params:           []byte(`{`),
			expectedCode:     http.StatusBadRequest,
			expectedResponse: []byte("unexpected EOF"),
		},
		{
			name: "err: orders err",
			params: test.MustMarshalJSON(s.T(), contractshttp.CreateOrderRequestV1{
				ProductID: "1",
				UserID:    "2",
				Quantity:  3,
			}),
			expectedCode:     http.StatusInternalServerError,
			expectedResponse: []byte("test"),
			ordersParams: &ordersdomain.CreateOrderParams{
				ProductID: "1",
				UserID:    "2",
				Quantity:  3,
			},
			ordersErr: testErr,
		},
		{
			name: "ok",
			params: test.MustMarshalJSON(s.T(), contractshttp.CreateOrderRequestV1{
				ProductID: "1",
				UserID:    "2",
				Quantity:  3,
			}),
			expectedCode: http.StatusCreated,
			expectedResponse: test.MustMarshalJSON(s.T(), contractshttp.CreateOrderResponseV1{
				OrderID: "4",
			}),
			ordersParams: &ordersdomain.CreateOrderParams{
				ProductID: "1",
				UserID:    "2",
				Quantity:  3,
			},
			ordersResult: &ordersdomain.CreateOrderResult{
				OrderID: "4",
			},
		},
	}

	for _, tt := range testArgs {
		s.Run(tt.name, func() {
			if tt.ordersParams != nil {
				s.ordersMock.EXPECT().
					CreateOrder(mock.Anything, *tt.ordersParams).
					Return(tt.ordersResult, tt.ordersErr).Once()
			}

			req := httptest.NewRequest(http.MethodPost, "/orders/v1", bytes.NewReader(tt.params))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			s.handler.Mux().ServeHTTP(rr, req)

			s.Equal(tt.expectedCode, rr.Code)
			s.EqualValues(string(tt.expectedResponse)+"\n", rr.Body.String())
		})
	}
}

func TestHandlerSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(HandlerSuite))
}

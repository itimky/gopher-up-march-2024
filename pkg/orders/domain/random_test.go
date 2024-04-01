package ordersdomain_test

import (
	"context"
	"testing"

	"github.com/itimky/gopher-up-march-2024/pkg"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
	"github.com/stretchr/testify/suite"
)

const (
	testOrderID pkg.OrderID = "4436567788" // generated from seed 0
)

type RandomSuite struct {
	suite.Suite
}

func (s *RandomSuite) Test_GenerateRandomID_OK() {
	randomID := ordersdomain.GenerateRandomID(context.Background())
	s.Len(randomID, ordersdomain.RandomIDDigitCount)
}

func (s *RandomSuite) Test_GenerateRandomID_NonRepeat() {
	randomID1 := ordersdomain.GenerateRandomID(context.Background())
	randomID2 := ordersdomain.GenerateRandomID(context.Background())
	s.NotEqualValues(randomID1, randomID2)
}

func (s *RandomSuite) Test_GenerateRandomID_FromCtx() {
	ctx := context.WithValue(context.Background(), ordersdomain.RandomSeedKey{}, int64(0))
	generatedRandomID := ordersdomain.GenerateRandomID(ctx)

	s.EqualValues(testOrderID, generatedRandomID)
}

func TestRandomSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RandomSuite))
}

package sys_test

import (
	"context"
	"testing"

	"github.com/itimky/gopher-up-march-2024/pkg/sys"
	"github.com/stretchr/testify/suite"
)

type RandomSuite struct {
	suite.Suite
}

func (s *RandomSuite) Test_GenerateRandomID_OK() {
	randomID, err := sys.GenerateRandomID(context.Background())
	s.Require().NoError(err)

	s.Len(randomID, sys.RandomIDDigitCount)
}

func (s *RandomSuite) Test_GenerateRandomID_NonRepeat() {
	randomID1, err := sys.GenerateRandomID(context.Background())
	s.Require().NoError(err)

	randomID2, err := sys.GenerateRandomID(context.Background())
	s.Require().NoError(err)

	s.NotEqualValues(randomID1, randomID2)
}

func (s *RandomSuite) Test_GenerateRandomID_FromCtx() {
	randomID := "test"
	ctx := sys.RandomIDToCtx(context.Background(), randomID)

	generatedRandomID, err := sys.GenerateRandomID(ctx)
	s.Require().NoError(err)

	s.EqualValues(randomID, generatedRandomID)
}

func TestRandomSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RandomSuite))
}

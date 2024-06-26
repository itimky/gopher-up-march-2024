package test_test

import (
	"testing"

	"github.com/itimky/gopher-up-march-2024/test"
	"github.com/stretchr/testify/suite"
)

type ConvertSuite struct {
	suite.Suite
}

func (s *ConvertSuite) Test_MustMarshalJSON_Panic() {
	s.Panics(func() {
		test.MustMarshalJSON(s.T(), make(chan struct{}))
	})
}

func (s *ConvertSuite) Test_MustMarshalJSON_OK() {
	var result []byte

	s.NotPanics(func() {
		result = test.MustMarshalJSON(s.T(), struct {
			Valid string `json:"valid"`
		}{
			Valid: "json",
		})
	})

	s.EqualValues([]byte(`{"valid":"json"}`), result)
}

func (s *ConvertSuite) Test_IntPtr() {
	value := 42

	result := test.IntPtr(value)

	s.Equal(&value, result)
}

func TestConvertSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ConvertSuite))
}

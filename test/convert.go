package test

import (
	"encoding/json"
	"testing"
)

func MustMarshalJSON(t *testing.T, value interface{}) []byte {
	t.Helper()

	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return data
}

func IntPtr(v int) *int {
	return &v
}

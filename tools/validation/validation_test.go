package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStruct struct {
	s string
	i int
}

func TestValidateInputNotEmpty(t *testing.T) {
	mockDataNotValid := []mockStruct{
		{
			s: "",
			i: 0,
		},
		{
			s: "",
			i: 1,
		},
		{
			s: "string",
			i: 0,
		},
		{
			s: "string",
			i: 1,
		},
	}
	for j, mock := range mockDataNotValid {
		if j == len(mockDataNotValid)-1 {
			err := ValidateInputNotEmpty(mock.s, mock.i)
			assert.Nil(t, err)
		} else {
			err := ValidateInputNotEmpty(mock.s, mock.i)
			assert.NotNil(t, err)
		}
	}
}

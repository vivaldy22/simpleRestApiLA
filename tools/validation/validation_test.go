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
	t.Run("", func(t *testing.T) {
		// mockDataValid := mockStruct{s: "string", i: 1, p: &i}
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
		}
		for _, mock := range mockDataNotValid {
			err := ValidateInputNotEmpty(mock.s, mock.i)
			assert.NotNil(t, err)
		}
	})
}

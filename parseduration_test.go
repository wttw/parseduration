package parseduration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validTests = []struct {
	in  string
	out string
}{
	{"P0Y", "0s"},
	{"PT0S", "0s"},
	{"PT2H", "2h0m0s"},
	{"PT6H10M", "6h10m0s"},
	{"P1Y3MT6M", "10920h6m0s"},
}

func Test8601(t *testing.T) {
	_, err := Parse8601("Pnotvalid")
	assert.Error(t, err, "invalid duration")
	for _, tt := range validTests {
		d, err := Parse8601(tt.in)
		assert.NoError(t, err, "no error %s", tt.in)
		assert.Equal(t, tt.out, d.String(), "input %s", tt.in)
	}
}

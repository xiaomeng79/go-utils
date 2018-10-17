package path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testValues = []struct {
		in string
	}{
		{"example.file"},
		{"example.json"},
		{"example/example.file"},
	}
)

func TestPathExists(t *testing.T) {

	for _, v := range testValues {
		res, err := PathExists(v.in)
		if v.in == "example.file" {
			assert.Truef(t, res, "file:%s", v.in)
			assert.NoErrorf(t, err, "file:%s", v.in)
		} else {
			assert.Falsef(t, res, "file:%s", v.in)
			assert.NoErrorf(t, err, "file:%s", v.in)
		}

	}
}

package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	as := assert.New(t)

	data := []struct {
		in     string
		expect string
	}{
		{in: "あいうえお", expect: "杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿柿杮柿柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿柿杮杮柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿杮柿柿柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿杮柿杮柿"},
		{in: "abc", expect: "柿杮杮柿柿柿柿杮柿杮杮柿柿柿杮柿柿杮杮柿柿柿杮杮"},
	}

	for _, v := range data {
		res, err := Encode(v.in)
		as.NoError(err)
		as.Equal(v.expect, res)
	}
}

func TestEncodeError(t *testing.T) {
	_, err := Encode("")
	assert.Error(t, err)
}

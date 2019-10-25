package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	data := []struct {
		in     string
		expect string
	}{
		{expect: "あいうえお", in: "杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿柿杮柿柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿柿杮杮柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿杮柿柿柿杮杮杮柿柿柿杮杮杮柿柿柿柿柿柿杮杮柿柿柿杮柿杮柿"},
		{expect: "abc", in: "柿杮杮柿柿柿柿杮柿杮杮柿柿柿杮柿柿杮杮柿柿柿杮杮"},
	}
	as := assert.New(t)
	for _, v := range data {
		res, err := Decode(v.in)
		as.NoError(err)
		as.Equal(v.expect, res)
	}
}

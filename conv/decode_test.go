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
		{in: "柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿柿柿柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿柿柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿柿柿杮杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿柿杮柿杮柿", expect: "あいうえお"},
		{in: "柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿柿柿杮杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿杮柿柿杮柿柿杮杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿柿杮柿杮柿柿杮杮", expect: "うんこ"},
		{in: "柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿杮杮杮杮杮杮杮杮杮杮杮", expect: "柿"},
		{in: "柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿柿杮杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿杮柿杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿杮杮柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿柿杮杮柿柿杮杮杮", expect: "abcdefg"},
	}
	as := assert.New(t)
	for _, v := range data {
		as.Equal(v.expect, Decode(v.in))
	}
}
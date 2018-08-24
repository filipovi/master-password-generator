package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	expected := "WQ,fH/LulHFCx%URrl_nng2p*y6OUH(;"
	p, _ := NewPassword("name", "passphrase", "service", 32, 1, "my scope")
	assert.Equal(t, expected, p)
	assert.Equal(t, 32, len(p))

	expected = "(pnpwSz`X&:yfb~pFc_ um.WOR46*mp7By}"
	p, _ = NewPassword("name", "passphrase", "service", 35, 2, "my new scope")
	assert.Equal(t, expected, p)
	assert.Equal(t, 35, len(p))
}

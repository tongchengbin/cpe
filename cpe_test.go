package cpe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCPE(t *testing.T) {
	c, err := Parse("cpe:2.3:a:hikvision:hikvision_dvr:*:*:*:*:*:*:*:*")
	assert.Nil(t, err)
	assert.Equal(t, c.Vendor, "hikvision")
	assert.Equal(t, c.Product, "hikvision_dvr")
}

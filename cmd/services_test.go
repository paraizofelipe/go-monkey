package cmd

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
)

var (
	ks *Service
)

func TestListInfo(t *testing.T) {
	kongConfig := []interface{}{
		map[string]interface{}{"url": "http://localhost:8001"},
	}
	viper.Set("kong.host", kongConfig)
	svc := ks.ListInfo()
	assert.NotNil(t, svc)
	assert.IsType(t, []Service{}, svc)
}

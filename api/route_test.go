package api

import (
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type SuiteRoute struct {
	suite.Suite
	svc *Service
}

func (s *SuiteRoute) SetupTest() {
	home, _ := homedir.Dir()
	viper.AddConfigPath(home)
	viper.SetConfigName(".go-monkey")
	err := viper.ReadInConfig()
	s.Nil(err)
}

func (s *SuiteRoute) TestRoute_List() {
	api := New("http://localhost:8001")
	err, routes := api.ListRoutes()

	s.Nil(err)
	s.IsType(&[]Route{}, routes)
}

func (s *SuiteRoute) TestRoute_CreateRoute() {
	api := New("http://localhost:8001")
	route := Route{
		Name:      "route_to_route",
		Protocols: []string{"http", "https"},
		Methods:   []string{"GET", "POST"},
		Hosts:     []string{"example.com", "foo.test"},
		Paths:     []string{"/foo", "/bar"},
	}

	err := api.CreateRoute(route)
	s.Nil(err)
}

func TestSuiteRoute(t *testing.T) {
	suite.Run(t, new(SuiteRoute))
}

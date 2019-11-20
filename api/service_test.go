package api

import (
	"testing"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/suite"
)

type SuiteService struct {
	suite.Suite
	svc *Service
}

func (s *SuiteService) SetupTest() {
	home, _ := homedir.Dir()
	viper.AddConfigPath(home)
	viper.SetConfigName(".go-monkey")
	err := viper.ReadInConfig()
	s.Nil(err)
}

func (s *SuiteService) TestService_List() {
	api := New("http://localhost:8001")
	err, services := api.Services()

	s.Nil(err)
	s.IsType(&[]Service{}, services)
}

func (s *SuiteService) TestService_CreateService() {
	api := New("http://localhost:8001")
	service := Service{
		Host:     "test.com",
		Protocol: "http",
		Name:     "teste",
		Port:     80,
		Path:     "/test",
	}

	err := api.CreateServices(service)
	s.Nil(err)
}

func TestSuiteService(t *testing.T) {
	suite.Run(t, new(SuiteService))
}

package translation_test

import (
	"testing"

	"github.com/luzeduardo/shipping-go/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestRemoteServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteServiceTestSuite))
}

type RemoteServiceTestSuite struct {
	suite.Suite
	client           *MockHelloClient
	underTestService *translation.RemoteService
}

type MockHelloClient struct {
	mock.Mock
}

func (suite *RemoteServiceTestSuite) SetupTest() {
	suite.client = new(MockHelloClient)
	suite.underTestService = translation.NewRemoteService(suite.client)
}

func (m *MockHelloClient) Translate(word string, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *RemoteServiceTestSuite) TestTranslate() {
	//org
	suite.client.On("Translate", "foo", "bar").Return("baz", nil)
	//act
	res := suite.underTestService.Translate("foo", "bar")

	suite.Equal(res, "baz")
	suite.client.AssertExpectations(suite.T())
}

func (suite *RemoteServiceTestSuite) TestTranslate_CaseSensitive() {
	suite.client.On("Translate", "foo", "bar").Return("baz", nil)

	res := suite.underTestService.Translate("foo", "bar")

	suite.Equal(res, "baz")
	suite.client.AssertExpectations(suite.T())
}

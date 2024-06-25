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
	client    *MockHelloClient
	underTest *translation.RemoteService
}

type MockHelloClient struct {
	mock.Mock
}

func (suite *RemoteServiceTestSuite) SetupTest() {
	suite.client = new(MockHelloClient)
	suite.underTest = translation.NewRemoteService(suite.client)
}

func (m *MockHelloClient) Translate(word string, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *RemoteServiceTestSuite) TestTranslate() {
	//org
	suite.client.On("Translate", "foo", "bar").Return("baz", nil)
	//act
	res, _ := suite.underTest.Translate("foo", "bar")

	suite.Equal(res, "baz")
	suite.client.AssertExpectations(suite.T())
}

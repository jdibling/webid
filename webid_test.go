package webid

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type meetApiTestSuite struct {suite.Suite}

func TestPkg_webid(t *testing.T) {
	suite.Run(t, new(meetApiTestSuite))
}

func (s *meetApiTestSuite) TestNew() {
	id, err := New()
	s.Require().NoError(err)
	s.Require().NotEmpty(id)
}

func (s *meetApiTestSuite) TestMustNew() {
	id := MustNew()
	s.Require().NotEmpty(id)
}
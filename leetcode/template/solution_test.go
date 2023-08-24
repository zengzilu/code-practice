package template

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SolutionTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SolutionTestSuite))
}

func (s *SolutionTestSuite) Test_01() {}

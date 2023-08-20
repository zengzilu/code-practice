package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SolutionTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SolutionTestSuite))
}

func (s *SolutionTestSuite) Test_01() {
	r := diffWaysToCompute("1")
	assert.True(s.T(), len(r) == 1)
	assert.True(s.T(), r[0] == 1)

}

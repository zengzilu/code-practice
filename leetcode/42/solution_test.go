package _40

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
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(s.T(), trap(arr), 6)
}

func (s *SolutionTestSuite) Test_02() {
	arr := []int{1, 2}
	assert.Equal(s.T(), trap(arr), 0)
}

func (s *SolutionTestSuite) Test_03() {
	arr := []int{3, 1, 2}
	assert.Equal(s.T(), trap(arr), 1)
}

func (s *SolutionTestSuite) Test_04() {
	arr := []int{3, 1, 2, 4}
	assert.Equal(s.T(), trap(arr), 3)
}

func (s *SolutionTestSuite) Test_05() {
	arr := []int{3, 1, 2, 0, 10}
	assert.Equal(s.T(), trap(arr), 6)
}

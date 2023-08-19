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
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	target := 20

	assert.False(s.T(), searchMatrix(matrix, target))
}

func (s *SolutionTestSuite) Test_02() {
	matrix := [][]int{{1}}
	target := 2

	assert.False(s.T(), searchMatrix(matrix, target))
}

func (s *SolutionTestSuite) Test_03() {
	matrix := [][]int{{1}}
	target := 1

	assert.True(s.T(), searchMatrix(matrix, target))
}

func (s *SolutionTestSuite) Test_04() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
	}
	target := 4

	assert.True(s.T(), searchMatrix(matrix, target))
}

func (s *SolutionTestSuite) Test_05() {
	matrix := [][]int{
		{1},
		{2},
		{3},
		{10},
		{18}}
	target := 18

	assert.True(s.T(), searchMatrix(matrix, target))
}

package _07

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
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1}}
	assert.Equal(s.T(), 4, trapRainWater(heightMap))
}

func (s *SolutionTestSuite) Test_02() {
	heightMap := [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3}}
	assert.Equal(s.T(), 10, trapRainWater(heightMap))
}

func (s *SolutionTestSuite) Test_03() {
	heightMap := [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3}}
	assert.Equal(s.T(), 0, trapRainWater(heightMap))
}

func (s *SolutionTestSuite) Test_04() {
	heightMap := [][]int{
		{3}}
	assert.Equal(s.T(), 0, trapRainWater(heightMap))
}

func (s *SolutionTestSuite) Test_05() {
	heightMap := [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 3, 2, 3},
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3}}
	assert.Equal(s.T(), 5, trapRainWater(heightMap))
}

func (s *SolutionTestSuite) Test_06() {
	heightMap := [][]int{
		{3, 3, 3, 3, 3},
		{3, 4, 4, 4, 3},
		{3, 4, 3, 4, 3},
		{3, 4, 4, 4, 3},
		{3, 3, 3, 3, 3}}
	assert.Equal(s.T(), 1, trapRainWater(heightMap))
}

func (s *SolutionTestSuite) Test_07() {
	heightMap := [][]int{
		{3, 3, 3, 3, 3},
		{3, 4, 4, 4, 3},
		{3, 4, 5, 4, 3},
		{3, 4, 4, 4, 3},
		{3, 3, 3, 3, 3}}
	assert.Equal(s.T(), 0, trapRainWater(heightMap))
}
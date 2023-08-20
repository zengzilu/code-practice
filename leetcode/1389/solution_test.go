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
	nums := []int{0, 1, 2, 3, 4}
	indexes := []int{0, 1, 2, 2, 1}
	result := createTargetArray(nums, indexes)

	expect := []int{0, 4, 1, 3, 2}

	for i, v := range result {
		assert.Equal(s.T(), expect[i], v)
	}

}

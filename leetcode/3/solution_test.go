package _3

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

	assert.Equal(s.T(), 1, lengthOfLongestSubstring("a"))
	assert.Equal(s.T(), 1, lengthOfLongestSubstring("aa"))
	assert.Equal(s.T(), 4, lengthOfLongestSubstring("abcda"))
	assert.Equal(s.T(), 4, lengthOfLongestSubstring("abcabcdabc"))
	assert.Equal(s.T(), 3, lengthOfLongestSubstring("abac"))
	assert.Equal(s.T(), 4, lengthOfLongestSubstring("abcd"))
	assert.Equal(s.T(), 0, lengthOfLongestSubstring(""))
}

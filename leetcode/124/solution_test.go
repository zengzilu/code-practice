package _24

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
	rootNode := &TreeNode{
		Var:   1,
		Left:  nil,
		Right: nil,
	}

	assert.Equal(s.T(), 1, maxPathSum(rootNode))
}

func (s *SolutionTestSuite) Test_02() {
	rootNode := &TreeNode{
		Var:   1,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left = &TreeNode{
		Var:   -2,
		Left:  nil,
		Right: nil,
	}

	assert.Equal(s.T(), 1, maxPathSum(rootNode))
}

func (s *SolutionTestSuite) Test_03() {
	rootNode := &TreeNode{
		Var:   -1,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left = &TreeNode{
		Var:   10,
		Left:  nil,
		Right: nil,
	}

	assert.Equal(s.T(), 10, maxPathSum(rootNode))
}

func (s *SolutionTestSuite) Test_04() {
	rootNode := &TreeNode{
		Var:   1,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left = &TreeNode{
		Var:   2,
		Left:  nil,
		Right: nil,
	}

	rootNode.Right = &TreeNode{
		Var:   3,
		Left:  nil,
		Right: nil,
	}

	assert.Equal(s.T(), 6, maxPathSum(rootNode))
}

func (s *SolutionTestSuite) Test_05() {
	rootNode := &TreeNode{
		Var:   1,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left = &TreeNode{
		Var:   2,
		Left:  nil,
		Right: nil,
	}

	rootNode.Right = &TreeNode{
		Var:   -3,
		Left:  nil,
		Right: nil,
	}

	assert.Equal(s.T(), 3, maxPathSum(rootNode))
}

func (s *SolutionTestSuite) Test_06() {
	rootNode := &TreeNode{
		Var:   -1,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left = &TreeNode{
		Var:   1,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left.Left = &TreeNode{
		Var:   2,
		Left:  nil,
		Right: nil,
	}

	rootNode.Left.Right = &TreeNode{
		Var:   3,
		Left:  nil,
		Right: nil,
	}

	assert.Equal(s.T(), 6, maxPathSum(rootNode))
}

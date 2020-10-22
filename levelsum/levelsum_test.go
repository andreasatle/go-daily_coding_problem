package levelsum_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/levelsum"
	"github.com/stretchr/testify/assert"
)

func TestLevelsumCase1(t *testing.T) {
	// Intermediate level gives smallest sum
	tree := &levelsum.Node{32,
		&levelsum.Node{15,
			&levelsum.Node{6, nil, nil},
			&levelsum.Node{8, nil, nil},
		},
		&levelsum.Node{12,
			&levelsum.Node{8, nil, nil},
			&levelsum.Node{9, nil, nil},
		},
	}
	maxLevel, _ := tree.LevelOfMinSum()
	assert.Equal(t, 1, maxLevel, "Intermediate level should get smallest sum")
}

func TestLevelsumCase2(t *testing.T) {
	// Leafs gives smallest sum
	tree := &levelsum.Node{32,
		&levelsum.Node{15,
			&levelsum.Node{3, nil, nil},
			&levelsum.Node{8, nil, nil},
		},
		&levelsum.Node{11,
			&levelsum.Node{8, nil, nil},
			&levelsum.Node{5, nil, nil},
		},
	}
	maxLevel, _ := tree.LevelOfMinSum()
	assert.Equal(t, 2, maxLevel, "leafs should get smallest sum")
}

func TestLevelsumCase3(t *testing.T) {
	// Root gives smallest sum
	tree := &levelsum.Node{25,
		&levelsum.Node{15,
			&levelsum.Node{8, nil, nil},
			&levelsum.Node{8, nil, nil},
		},
		&levelsum.Node{11,
			&levelsum.Node{8, nil, nil},
			&levelsum.Node{9, nil, nil},
		},
	}
	maxLevel, _ := tree.LevelOfMinSum()
	assert.Equal(t, 0, maxLevel, "root should give smallest sum")
}

func TestLevelsumCase4(t *testing.T) {
	// Empty tree
	var tree *levelsum.Node
	_, err := tree.LevelOfMinSum()
	assert.Equal(t, errors.New("empty tree"), err, "tree should be empty")
}

func TestLevelsumCase5(t *testing.T) {
	// Root gives smallest sum
	tree := &levelsum.Node{25,
		&levelsum.Node{15,
			nil,
			&levelsum.Node{8, nil, nil},
		},
		&levelsum.Node{11,
			&levelsum.Node{8, nil, nil},
			&levelsum.Node{9, nil, nil},
		},
	}
	maxLevel, _ := tree.LevelOfMinSum()
	assert.Equal(t, 0, maxLevel, "root should give smallest sum")
}

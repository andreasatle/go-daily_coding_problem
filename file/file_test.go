package file_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/file"
	"github.com/stretchr/testify/assert"
)

func TestLruCache(t *testing.T) {
	file1 := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	file2 := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	assert.Equal(t, 20, file.GetMaxPathLength(file1), "Wrong length of largest path.")
	assert.Equal(t, 32, file.GetMaxPathLength(file2), "Wrong length of largest path.")
}

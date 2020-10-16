// Package allperm contains Daily Coding Problem #96
//
// This problem was asked by Microsoft.
//
// Given a number in the form of a list of digits,
// return all possible permutations.
//
// For example, given [1,2,3], return
// [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]].
//
// Author: Andreas Atle, atle.andreas@gmail.com
package allperm

// Permutations returns all permutations of an integer slice.
func Permutations(vals []int) [][]int {
	return allPerms(vals, []int{}, [][]int{})
}

func allPerms(vals []int, perm []int, perms [][]int) [][]int {
	if len(vals) == 0 {
		return append(perms, perm)
	}
	vals2 := make([]int, len(vals)-1)

	for i, val := range vals {
		removeElem(vals2, vals, i)
		perms = allPerms(vals2, append(perm, val), perms)
	}
	return perms
}

func removeElem(slice2, slice []int, index int) {
	copy(slice2, slice[:index])
	copy(slice2[index:], slice[index+1:])
}

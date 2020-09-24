// Package prereq contains Daily Coding Problem #92
// This problem was asked by Airbnb.
// We're given a hashmap associating each courseId key with a
// list of courseIds values, which represents that the prerequisites
// of courseId are courseIds. Return a sorted ordering of courses
// such that we can finish all courses.
// Return null if there is no such ordering.
// For example, given
//  {
//      'CSC300': ['CSC100', 'CSC200'],
//      'CSC200': ['CSC100'],
//      'CSC100': []
//  },
// should return
//  ['CSC100', 'CSC200', 'CSCS300'].
//
// Author: Andreas Atle, atle.andreas@gmail.com
package prereq

// Empty defines an empty struct
type Empty struct{}

// StringSet defines a set of strings
type StringSet map[string]Empty

// StringSetMap defines a map from strings to a set of strings
type StringSetMap map[string]StringSet

// StringSliceMap defines a map from strings to a slice of strings
type StringSliceMap map[string][]string

// GetOrderedList returns an ordered slice of events that fulfills the prerequisites
func GetOrderedList(m StringSliceMap) []string {
	fwdM, bwdM := initMaps(m)
	return findList(fwdM, bwdM)
}

// initMaps defines the local work-maps. The bwdM simplifize deletion in the fwdM.
func initMaps(m StringSliceMap) (StringSetMap, StringSetMap) {
	fwdM := StringSetMap{}
	bwdM := StringSetMap{}
	for k, value := range m {
		fwdM[k] = StringSet{}
		for _, v := range value {
			fwdM[k][v] = Empty{}
			if _, ok := bwdM[v]; !ok {
				bwdM[v] = StringSet{}
			}
			bwdM[v][k] = Empty{}
		}
	}
	return fwdM, bwdM
}

// findList finds the ordered list of events.
// This could (should?) have been made recursive!
func findList(fwdM, bwdM map[string]StringSet) []string {
	output := []string{}
	flag := true
	for flag {
		flag = false
		for k, value := range fwdM {
			// Check if events have no prerequisites
			if len(value) == 0 {
				flag = true
				// Add event to output
				output = append(output, k)

				// Delete event from fwd and bwd maps
				delete(fwdM, k)
				for v2 := range bwdM[k] {
					delete(fwdM[v2], k)
				}
				delete(bwdM, k)
			}
		}
	}
	// Returns nil ([]) if the forward map is not empty (cyclic graph)
	if len(fwdM) > 0 {
		return nil
	}

	// Return the ordered list of events
	return output
}

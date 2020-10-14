# go-daily_coding_problem
Some problems from dailycodingproblem.com are solved using Golang.

Some effort have been made to make the documentation look good with godoc.

All problems (except number 10) has a test-suite. Run the tests with "go test ./..." in the repository root-directory.

My two favorite solutions are number 2 and 40. Number 2 computes the product of all numbers in a slice, except
for the current entry. This is trivial as long as there are no zero entries. For zero entries it becomes quite interesting. I solved the problem with a divide and conquer algorithm. Problem number 40 has a slice with triplets except for one single entry. The task is to find this single entry. I use some interesting modulo-arithmetics, that will be slow in reality, but still fulfil the complexity requirement, let be with a large constant when the entries are large.
|Number|Package|Description|
|-----:|:-----:|-----------|
|1|sumk|Given a list of numbers and a number k, return whether any two numbers from the list add up to k.|
|2|exprod|Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.|
|3|strtree|Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.|
|4|missing|Given an array of integers, find the first missing positive integer in linear time and constant space.|
|5|cons|cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4. Implement car and cdr (given implementation of cons, see doc.).|
|6|xor|Implement an XOR linked list; it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.|
|7|encode|Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.|
|8|unival|A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value. Given the root to a binary tree, count the number of unival subtrees.|
|9|nonadj|Given a list of integers, write a function that returns the largest sum of non-adjacent numbers.|
|10|schedule|Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.|
|12|stair|There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time. Given N, write a function that returns the number of unique ways you can climb the staircase.|
|17|file|Given a string representing the file system in the above format, return the length of the longest absolute path to a file in the abstracted file system.|
|25|regexp|Implement regular expression matching with the following special characters: . (period) which matches any single character * (asterisk) which matches zero or more of the preceding element.|
|35|rgb|Given an array of strictly the characters 'R', 'G', and 'B', segregate the values of the array so that all the Rs come first, the Gs come second, and the Bs come last. You can only swap elements of the array.|
|38|queen|You have an N by N board. Write a function that, given N, returns the number of possible arrangements of the board where N queens can be placed on the board without threatening each other, i.e. no two queens share the same row, column, or diagonal.|
|40|once|Given an array of integers where every integer occurs three times except for one integer, which only occurs once, find and return the non-duplicated integer.|
|52|lru|Implement an LRU (Least Recently Used) Cache.|
|64|knight|Given N, write a function to return the number of knight's tours on an N by N chessboard.|
|69|product|Given a list of integers, return the largest product that can be made by multiplying any three integers.|
|72|path|Given a graph with n nodes and m directed edges, return the largest value path of the graph. If the largest value is infinite, then return null.|
|75|subseq|Given an array of numbers, find the length of the longest increasing subsequence in the array.|
|83|inverttree|Invert a binary tree.|
|84|islands|Given a Matrix of 1s and 0s, return the number of "islands" in the matrix.|
|85|mathor|Given three 32-bit integers x, y, and b, return x if b is 1 and y if b is 0, using only mathematical or bit operations. You can assume b can only be 1 or 0.|
|89|validtree|Determine whether a tree is a valid binary search tree.|
|90|rnd|Given an integer n and a list of integers l, write a function that randomly generates a number from 0 to n-1 that isn't in l (uniform).|
|91|lambda|What does the below code snippet print out? How can we fix the anonymous functions to behave as we'd expect?|
|92|prereq|We're given a hashmap associating each courseId key with a list of courseIds values, which represents that the prerequisites of courseId are courseIds. Return a sorted ordering of courses such that we can finish all courses.|
|96|allperm|Given a number in the form of a list of digits, return all possible permutations.|
|97|recent|Write a map implementation with a get function that lets you retrieve the value of a key at a particular time.|
|98|letters|Given a 2D board of characters and a word, find if the word exists in the grid. The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.|
|100|minpath|You are given a sequence of points and the order in which you need to cover the points. Give the minimum number of steps in which you can achieve it. You start from the first point.|
|101|sumprimes|Given an even number (greater than 2), return two prime numbers whose sum will be equal to the given number.|
|102|cont|Given a list of integers and a number K, return which contiguous elements of the list sum to K.|
|103|substr|Given a string and a set of characters, return the shortest substring containing all the characters in the set.|
|104|palin|Determine whether a doubly linked list is a palindrome. What if it’s singly linked?|
|106|jumps|Given an integer list where each number represents the number of hops you can make, determine whether you can reach to the last index starting at index 0.|
|107|levelwise|Print the nodes in a binary tree level-wise.|
|108|stringshift|Given two strings A and B, return whether or not A can be shifted some number of times to get B.|
|109|bitswap|Given an unsigned 8-bit integer, swap its even and odd bits.|
|110|leaf|Given a binary tree, return all paths from the root to leaves.|
|111|anagram|Given a word W and a string S, find all starting indices in S which are anagrams of W.|

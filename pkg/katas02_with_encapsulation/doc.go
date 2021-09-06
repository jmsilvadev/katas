// Copyright 2021 JMSilvaDev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package katas02withencapsulation is a library that contains methods to solve the Kata02 problem, Karate Chop, see: http://codekata.com/kata/kata02-karate-chop/
// This library uses concepts of OOP and SOLID to implement best practices of programming
/* Usage:
1 - Instantiate the object of binary search and call the Chop method:
	k := kataswithinternals.NewIterable()
	pos := k.Chop(5, []int{1, 3, 5})
2 - Types of binary search:
	iterable: k := kataswithinternals.NewIterable()
	recursive: k := kataswithinternals.NewRecursive()
	recursivereference: k := kataswithinternals.NewRecursiveReference()
	slicerecursive: k := kataswithinternals.NewRecursiveSlice()
*/
package katas02withencapsulation

// Copyright 2021 JMSilvaDev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package katas19 is a library that contains methods to solve the Kata19 problem, Word Chains, see: http://codekata.com/kata/kata19-word-chains/
/* Usage:
1 - Set the path of the dictionary to be used:
	pathToDictionary, _ := filepath.Abs("../../../storage/wordlist.txt")
2 - Instantiate the wordschain package and call the Discover method with the subject, target and path to dictionary:
	chain, _ := wordschain.Discover("the", "end", pathToDictionary)
*/
package katas19

package main

import (
	"fmt"
)

func main() {
	fmt.Println(find([]rune("could be greater than linear"), []rune(`
In computer science, the Rabin–Karp algorithm or Karp–Rabin algorithm is a
string-searching algorithm created by Richard M. Karp and Michael O. Rabin (1987) that uses
hashing to find an exact match of a pattern string in a text. It uses a rolling hash
to quickly filter out positions of the text that cannot match the pattern, and
then checks for a match at the remaining positions. Generalizations of the same
idea can be used to find more than one match of a single pattern, or to find
matches for more than one pattern.  To find a single match of a single pattern, the
expected time of the algorithm is linear in the combined length of the pattern and
text, although its worst-case time complexity is the product of the two lengths. To
find multiple matches, the expected time is linear in the input lengths, plus the
combined length of all the matches, which could be greater than linear. In contrast,
the Aho–Corasick algorithm can find all matches of multiple patterns in worst-case
time and space linear in the input length and the number of matches (instead of the
total length of the matches).  A practical application of the algorithm is detecting
plagiarism. Given source material, the algorithm can rapidly search through a paper for
instances of sentences from the source material, ignoring details such as case and
punctuation. Because of the abundance of the sought strings, single-string searching
algorithms are impractical.`)))
}

const BASE = int64(262_144)      // 2^18 enough to store the ~150k different unicode points
const MOD = int64((1 << 31) - 1) // 8th mersenne prime, aka math.MaxInt32
// from my understanding the MOD should be a prime number >> BASE but not so large to make BASE * MOD overflow

// rabin-karp algorithm
// for the algorithm we'll need a rolling hash, the basic principle is :
// make a hash of the substring, make a rolling hash of the same length of our text
// for each position in the text, remove the first char from the hash and add the current char
// if the two hashes matches we might have a match
// - we need to have a base >= the != possible characters
// - we need a hash that is also prime
// don't quote me on all that, I'm winging it
// because unicode currently have ~150k char I'll choose the next power of 2 as a limit
func find(substring, text []rune) int {
	subLen := len(substring)
	subHash := hash(substring)
	h := hash(text[:subLen])

	basePower := int64(1)
	for i := 0; i < subLen-1; i++ {
		basePower = (basePower * BASE) % MOD
	}

	for i := subLen; i < len(text); i++ {
		if h == subHash {
			match := true
			// now we need to double-check
			for j := 0; j < subLen; j++ {
				if substring[j] != text[i+j-subLen] {
					match = false
					break
				}
			}
			if match {
				return i - subLen
			}
		}
		cToRemove := (int64(text[i-subLen]) * basePower) % MOD
		h = (h + MOD - cToRemove) % MOD
		h *= BASE
		h = h + int64(text[i])
		h %= MOD
	}

	return -1

}

func hash(s []rune) int64 {
	h := int64(0)
	for _, c := range s {
		h *= BASE
		h += int64(c)
		h %= MOD
	}
	return h
}

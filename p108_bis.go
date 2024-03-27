package main

import "fmt"

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

const BASE = uint32(524_287) // 2^19-1 7th mersenne prime

// rabin-karp algorithm
// same as in p108 but this uses uint32 as a way to do an automatic modulo, overflowing is the same as modulo 2^32
func find(substring, text []rune) int {
	subLen := len(substring)
	subHash := hash(substring)
	h := hash(text[:subLen])

	basePower := uint32(1)
	for i := 0; i < subLen-1; i++ {
		basePower *= BASE
	}

	for i := subLen; i < len(text); i++ {
		// the hash are the same, but we could have a collision
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
		cToRemove := (uint32(text[i-subLen])) * basePower
		h = h - cToRemove
		h *= BASE
		h = h + uint32(text[i])
	}

	return -1

}

func hash(s []rune) uint32 {
	h := uint32(0)
	for _, c := range s {
		h *= BASE
		h += uint32(c)
	}
	return h
}

package solve

import (
	"log"
	"strings"
	"trust-pilot-challenge/hashutils"
)

type Permutation struct {
	permLength, dataSize int
	data                 []string
	phrase               []string
}

func CreatePermutation(permLength int, data []string) Permutation {
	return Permutation{permLength: permLength, data: data,
		dataSize: len(data), phrase: []string{}}
}

func (p Permutation) Solve(targetHash string) {

	ppn := make([]int, p.permLength)
	pp := make([]string, p.permLength)

	for {
		// generate permutaton
		for i, x := range ppn {
			pp[i] = p.data[x]
		}
		// pass to deciding function
		if hashutils.IsTargetHash(strings.Join(pp, " "), targetHash) {
			p.phrase = pp
			log.Println(p.phrase)
			return // terminate early
		}
		// increment permutation number
		for i := 0; ; {
			ppn[i]++
			if ppn[i] < p.dataSize {
				break
			}
			ppn[i] = 0
			i++
			if i == p.permLength {
				return // all permutations generated
			}
		}
	}
}

package main

import (  "fmt"		
)

type suffix struct {
	index int
	rank [2]int	
}

/*

type Suffixes []suffix

func (s Suffixes) Len() int {
	return len(s)
}

func (s Suffixes) Less (i, j int) bool {
	if s[i].rank[0] == s[j].rank[0] {
		if s[i].rank[1] < s[j].rank[1] {
			return true
		} else {
			return false
		}
	} else {
		if s[i].rank[0] < s[j].rank[0] {
			return true
		} else {
			return false
		}
	}
}

func (s Suffixes) Swap (i, j int) {
	s[i], s[j] = s[j], s[i]
}
*/

func countSort (suffixes []suffix, n, exp, rankIndx, startIndx, endIndx int) []suffix {
	
	sortedSuffixLen := endIndx - startIndx
	sortedSuffix := make([]suffix, sortedSuffixLen)
	count := make([]int, 10)
	negBucket := 0
	
	// Fill buckets and count negative numbers
	for i:=startIndx; i<endIndx; i++ {
		bucket := (suffixes[i].rank[rankIndx]/exp)%10
		if bucket < 0 {
			negBucket++
		} else {
			count[bucket]++
		}
	}
	
	// Adjust bucket counts w.r.t negative numbers
	if negBucket > 0 {
		count[0] += negBucket		
	}

	// Compute actual positions
	for i:=1; i<10; i++ {
		count[i] += count[i-1]
	}
	
	// Build sorted array
	for i:=endIndx-1; i>=startIndx; i-- {
		bucket := (suffixes[i].rank[rankIndx]/exp)%10
		if bucket < 0 {
			sortedSuffix[negBucket-1] = suffixes[i]
			negBucket--
		} else {
			sortedSuffix[count[bucket] - 1] = suffixes[i]
			count[bucket]--
		}			
	}
	
	// Adjust original array with sorted output
	for i:=0; i<sortedSuffixLen; i++ {
		suffixes[i+startIndx] = sortedSuffix[i]
	}
	
	return suffixes
}

func getMax (suffixes []suffix, n, rankIndx int) int {
	
	max := suffixes[0].rank[rankIndx]
	
	for i:=1; i<n; i++ {
		if suffixes[i].rank[rankIndx] > max {
			max = suffixes[i].rank[rankIndx]
		}
	}
	return max
}

func radixSort (suffixes []suffix, n int, sortRank1 bool) []suffix {

	max := 0
	
	if sortRank1 {		
		//Get max no. of first rank
		max = getMax(suffixes, n, 0)	
		//Sort first rank	
		for exp:=1; max/exp>0; exp*=10 {
			fmt.Println("exp = ", exp)
			suffixes = countSort(suffixes, n, exp, 0, 0, n)
		}
		fmt.Println("first rank max  = ", max)
		fmt.Println("first rank sort : ")
		fmt.Println(suffixes, "\n")
	}
		
	//Get max no. of second rank
	max = getMax(suffixes, n, 1)
	//Sort second rank	
	fmt.Println("second rank max  = ", max)
	fmt.Println("second rank sort : ")
	for exp:=1; max/exp>0; exp*=10 {
		
		prevRank0 := suffixes[0].rank[0]
		startIndx := 0		
		
		for i:=1; i<n; i++ {			
			if suffixes[i].rank[0] != prevRank0 {
				prevRank0 = suffixes[i].rank[0]
				if startIndx < (i-1) { // to ensure there are atleast 2 entries to be sorted	
					fmt.Println("exp = ", exp, " ; start index = ", startIndx, " ; end index = ", i)				
					suffixes = countSort(suffixes, n, exp, 1, startIndx, i)					
					fmt.Println(suffixes, "\n")	
				}						
				startIndx = i		
			} else if i == (n-1) && startIndx < i { //to ensure last entry is considered
				fmt.Println("exp = ", exp, " ; start index = ", startIndx, " ; end index = ", i+1)
				suffixes = countSort(suffixes, n, exp, 1, startIndx, i+1)
				fmt.Println(suffixes, "\n")
			}
		}		 
	}
	
	return suffixes	
}

func buildSuffixArray (txt string, n int, ) []suffix {
	
	suffixes := make([]suffix, n)
	
	// Build the Ranks according to alphabets in string
	for i:=0; i<n; i++ {
		suffixes[i].index 	= i
		suffixes[i].rank[0] = int(txt[i] - 'a')
		if (i+1) < n {
			suffixes[i].rank[1] = int(txt[i+1] - 'a')
		} else {
			suffixes[i].rank[1] = -1
		}
	}
	
	fmt.Println("Ranks according to alphabets in string")
	fmt.Println(suffixes, "\n")	
	
	fmt.Println("Sort Ranks")
	suffixes = radixSort(suffixes, n, true)
	
	//sort.Sort(Suffixes(suffixes))
	
	for k:=4; k<n ; k*=2 {		
		ind := make([]int, n)
		rank := 0
		prev_rank := suffixes[0].rank[0]
		suffixes[0].rank[0] = rank
		ind[suffixes[0].index] = 0
		
		//Build rank[0] with sequential numbers for unique pair of Ranks
		for i:=1; i<n; i++ {					
			if suffixes[i].rank[0] == prev_rank &&
				suffixes[i].rank[1] == suffixes[i-1].rank[1] {
					prev_rank = suffixes[i].rank[0]
					suffixes[i].rank[0] = rank
			} else {
					rank++
					prev_rank = suffixes[i].rank[0]
					suffixes[i].rank[0] = rank
			}			
			ind[suffixes[i].index] = i
		}
		//Build rank[1]	w.r.t rank[0] of (i + K/2) alphabet
		for i:=0; i<n; i++ {
			nxtindx := suffixes[i].index + (k/2) 
			if nxtindx < n {
				suffixes[i].rank[1] = suffixes[ind[nxtindx]].rank[0]
			} else {
				suffixes[i].rank[1] = -1
			}
		}
		
		fmt.Println("New ranks for K = ", k)
		fmt.Println(suffixes)
		fmt.Println("Sort Ranks")
		suffixes = radixSort(suffixes, n, false)
	}
	
	return suffixes
}

func main () {
	
	txt := "abbcbacba"
	n 	:= len(txt)	
	suffixes := buildSuffixArray( txt, n )	
	fmt.Printf("%v", suffixes)	

}
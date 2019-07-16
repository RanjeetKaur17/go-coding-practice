package strings

import "strconv"

/*
Problem Statement:
The count-and-say sequence is the sequence of integers beginning as follows:
1 is read off as one 1 or 11.
11 is read off as two 1s or 21.

Given an integer n, generate the nth sequence.
*/

func CountAndSay(A int)  (string) {
	//sequence start
	seq := "1"

	for i := 1; i < A; i++ {
		newSeq := ""
		//first element of prev seq and count
		elem := seq[0]
		count := 1
		for i := 1; i < len(seq); i++ {

			if seq[i] == elem {
				//increment count if same value repeats
				count++
			} else {
				//add prev details to new seuence and reset elem and count to current values
				newSeq += strconv.Itoa(count) + string(elem)
				elem = seq[i]
				count = 1
			}
		}
		//add last elem that was counted, to the sequence
		newSeq += strconv.Itoa(count) + string(elem)
		seq = newSeq
	}

	return seq
}
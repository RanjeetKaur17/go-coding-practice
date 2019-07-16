package strings

/*
Problem Statement:
Given an integer, convert it to a roman numeral, and return a string corresponding to its roman numeral version
Input is guaranteed to be within the range from 1 to 3999.

Sample Input:
14

Sample Output:
XIV
 */

func IntToRoman(A int)  (string) {
	s := ""
	for A > 0 {
		if A >= 1000 {
			s += "M"
			A -= 1000
		} else if A >= 900 {
			s += "CM"
			A -= 900
		} else if A >= 500 {
			s += "D"
			A -= 500
		} else if A >= 400 {
			s += "CD"
			A -= 400
		} else if A >= 100 {
			s += "C"
			A -= 100
		} else if A >= 90 {
			s += "XC"
			A -= 90
		} else if A >= 50 {
			s += "L"
			A -= 50
		} else if A >= 40 {
			s += "XL"
			A -= 40
		} else if A >= 10 {
			s += "X"
			A -= 10
		} else if A >= 9 {
			s += "IX"
			A -= 9
		} else if A >= 5 {
			s += "V"
			A -= 5
		} else if A >= 4 {
			s += "IV"
			A -= 4
		} else if A >= 1 {
			s += "I"
			A -= 1
		}
	}

	return s
}

package strings

/*Problem Statement:
Given a roman numeral, convert it to an integer.
Input is guaranteed to be within the range from 1 to 3999.

Sample Input:
XIV

Sample Output:
14
*/

var mapping = map[byte]int {
	'M': 1000,
	'D' : 500,
	'C' : 100,
	'L' : 50,
	'X' : 10,
	'V' : 5,
	'I' : 1,
}

func RomanToInt(A string)  (int) {
	s := 0
	for i := range A {
		mul := 1

		// If a smaller value comes before bigger value,
		// then this value needs to be subtracted from the bigger value to get the addition
		if i < len(A) - 1 && mapping[A[i]] < mapping[A[i+1]] {
			mul = -1
		}
		s+= mul * mapping[A[i]]
	}
	return s
}
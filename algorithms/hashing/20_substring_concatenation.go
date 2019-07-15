package hashing

/*Problem Statement:
You are given a string, S, and a list of words, L, that are all of the same length.
Find all starting indices of substring(s) in S that is a concatenation of each word in L exactly once and without any intervening characters.

Sample Input:
S: "barfoothefoobarman"
L: ["foo", "bar"]

Sample Output:
[0,9]
*/

func FindSubstring(A string , B []string )  []int {

	wm := map[string]int{}
	mA := make([]int, 26)
	mB := make([]int, 26)

	//Find the total number of characters in B
	bTotalLen := 0
	for i := range B {
		for j := range B[i] {
			//Map for characters in list B
			mB[B[i][j] - 'a'] = mB[B[i][j] - 'a'] + 1
			bTotalLen++
		}
		//Map of words in list B
		wm[B[i]]++
	}

	var result []int
	for i := 0; i < len(A); i++ {
		//Add current character from the substring
		mA[A[i] - 'a'] = mA[A[i] - 'a'] + 1

		//Skip till first position where solution can be found
		if i < bTotalLen- 1 {
			continue
		}

		//if substring has same characters as list B
		if equals(mB, mA) {
			//Check if the substring has permutations of words from list B
			if hasAllWords(A[i -bTotalLen+ 1:i+1], copy(wm)) {
				result = append(result, i -bTotalLen+ 1)
			}
		}

		//Remove first character from the substring
		p := A[i -bTotalLen+ 1]
		mA[p - 'a'] = mA[p - 'a'] - 1
	}

	return result
}

//Create copy of the map
func copy(originalMap map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

//Check if current string has all words from the list
func hasAllWords(A string , wm map[string]int ) bool {
	start, end := 0, 0
	for i := range A {
		//check if current substring is a word
		if _, ok := wm[A[start:end+1]]; ok {
			//if yes, update map
			wm[A[start:end+1]]--
			//if count in map goes negative, return false
			if wm[A[start:end+1]] < 0 {
				return false
			} else if wm[A[start:end+1]] == 0 {
				delete(wm, A[start:end+1])
			}

			start = i+1
			end = i
		}
		end++
	}

	//check if all elements found or not
	if len(wm) != 0 {
		return false
	}
	return true
}

//Check if two arrays are equal
func equals(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}

	for i := range A {
		if B[i] != A[i] {
			return false
		}
	}
	return true
}
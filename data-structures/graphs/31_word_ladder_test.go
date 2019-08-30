package graphs

import "testing"

func Test_WordLadderLength(t *testing.T) {
	l := WordLadderLength("hit", "cog", []string{})
	if l != 0 {
		t.Fail()
	}

	l = WordLadderLength("hit", "cog", []string{"hot","dot","dog","lot","log"})
	if l != 5 {
		t.Fail()
	}
}
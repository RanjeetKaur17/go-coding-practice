package graphs

import (
	"testing"
)

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

func Test_WordLadder(t *testing.T) {
	l := WordLadder("hit", "cog", []string{"hot","dot","dog","lot","log"})

	if len(l) != 2 || len(l[0]) != 5 || len(l[1]) != 5 {
		t.Fail()
	}
}
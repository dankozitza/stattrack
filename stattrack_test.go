package stattrack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New("first use of New() in TestNew")

	if (s != StatTrack{1, "INIT", "testing.go:422::stattrack_test.go:9",
		"first use of New() in TestNew", ""}) {

		fmt.Println("TestNew: object template did not match StatTrack object:", s)
		t.Fail()
	}
}

func TestResolve(t *testing.T) {
	sp := New("first use of New() in TestResolve")
	sp.Pass("first use of Pass() in TestResolve")
	sw := New("nth use of New() in TestResolve")
	sw.Warn("first use of Warn() in TestResolve")
	se := New("nth use of New() in TestResolve")
	err := se.Err("first use of Error() in TestResolve")
	fmt.Println("TestResolve: err:", err)
	//sP := New("nth use of New() in TestResolve")
	//sP.Panic("first use of Panic() in TestResolve")
	//spe := New("nth use of New() in TestResolve")
	//spe.PanicErr("first use of Panic() in TestResolve", err)
}

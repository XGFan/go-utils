package utils

import "testing"

func TestPrefixFilter(t *testing.T) {
	var pf PrefixMatcher = make(map[string]interface{})
	t.Run("{A{B{C}}}", func(t *testing.T) {
		pf.AddNode([]string{"A", "B", "C"})
		if pf.MatchNode([]string{"A"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"B"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"A", "B"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "C"}) {
			t.Failed()
		}
	})
	t.Run("{A{B{C,D}}}", func(t *testing.T) {
		pf.AddNode([]string{"A", "B", "D"})
		if pf.MatchNode([]string{"A"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"B"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"A", "B"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "D"}) {
			t.Failed()
		}
	})
	t.Run("{A{B}}", func(t *testing.T) {
		pf.AddNode([]string{"A", "B"})
		if pf.MatchNode([]string{"A"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"B"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"A", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "D"}) {
			t.Failed()
		}
	})
	t.Run("{A}", func(t *testing.T) {
		pf.AddNode([]string{"A"})
		if !pf.MatchNode([]string{"A"}) {
			t.Failed()
		}
		if pf.MatchNode([]string{"X"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "D"}) {
			t.Failed()
		}
	})
	t.Run("{A,X}", func(t *testing.T) {
		pf.AddNode([]string{"X"})
		if !pf.MatchNode([]string{"A"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"X"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"X", "A", "B", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "C"}) {
			t.Failed()
		}
		if !pf.MatchNode([]string{"A", "B", "D"}) {
			t.Failed()
		}
	})
}

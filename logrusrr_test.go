package glogr

import "testing"

func TestInfo(t *testing.T) {
	l, err := New()
	if err != nil {
		t.Error("Error creating glogr")
	}
	l.Error("Brian")
	lf := l.WithField("color", "blue")
	lf.Error("Brian was here")

	lf2 := lf.WithField("name", "bob")
	lf2.Error("Brian was here again")
}

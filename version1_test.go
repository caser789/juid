package uuid

import (
	"bytes"
	"testing"
)

func TestVersion1(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid1 := Must(NewUUID())
		uuid2 := Must(NewUUID())

		if uuid1 == uuid2 {
			t.Errorf("%s:duplicate uuid\n", uuid1)
		}

		if v := uuid1.Version(); v != 1 {
			t.Errorf("%s: version %s expected 1\n", uuid1, v)
		}
		if v := uuid2.Version(); v != 1 {
			t.Errorf("%s: version %s expected 1\n", uuid2, v)
		}

		n1 := uuid1.NodeID()
		n2 := uuid2.NodeID()
		if !bytes.Equal(n1, n2) {
			t.Errorf("Different nodes %x != %x\n", n1, n2)
		}

		t1 := uuid1.Time()
		t2 := uuid2.Time()
		q1 := uuid1.ClockSequence()
		q2 := uuid2.ClockSequence()
		switch {
		case t1 == t2 && q1 == q2:
			t.Errorf("time stopped\n")
		case t1 > t2 && q1 == q2:
			t.Errorf("time reversed\n")
		case t1 < t2 && q1 != q2:
			t.Errorf("clock sequence chaned unexpectedly\n")
		}
	}
}

package _snowflake

import (
	"testing"
	"time"

	"github.com/junyang7/go-common/_exception"
)

func TestNewInitializesBitLayout(t *testing.T) {
	sf := New(1700000000000, 10, 12, 7)

	if sf.nodeShift != 12 {
		t.Fatalf("nodeShift mismatch: got %d, want %d", sf.nodeShift, 12)
	}
	if sf.timeShift != 22 {
		t.Fatalf("timeShift mismatch: got %d, want %d", sf.timeShift, 22)
	}
	if sf.maxNodeID != 1023 {
		t.Fatalf("maxNodeID mismatch: got %d, want %d", sf.maxNodeID, 1023)
	}
	if sf.maxSequence != 4095 {
		t.Fatalf("maxSequence mismatch: got %d, want %d", sf.maxSequence, 4095)
	}
	if sf.nodeId != 7 {
		t.Fatalf("nodeId mismatch: got %d, want %d", sf.nodeId, 7)
	}
}
func TestIdEncodesTimeNodeAndSequence(t *testing.T) {
	epoch := time.Now().Add(-time.Second).UnixMilli()
	sf := New(epoch, 10, 12, 25)

	before := time.Now().UnixMilli()
	id := sf.Id()
	after := time.Now().UnixMilli()

	timePart := id >> sf.timeShift
	nodePart := (id >> sf.nodeShift) & sf.maxNodeID
	sequencePart := id & sf.maxSequence

	if nodePart != sf.nodeId {
		t.Fatalf("node part mismatch: got %d, want %d", nodePart, sf.nodeId)
	}
	if sequencePart != 0 {
		t.Fatalf("first sequence should be 0: got %d", sequencePart)
	}
	if timePart < before-epoch || timePart > after-epoch {
		t.Fatalf("time part out of expected range: got %d, want [%d, %d]", timePart, before-epoch, after-epoch)
	}
}
func TestIdListReturnsOrderedUniqueIDs(t *testing.T) {
	sf := New(time.Now().Add(-time.Second).UnixMilli(), 10, 12, 1)

	ids := sf.IdList(5)
	if len(ids) != 5 {
		t.Fatalf("unexpected id count: got %d, want %d", len(ids), 5)
	}

	seen := make(map[int64]struct{}, len(ids))
	for i, id := range ids {
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate id found at index %d: %d", i, id)
		}
		seen[id] = struct{}{}

		if i > 0 && ids[i-1] >= id {
			t.Fatalf("ids are not strictly increasing: prev=%d current=%d", ids[i-1], id)
		}
	}
}
func TestNewRejectsInvalidNodeID(t *testing.T) {
	exception := catchException(t, func() {
		New(1700000000000, 2, 12, 4)
	})

	if exception.Message != "nodeId 超出可配置范围" {
		t.Fatalf("unexpected message: %s", exception.Message)
	}
}

func TestNewReturnsSameInstanceForSameConfig(t *testing.T) {
	first := New(1700000000000, 10, 12, 3)
	second := New(1700000000000, 10, 12, 3)

	if first != second {
		t.Fatal("expected same instance for identical config")
	}
}

func TestNewReturnsDifferentInstancesForDifferentConfig(t *testing.T) {
	first := New(1700000000000, 10, 12, 3)
	second := New(1700000000001, 10, 12, 3)

	if first == second {
		t.Fatal("expected different instances for different config")
	}
}

func TestIdListRejectsNonPositiveCount(t *testing.T) {
	sf := New(1700000000000, 10, 12, 1)

	for _, count := range []int{0, -1} {
		exception := catchException(t, func() {
			sf.IdList(count)
		})
		if exception.Message != "count必须大于0" {
			t.Fatalf("unexpected message for count=%d: %s", count, exception.Message)
		}
	}
}
func catchException(t *testing.T, fn func()) (got *_exception.Exception) {
	t.Helper()

	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("expected panic, got nil")
		}
		exception, ok := err.(*_exception.Exception)
		if !ok {
			t.Fatalf("unexpected panic type: %T", err)
		}
		got = exception
	}()

	fn()
	return got
}

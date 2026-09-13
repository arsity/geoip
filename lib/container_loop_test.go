package lib

import (
	"fmt"
	"testing"
)

func TestContainerLoopAllowsRemoval(t *testing.T) {
	for _, count := range []int{0, 1, 301, 1024} {
		t.Run(fmt.Sprint(count), func(t *testing.T) {
			c := NewContainer()
			for i := 0; i < count; i++ {
				if err := c.Add(NewEntry(fmt.Sprintf("AS%d", i))); err != nil {
					t.Fatal(err)
				}
			}
			seen := make(map[string]bool)
			for entry := range c.Loop() {
				if seen[entry.GetName()] {
					t.Fatalf("duplicate entry %s", entry.GetName())
				}
				seen[entry.GetName()] = true
				if err := c.Remove(entry, CaseRemoveEntry); err != nil {
					t.Fatal(err)
				}
			}
			if len(seen) != count || c.Len() != 0 {
				t.Fatalf("visited %d/%d entries; remaining=%d", len(seen), count, c.Len())
			}
		})
	}
}

func TestContainerLoopSnapshotsBeforeReturning(t *testing.T) {
	c := NewContainer()
	if err := c.Add(NewEntry("ORIGINAL")); err != nil {
		t.Fatal(err)
	}
	entries := c.Loop()
	if err := c.Add(NewEntry("ADDED-LATER")); err != nil {
		t.Fatal(err)
	}
	count := 0
	for entry := range entries {
		count++
		if entry.GetName() != "ORIGINAL" {
			t.Fatalf("unexpected entry in snapshot: %s", entry.GetName())
		}
	}
	if count != 1 {
		t.Fatalf("snapshot has %d entries, want 1", count)
	}
}

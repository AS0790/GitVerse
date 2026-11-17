package catalog

import "testing"

func TestCatalog_NonEmpty(t *testing.T) {
	if len(Achievements) < 10 {
		t.Fatalf("expected at least 10 achievements, got %d", len(Achievements))
	}
}

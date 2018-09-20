package misavo

import (
	"testing"
	"fmt"
)

func TestGuid(t *testing.T) {
	// Seed()
	id := Guid()
	fmt.Printf("id=%s\n", id)
	if (len(id) != GSIZE*2) {
		t.Error(`id.length != GSIZE*2`)
	}
}


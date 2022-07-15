package base63

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	id := rand.Int63()
	short := Encode(id)
	if len(short) != 10 {
		t.Fatalf("lenght of short url is %d, want 10", len(short))
	}
	for c := range short {
		if !strings.Contains(ALPHABET, strconv.Itoa(c)) {
			t.Fatalf("invalid character in short url: %c", c)
		}
	}
}

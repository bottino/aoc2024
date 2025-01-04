package day22

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSecret(t *testing.T) {
	expSecrets := `15887950
16495136
527345
704524
1553684
12683156
11100544
12249484
7753432
5908254`

	secret := 123
	for i, expStr := range strings.Split(expSecrets, "\n") {
		exp, _ := strconv.Atoi(expStr)
		secret = getSecret(secret)
		if secret != exp {
			t.Errorf("Round %d: exp %d, got %d", i+1, exp, secret)
		}
	}
}

func TestDiff(t *testing.T) {
	expBananas := []int{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}
	expChanges := []int{-3, 6, -1, -1, 0, 2, -2, 0, -2}

	bananas, changes := getDiff(123, 10)

	if diff := cmp.Diff(expBananas, bananas); diff != "" {
		t.Errorf("Bananas mismatch: %s", diff)
	}
	if diff := cmp.Diff(expChanges, changes); diff != "" {
		t.Errorf("Changes mismatch: %s", diff)
	}
}

func TestSequences(t *testing.T) {
	expSeq := map[seq]int{
		{-3, 6, -1, -1}: 4,
		{6, -1, -1, 0}:  4,
		{-1, -1, 0, 2}:  6,
		{-1, 0, 2, -2}:  4,
		{0, 2, -2, 0}:   4,
		{2, -2, 0, -2}:  2,
	}

	seqs := getSequences(123, 10)

	if diff := cmp.Diff(expSeq, seqs); diff != "" {
		t.Errorf("Sequences mismatch: %s", diff)
	}
}

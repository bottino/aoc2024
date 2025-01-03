package day22

import (
	"strconv"
	"strings"
	"testing"
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

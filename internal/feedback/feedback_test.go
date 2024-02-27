package feedback

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExceedsThreshold(t *testing.T) {
	var level, suppress Level

	for level = 0; level <= levelCap; level++ {
		flag := ""
		if level > 0 {
			flag = fmt.Sprintf("-%s ", strings.Repeat("q", int(level)))
		}
		for suppress = 0; suppress <= levelCap; suppress++ {
			SuppressNoise(suppress)
			limit := GetNoiseLimit()
			t.Run(fmt.Sprintf("suppress=%d limit=%d level=%d", suppress, limit, level), func(t *testing.T) {
				expect := level > GetNoiseLimit()
				check := exceedsLimit(level)
				t.Logf("corsair %s (suppress=%d limit=%d level=%d) expect %t outcome %t", flag, suppress, GetNoiseLimit(), level, expect, check)

				assert.Equal(t, expect, check)
			})
		}
	}
}

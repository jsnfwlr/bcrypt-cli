package feedback

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapErrors(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		format    string
		args      []any
		expectErr bool
		expectLen int
		expect    string
	}{
		{
			name:      "err fmt args",
			err:       errors.New("error msg"),
			format:    "%[2]s fmt: %[1]w",
			args:      []any{"arg"},
			expectErr: true,
			expectLen: 102,
			expect:    "github.com.com/corsair/app/corsair/internal/feedback/error_test.go:[0-9]+ - Error: arg fmt: error msg\n",
		},
		{
			name:      "err args",
			err:       errors.New("error msg"),
			args:      []any{"arg"},
			expectErr: true,
			expectLen: 97,
			expect:    "github.com.com/corsair/app/corsair/internal/feedback/error_test.go:[0-9]+ - Error: error msg arg\n",
		},
		{
			name:      "err",
			err:       errors.New("error msg"),
			expectErr: true,
			expectLen: 93,
			expect:    "github.com.com/corsair/app/corsair/internal/feedback/error_test.go:[0-9]+ - Error: error msg\n",
		},
		{
			name:      "err fmt",
			err:       errors.New("error msg"),
			format:    "fmt: %w",
			expectErr: true,
			expectLen: 98,
			expect:    "github.com.com/corsair/app/corsair/internal/feedback/error_test.go:[0-9]+ - Error: fmt: error msg\n",
		},
		{
			name:      "nil fmt",
			err:       nil,
			format:    "fmt: %w",
			expectErr: false,
			expectLen: 0,
			expect:    "",
		},
		{
			name:      "nil fmt args",
			err:       nil,
			format:    "fmt: %w",
			args:      []any{"arg"},
			expectErr: false,
			expectLen: 0,
			expect:    "",
		},
		{
			name:      "nil args",
			err:       nil,
			args:      []any{"arg"},
			expectErr: false,
			expectLen: 0,
			expect:    "",
		},
	}

	buf := &bytes.Buffer{}
	file, err := os.Create("my_file")
	if err != nil {
		t.Fatal(err)
	}

	outputs := []io.Writer{
		// os.Stderr,
		buf,
		file,
	}

	mw := io.MultiWriter(outputs...)
	SetDestination(mw)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleWErr(tt.format, tt.err, tt.args...)
			if !assert.Equal(t, tt.expectErr, tt.err != nil) {
				return
			}
			if !assert.Equal(t, tt.expectLen, len(buf.String())) {
				return
			}
			assert.Regexp(t, tt.expect, buf.String())
		})
		buf.Reset()
	}
}

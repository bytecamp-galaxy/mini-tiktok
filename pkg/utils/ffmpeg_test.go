package utils

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseVideoResolution(t *testing.T) {
	f, err := os.Open("assets/test.mp4")
	assert.NoError(t, err)
	defer f.Close()

	weight, height, err := ParseVideoResolution(context.Background(), f)
	assert.NoError(t, err)
	assert.Equal(t, weight, 320)
	assert.Equal(t, height, 240)
}

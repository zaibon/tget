package betaseries

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	testBs := New(os.Getenv("BS_TOKEN"), "", "")
	show, err := testBs.Show("Vikings")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, 5517, show.ID)
}

func TestEpisodes(t *testing.T) {
	testBs := New(os.Getenv("BS_TOKEN"), "", "")
	episode, err := testBs.Episode("Vikings", 1, 1)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	assert.Equal(t, "Rites of Passage", episode.Title)
}

package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper_getPageQueryParam(t *testing.T) {
	page := getPageQueryParam("10")
	assert.Equal(t, 10, page)

	page = getPageQueryParam("no number")
	assert.Equal(t, 1, page)

	page = getPageQueryParam("")
	assert.Equal(t, 1, page)
}

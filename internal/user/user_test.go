package user

import "testing"
import "github.com/stretchr/testify/assert"

func TestNew(t *testing.T) {
	u := New()

	assert.IsType(t, User{}, u)
	assert.Len(t, u.Id.String(), 36)
}

package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreInPiggyBank(t *testing.T) {

	testBank := piggyBank{0.0, 0, 0, 0, 0}
	storeInPiggyBank(&testBank, 0.01)
	assert.Equal(t, testBank, piggyBank{0.01, 0, 0, 0, 1}, "Balance should be 1 and pennies should be one after deposit")
}

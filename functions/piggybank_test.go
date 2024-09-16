package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSingleCoinDeposit struct {
	amount    float32
	finalBank piggyBank
	coinName  string
}

// TestStoreInPiggyBank tests the storeInPiggyBank function
func TestStoreInPiggyBank(t *testing.T) {

	singleCoinTests := []testSingleCoinDeposit{
		{0.10, piggyBank{0.10, 1, 0, 0, 0}, "Dime"},
		{0.25, piggyBank{0.25, 0, 1, 0, 0}, "Quarter"},
		{0.05, piggyBank{0.05, 0, 0, 1, 0}, "Nickel"},
	}

	for _, test := range singleCoinTests {
		t.Run("Check "+test.coinName+" deposit", func(t *testing.T) {
			testBank := piggyBank{0.0, 0, 0, 0, 0}
			storeInPiggyBank(&testBank, test.amount)
			assert.Equal(t, testBank, test.finalBank,
				fmt.Sprintf("Balance should be %.2f and there should be one %s after deposit", test.amount, test.coinName))
		})
	}

	t.Run("Check multiple coin deposit", func(t *testing.T) {
		testBank := piggyBank{0.0, 0, 0, 0, 0}
		storeInPiggyBank(&testBank, 0.01)
		assert.Equal(t, testBank, piggyBank{0.01, 0, 0, 0, 1}, "Balance should be .01 and there should be one penny after deposit")
		storeInPiggyBank(&testBank, 0.05)
		assert.Equal(t, testBank, piggyBank{0.06, 0, 0, 1, 1}, "Balance should be .06 and there should be one penny and one nickel after deposit")
		storeInPiggyBank(&testBank, 0.25)
		assert.Equal(t, testBank, piggyBank{0.31, 0, 1, 1, 1}, "Balance should be .31 and there should be one penny, one nickel, and one quarter after deposit")
		storeInPiggyBank(&testBank, 0.10)
		assert.Equal(t, testBank, piggyBank{0.41, 1, 1, 1, 1}, "Balance should be .41 and there should be one of each coin after deposit")
	})

}

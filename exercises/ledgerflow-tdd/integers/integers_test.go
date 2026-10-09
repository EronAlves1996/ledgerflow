package integers

import "testing"

func TestAdd(t *testing.T) {
	table := []struct {
		test           string
		number1        int
		number2        int
		expectedResult int
	}{
		{
			test:           "positive case",
			number1:        2,
			number2:        2,
			expectedResult: 4,
		},
		{
			test:           "negative case",
			number1:        3,
			number2:        -4,
			expectedResult: -1,
		},
		{
			test:           "zero case",
			number1:        5,
			number2:        -5,
			expectedResult: 0,
		},
		{
			test:           "another zero case",
			number1:        0,
			number2:        0,
			expectedResult: 0,
		},
	}

	for _, v := range table {
		t.Run(v.test, func(t *testing.T) {
			result := Add(v.number1, v.number2)

			if result != v.expectedResult {
				t.Errorf("want %d, have %d", v.expectedResult, result)
			}
		})
	}

}

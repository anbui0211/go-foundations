package unittest

import "testing"

type AddData struct {
	x, y   int
	result int
}

func TestSum(t *testing.T) {
	testData := []AddData{
		{1, 2, 3},
		{100, 200, 300},
		{69, 96, 165},
	}

	for _, data := range testData {
		result := Sum(data.x, data.y)
		if result != data.result {
			t.Errorf("Sum(%d, %d) FAILED.  Expected %d got %d\n", data.x, data.y, data.result, result)
		}
	}

}

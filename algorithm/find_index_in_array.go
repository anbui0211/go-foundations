package algorithm

import "fmt"

// FindIndexInArray Find the index of the 2nd largest number in the array
func FindIndex() {
	var (
		arr       = []int{5, 2, 7, -4, 2, 7, 0, 6, 1, 3, 7, 6}
		max       = arr[0]
		secondMax = arr[1]
		result    = make([]int, 0)
	)

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		} else if arr[i] > secondMax && arr[i] != max {
			secondMax = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == secondMax {
			result = append(result, i)
		}
	}

	fmt.Println(result)
}

func findIndexOfSecondLargest() int {
	// Parameter
	nums := []int{10, 20, 30, 20, 10}

	if len(nums) < 2 {
		return -1 // Trả về -1 nếu mảng không có đủ phần tử
	}

	firstLargest, secondLargest := 0, -1

	// Tìm chỉ số của số lớn nhất
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[firstLargest] {
			firstLargest = i
		}
	}

	// Tìm chỉ số của số lớn thứ hai
	for i, value := range nums {
		if i != firstLargest { // Bỏ qua số lớn nhất
			if secondLargest == -1 || value > nums[secondLargest] {
				secondLargest = i
			}
		}
	}
	return secondLargest
}

package signal

/*
You are given an array of integers. On each move you are allowed to
increase exactly one of its element by one. Find the minimal number of moves required to obtain a strictly
increasing sequence from the input.

Example

For inputArray = [1, 1, 1], the output should be
solution(inputArray) = 3.

Input/Output

[execution time limit] 4 seconds (go)

[memory limit] 1 GB

[input] array.integer inputArray

Guaranteed constraints:
3 ≤ inputArray.length ≤ 105,
-105 ≤ inputArray[i] ≤ 105.

[output] integer

The minimal number of moves needed to obtain a strictly increasing sequence from inputArray.
It's guaranteed that for the given test cases the answer always fits signed 32-bit integer type.
*/

func ArrayChange(inputArray []int) int {
	sum := 0
	lastNumber := inputArray[0]

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] < lastNumber {
			diff := lastNumber - inputArray[i] + 1
			sum += diff
			inputArray[i] = inputArray[i] + diff
		}

		if inputArray[i] == lastNumber {
			sum += 1
			inputArray[i] = inputArray[i] + 1
		}

		lastNumber = inputArray[i]
	}

	return sum
}

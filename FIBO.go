package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Copied from Mostafa's response to http://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
		987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393,
		196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887,
		9227465, 14930352, 24157817, 39088169}

	nums, err := readFile("input/rosalind_fibo.txt")
	check(err)

	fmt.Println(nums)
	fmt.Println(fibs[nums[0]])
	fmt.Println(string(fibs[nums[0]]))
	fmt.Println(strconv.Itoa(fibs[nums[0]]))
	fmt.Println([]byte(strconv.Itoa(fibs[nums[0]])))

	err = ioutil.WriteFile("output/FIBO.txt",
		[]byte(strconv.Itoa(fibs[nums[0]])),
		0644)
	check(err)
}

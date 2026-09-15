package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func isPalindrome(phrase string) bool {
	if len(phrase) < 2 {
		return false
	}

	specialChars := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	phrase = specialChars.ReplaceAllString(strings.ToLower(phrase), "")
	pointerLeft, pointerRight := 0, len(phrase)-1
	for pointerLeft < pointerRight {
		if phrase[pointerLeft] != phrase[pointerRight] {
			return false
		}
		pointerLeft, pointerRight = pointerLeft+1, pointerRight-1
	}
	return true
}

func reverseArray(input []string) []string {
	for pointerLeft := 0; pointerLeft < len(input)/2; pointerLeft++ {
		pointerRight := len(input) - 1 - pointerLeft
		input[pointerLeft], input[pointerRight] = input[pointerRight], input[pointerLeft]
	}

	return input
}

func pairWithGivenSum(numbers []int, targetSum int) ([2]int, bool) {
	// numbers is sorted
	// if it wasn't, we could do this
	sort.Ints(numbers)
	pointerLeft, pointerRight := 0, len(numbers)-1
	for pointerLeft < pointerRight {
		pairSum := numbers[pointerLeft] + numbers[pointerRight]

		if pairSum == targetSum {
			return [2]int{numbers[pointerLeft], numbers[pointerRight]}, true
		}
		if pairSum > targetSum {
			// increment pointerRight, hold pointerLeft
			pointerRight -= 1
		} else if pairSum < targetSum {
			// decrement pointerRight, hold pointerLeft
			pointerLeft += 1
		}
	}
	return [2]int{-1, -1}, false
}

func trioWithGivenSum(numbers []int, targetSum int) ([3]int, bool) {
	// numbers is sorted
	// if it wasn't, we could do this
	sort.Ints(numbers)

	for i := 0; i < len(numbers)-3; i++ {
		pointerLeft := i + 1
		pointerRight := len(numbers) - 1

		for pointerLeft < pointerRight {
			trioSum := numbers[i] + numbers[pointerLeft] + numbers[pointerRight]
			if trioSum == targetSum {
				return [3]int{numbers[i], numbers[pointerLeft], numbers[pointerRight]}, true
			}
			if trioSum > targetSum {
				// decrement pointerRight, hold pointerLeft
				pointerRight -= 1
			} else if trioSum < targetSum {
				// increment pointerLeft, hold pointerRight
				pointerLeft += 1
			}
		}
	}

	return [3]int{-1, -1, -1}, false
}

func comboWithGivenSum(numbers []int, targetSum int, comboSize int) ([]int, bool) {
	// numbers is sorted
	// if it wasn't, we could do this
	sort.Ints(numbers)

	for i := 0; i < len(numbers)-comboSize; i++ {
		pointerLeft, pointerRight := i+comboSize-2, len(numbers)-1

		for pointerLeft < pointerRight {
			combo := []int{}
			comboSum := 0

			for comboMemberIndex := 0; comboMemberIndex < comboSize-2; comboMemberIndex++ {
				combo = append(combo, numbers[comboMemberIndex])
				comboSum += numbers[comboMemberIndex]
			}

			combo = append(combo, numbers[pointerLeft], numbers[pointerRight])
			comboSum += numbers[pointerLeft] + numbers[pointerRight]

			if comboSum == targetSum {
				return combo, true
			}

			if comboSum > targetSum {
				// decrement pointerRight, hold pointerLeft
				pointerRight -= 1
			} else if comboSum < targetSum {
				// increment pointerLeft, hold pointerRight
				pointerLeft += 1
			}
		}

	}
	return []int{-1}, false
}

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
}

func deleteNthNodeFromLinkedList(head *LinkedListNode, n int) *LinkedListNode {
	pointerLeft, pointerRight := head, head

	for i := 0; i < n; i++ {
		pointerRight = pointerRight.next
	}

	if pointerRight == nil {
		// this means that we are past the tail, and since n elements back
		// is the head, this means that the head is the target to be removed
		// hence a new head is necessary
		return head.next
	}

	// then, we keep moving everyone forwards until
	// pointerRight is the tail, until
	// it has no next to point to
	for pointerRight.next != nil {
		pointerRight = pointerRight.next
		pointerLeft = pointerLeft.next
	}
	// after loop is done, pointerLeft and pointerRight are
	// n elements apart, meaning we should remove pointerLeft.next
	// or re-assign pointerLeft to point to the element
	// after the one it points to
	// Meanwhile, our head was unaffected, since we had enough elems
	// to loop thru w/o touching the head
	pointerLeft.next = pointerLeft.next.next

	return head
}

func reverseSentencePreserveWords(sentence string) string {
	sentence = strings.TrimSpace(sentence)
	doubleSpaces := regexp.MustCompile("\\s+")
	sentence = doubleSpaces.ReplaceAllString(sentence, " ")

	reverseBytes := func(bytes []byte, start, end int) []byte {
		pointerLeft, pointerRight := start, end
		for pointerLeft < pointerRight {
			bytes[pointerLeft], bytes[pointerRight] = bytes[pointerRight], bytes[pointerLeft]
			pointerLeft++
			pointerRight--
		}
		return bytes
	}

	sentenceAsBytesReversed := reverseBytes([]byte(sentence), 0, len(sentence)-1)
	pointerLeft, pointerRight := 0, 0

	for pointerRight < len(sentenceAsBytesReversed) {
		// If we are currently at a space, it means we just passed a full word
		// We revers the full word, because it is out of order after ln 181
		if unicode.IsSpace(rune(sentenceAsBytesReversed[pointerRight])) {
			sentenceAsBytesReversed = reverseBytes(sentenceAsBytesReversed, pointerLeft, pointerRight-1)
			pointerLeft = pointerRight + 1
		}
		pointerRight++
	}

	// reverse the last word
	// added this because otherwise the last word is unreversed
	// and I cannot see why
	sentenceAsBytesReversed = reverseBytes(sentenceAsBytesReversed, pointerLeft, pointerRight-1)
	return string(sentenceAsBytesReversed)
}

func main() {
	phrases := []string{"dennis", "abba", "aba", "Signa te signa temere, me tangis et angis", " aa ", "aabbc cbbaa"}

	phrases = reverseArray(phrases)

	for i := 0; i < len(phrases); i++ {
		var a []any = []any{phrases[i], "isPalindrome?", isPalindrome(phrases[i])}
		fmt.Println(a...)
	}

	// fmt.Println(strings.Repeat("-", 100))

	// numbers := []int{-5, -2, 7, 1, -4, 2, 8, 4, 5}
	// targets := []int{0}

	// for i := 0; i < len(targets); i++ {
	// 	pair, exists := pairWithGivenSum(numbers, targets[i])
	// 	fmt.Println("Array::", numbers, "Target::", targets[i], "Exists?::", exists, "Pair::", pair)

	// 	trio, trioExists := trioWithGivenSum(numbers, targets[i])
	// 	fmt.Println("Array::", numbers, "Target::", targets[i], "Exists?::", trioExists, "Trio::", trio)

	// 	n := 3
	// 	comboN, comboNExists := comboWithGivenSum(numbers, targets[i], n)
	// 	fmt.Println("Array::", numbers, "Target::", targets[i], "Exists?::", comboNExists, "ComboN::", comboN, "ComboSize::", n)

	// 	fmt.Println(strings.Repeat("--", 50))
	// }
	//
	//
	// sentences := []string{"They want a hippopotamus for Christmas", "We love Go "}
	// for i := 0; i < len(sentences); i++ {
	// 	fmt.Println("sentence/reversed::", sentences[i], "/", reverseSentencePreserveWords(sentences[i]))
	// }
}

// Fast amd slow pointers
// Hare-Tortoise algorithm. Floyd's cycle detection alogorithm

func middleNodeOfLinkedList(list []LinkedListNode) LinkedListNode {
	pointerSlow, pointerFast := 0, 0

	for {
		pointerFast += 2
		pointerSlow += 1

		if list[pointerFast].next == nil {
			return list[pointerSlow]
		}
	}
}

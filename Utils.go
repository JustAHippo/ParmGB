package main

import (
	"fmt"
	"strings"
)

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func Len(arr []string) string {
	maxLenInt := len(arr)
	maxLen := fmt.Sprintf("%x", maxLenInt)

	return maxLen
}

//goland:noinspection GoRedundantParens
func ReadOne(arr []string, from string) string {
	var out = ""
	var i = int32(0)

	for {
		end := false

		end = fmt.Sprintf("%x", i) == from

		if end {
			out = arr[i]
			break
		}
		i++
		if i < 0 {
			panic("Couldn't find value at addr: " + from)
		}
	}

	return out
}

//goland:noinspection GoRedundantParens
func Read(arr []string, from string, to string) string {
	var out = ""
	var i = int32(0)

	for {
		end := false

		end = fmt.Sprintf("%x", i) == from

		if end {
			out = arr[i]
			break
		}
		i++
		if i < 0 {
			panic("Couldn't find value at addr (FROM): " + from)
		}
	}

	for {
		end := false

		end = fmt.Sprintf("%x", i) == to

		out += arr[i]
		if end {
			break
		}

		i++
		if i < 0 {
			panic("Couldn't find value at addr (END): " + to)
		}
	}

	return out
}

func normalizePointer(pointer string) string {
	for {
		pointer = pointer[1:]

		if !strings.HasPrefix(pointer, "0") {
			if len(pointer) == 0 {
				pointer = "0"
			}
			break
		}
	}
	return pointer
}

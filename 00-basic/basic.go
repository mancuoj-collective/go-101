package hello

import "strings"

// 01 Constants
const (
	spanish            = "Spanish"
	french             = "French"
	chinese            = "Chinese"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	chineseHelloPrefix = "你好, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// 02 Integers
func Add(a, b int) int {
	return a + b
}

// 03 Iteration
func Repeat1(character string, repeatCount int) (repeated string) {
	for range repeatCount {
		repeated += character
	}
	return
}

// Strings in Go are immutable, meaning every concatenation, such as in our Repeat function,
// involves copying memory to accommodate the new string.
// This impacts performance, particularly during heavy string concatenation.
// go test -bench=.
// go test -bench=. -benchmem
func Repeat2(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func Repeat3(character string, repeatCount int) string {
	repeated := strings.Repeat(character, repeatCount)
	return repeated
}

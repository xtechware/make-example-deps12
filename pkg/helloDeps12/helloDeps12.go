package helloDeps12

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps12.PrintPhrase")
	return phrase
}

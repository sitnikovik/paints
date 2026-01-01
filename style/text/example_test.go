package text_test

import (
	"fmt"

	"github.com/sitnikovik/paints/style/text"
)

// Example shows basic usage of text style functions.
func Example() {
	fmt.Println(text.Bold("Important title"))
	fmt.Println(text.Italic("Emphasis text"))
	fmt.Println(text.Underline("Clickable link"))
	fmt.Println(text.Dim("Less important note"))
}

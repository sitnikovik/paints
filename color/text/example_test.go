package text_test

import (
	"fmt"

	"github.com/sitnikovik/paints/color/text"
)

// Example shows basic usage of text color functions.
func Example() {
	fmt.Println(text.Red("Error message"))
	fmt.Println(text.Green("Success"))
	fmt.Println(text.Blue("Information"))
}

package background_test

import (
	"fmt"

	"github.com/sitnikovik/paints/color/background"
)

// Example shows basic usage of background color functions.
func Example() {
	fmt.Println(background.Red("Error background"))
	fmt.Println(background.Green("Success background"))
	fmt.Println(background.Yellow("Warning background"))
}

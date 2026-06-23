package main

// multiple inmports within go
import (
	"fmt"
	"os"
)

func formatMessage(message string) string {
	return fmt.Sprintf("(fix) %s.", message)
}

// should essentially parse and get args by index
// We would ideally need to know what flags potentially live within the function
// So flags then parse those flags as 1,2,3 and so on.
func main() {

	// var w. inferred type on the os args
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Not enough args present")
		return
	}

	// saving flags as inferred vars
	flags := args[0]
	message := args[1]

	revMessage := fmt.Sprintf("Lookin good %s %s", flags, message)

	fmt.Println(revMessage)

}

// pass_fail reports whether a grade is passing or failing
package pass_fail

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func Grade() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil{
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(status)
}

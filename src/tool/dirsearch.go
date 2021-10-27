package tool

import (
	"fmt"
	"os/exec"
)

func Dirsearch(url string) {
	args := "-u"
	args2 := url
	out, err := exec.Command("dirsearch", args, args2).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out[:])
	fmt.Println("dirsearch output ", output)
}

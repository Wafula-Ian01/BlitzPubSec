package Functionality

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

func getOption(r io.Reader, w io.Writer) (string, error) {
	msg := "Blitz Pub Sec \n Choose one of the options below:\n" +
		" \t 1. Get my device's details\n \t 2. Get another device's details\n" +
		"To select an option enter it's number below to proceed. Thanks\n"
	fmt.Fprint(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	option := scanner.Text()

	if len(option) == 0 {
		return "", errors.New("you cannot proceed without an option selected")
	}

	return "", nil
}

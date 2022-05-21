package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		if input, err := reader.ReadString('\n'); err != nil {
			fmt.Println(err)
		} else {
			if err = execLine(input); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func execLine(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		if err := os.Chdir(args[1]); err != nil {
			return err
		} else {
			return nil
		}
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

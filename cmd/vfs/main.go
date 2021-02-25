package main

import (
	"bufio"
	"fmt"
	"os"
	"virtual-file-system/internal/actions"

	"github.com/google/shlex"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for fmt.Print("# "); scanner.Scan(); fmt.Print("# ") {
		text := scanner.Text()
		fmt.Println("$cmd:", text)

		args, err := shlex.Split(text)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		f := &actions.Factory{}
		if act := f.CreateAction(args); !act.Exec(args) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

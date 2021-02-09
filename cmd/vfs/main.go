package main

import (
	"fmt"

	"github.com/lukefan6/virtual-file-system/internal/pkg/greetings"
	"github.com/lukefan6/virtual-file-system/internal/pkg/logging"
)

func main() {
    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin", ""}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        logging.NewLogger().UnexpectedError(err);
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}
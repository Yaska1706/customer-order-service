package main

import (
	"fmt"
	"os"

	"github.com/yaska1706/customer-order-service/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

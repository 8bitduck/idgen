package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <uuid|ulid> [options]", os.Args[0])
		os.Exit(2)
	}

	identifier := ""
	switch strings.ToLower(os.Args[1]) {
	case "ulid":
		entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		u, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(2)
		}
		identifier = u.String()
	case "uuid":
		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(2)
		}
		identifier = u.String()
	default:
		fmt.Fprintf(os.Stderr, "Invalid identifier type")
		os.Exit(2)
	}

	// Default to Lower Case
	identifier = strings.ToLower(identifier)

	if len(os.Args) > 2 {
		switch strings.ToLower(os.Args[2]) {
		case "--upper":
			identifier = strings.ToUpper(identifier)
		case "--lower":
			//identifier = strings.ToLower(identifier)
			//noop since we already did that
		}
	}

	fmt.Fprintf(os.Stdout, identifier)
}

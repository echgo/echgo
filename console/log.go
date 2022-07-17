package console

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	cyan  = "\u001B[36m"
	red   = "\u001B[31m"
	reset = "\u001B[0m"
)

var color string

// Log is to print the console output
// First check the color / kind & then print the kind, message
// Now check the attribute keys & sort them to print them
func Log(kind, message string, attributes map[string]any) {

	switch {
	case strings.EqualFold("error", kind):
		color = red
	default:
		color = cyan
	}

	kind = fmt.Sprintf("%s%s%s[%d]", color, strings.ToUpper(kind), reset, time.Now().Unix())

	fmt.Printf("%-30s", kind)
	fmt.Printf("%s\t", message)

	keys := make([]string, 0, len(attributes))
	for index := range attributes {
		keys = append(keys, index)
	}
	sort.Strings(keys)

	for _, value := range keys {
		fmt.Printf("\t%s%s%s=%v", color, strings.ToUpper(value), reset, attributes[value])
	}

	fmt.Printf("\n")

}

package lexer_test

import (
  "fmt"
  "strings"
  "testing"
)

func Status(format string, args ...interface{}) {
  if testing.Verbose() {
    fmt.Printf(strings.TrimSpace(format)+"\n", args...)
  }
}

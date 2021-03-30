package main

import "testing"

var tests = []struct{
      s string
      x int
}{
    {"-v", 1},
    {"-version", 3},
    {"-help", 3},
    {"--version", 1},
    {"--help", 0},
    {"-h", 0},

}

func TestChoice(t *testing.T) {
      for _, e := range tests{
        result := choice(e.s)
        if result != e.x{
          t.Errorf("Got %d, expected %d", result, e.x)
        }
      }
}

package main

import "fmt"

// Define a function for each basic gate
func AND(a, b bool) bool {
        return a && b
}

func OR(a, b bool) bool {
        return a || b
}

func NOT(a bool) bool {
        return !a
}

// The circuit: (A OR B) AND (C OR NOT D)
func booleanCircuit(a, b, c, d bool) bool {
        return AND(OR(a, b), OR(c, NOT(d)))
}

func main() {
        // Try all combinations of inputs
        inputs := []struct {
                a, b, c, d bool
        }{
                {false, false, false, false},
                {false, false, false, true},
                {false, false, true, false},
                {false, false, true, true},
                {false, true, false, false},
                {false, true, false, true},
                {false, true, true, false},
                {false, true, true, true},
                {true, false, false, false},
                {true, false, false, true},
                {true, false, true, false},
                {true, false, true, true},
                {true, true, false, false},
                {true, true, false, true},
                {true, true, true, false},
                {true, true, true, true},
        }

        fmt.Println("A\tB\tC\tD\tOutput")
        for _, input := range inputs {
                result := booleanCircuit(input.a, input.b, input.c, input.d)
                fmt.Printf("%v\t%v\t%v\t%v\t%v\n", input.a, input.b, input.c, input.d, result)
        }
}
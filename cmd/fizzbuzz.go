/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func fizzbuzz(max int) {
    for i := 0; i <= max; i++ {
        fizz := i%3 == 0
        buzz := i%5 == 0
        switch {
        case fizz && buzz:
            fmt.Println("fizzbuzz")
        case fizz && !buzz:
            fmt.Println("fizz")
        case !fizz && buzz:
            fmt.Println("buzz")
        default:
            fmt.Println(i)
        }
    }
}

// fizzbuzzCmd represents the fizzbuzz command
var fizzbuzzCmd = &cobra.Command{
	Use:   "fizzbuzz [int]",
	Short: "return Fizzbuzz",
	Long: `return Fizzbuzz
	return Fizzbuzz There is no particular reason because it is a suitable sample`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fizzbuzz called")
		var m int
		m, _ = strconv.Atoi(args[0])
        fizzbuzz(m)
	},
}

func init() {
	rootCmd.AddCommand(fizzbuzzCmd)
}

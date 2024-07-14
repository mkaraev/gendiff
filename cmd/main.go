package main

import (
	"errors"
	"fmt"
	"gendiff/pkg/gendiff"
	"github.com/spf13/cobra"
)

func main() {
	var format string
	cmd := &cobra.Command{
		Use:   "gendiff first_file second_file",
		Short: "Compares two configuration files and shows a difference.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("first_file and second_file are required")
			}
			diff, err := gendiff.GenerateDiff(args[0], args[1], format)
			if err != nil {
				return err
			}
			fmt.Println(diff)
			return nil
		},
	}
	cmd.Flags().StringVarP(&format, "format", "f", "stylish", "set format of output: plain, json, stylish")
	if err := cmd.Execute(); err != nil {
		fmt.Println(fmt.Errorf("error in gendiff: %v", err))
	}
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/westford14/gimmedat/internal/parser"
	"github.com/westford14/gimmedat/internal/times"
)

type movieTimeFlags struct {
	theater string
	zipCode string
}

func init() {
	rootCmd.AddCommand(movieTimesCmd())
}

func configureMovieTimeFlags(flags *movieTimeFlags, cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(
		&flags.theater,
		"theater",
		"t",
		"",
		"give the lowercase name of the movie theater without any unicode accents etc.",
	)
	if err := cmd.MarkFlagRequired("theater"); err != nil {
		panic(err)
	}

	cmd.Flags().StringVarP(
		&flags.zipCode,
		"zipCode",
		"z",
		"",
		"give the lowercase name of the movie theater without any unicode accents etc.",
	)
	if err := cmd.MarkFlagRequired("zipCode"); err != nil {
		panic(err)
	}
}

func movieTimesCmd() *cobra.Command {
	flags := movieTimeFlags{}
	movieTimeCmd := &cobra.Command{
		Use:   "times",
		Short: "Gather movie times from a specified movie theater",
		Long:  "Given a movie theater (assumed in a specific postal code) -- gather movie times",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMovieTimes(cmd, args, flags)
		},
	}

	configureMovieTimeFlags(&flags, movieTimeCmd)
	return movieTimeCmd
}

func runMovieTimes(cmd *cobra.Command, args []string, flags movieTimeFlags) error {
	zipCode, err := parser.ParseZipCodes(flags.zipCode)
	if err != nil {
		return err
	}

	result, err := times.CallAPI(zipCode)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

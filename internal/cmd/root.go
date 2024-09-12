package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const defaultLogLevel = "error"

var (
	logLevels = map[string]slog.Level{
		"":        slog.LevelError, // Default
		"debug":   slog.LevelDebug,
		"info":    slog.LevelInfo,
		"warn":    slog.LevelWarn,
		"warning": slog.LevelWarn,
		"error":   slog.LevelError,
	}
	rootCmd = newRootCmd()
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:          "gimmedat",
		SilenceUsage: true,
		Short:        "CLI for getting movie recommendations and stuff",
	}

	root.PersistentFlags().StringP("verbosity", "v", defaultLogLevel,
		"Level of logging verbosity. One of: debug, info, warn, error. "+
			fmt.Sprintf("Defaults to %s.", defaultLogLevel))
	root.PersistentFlags().Bool("json", false,
		"When set, log entries are emitted in JSON format")

	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		defer slog.DebugContext(cmd.Context(), "Command starting", "command", cmd.CommandPath())

		v, err := cmd.Root().Flags().GetString("verbosity")
		if err != nil {
			return err
		}

		var cmdLoggingAttrs []slog.Attr
		if len(args) > 0 {
			cmdLoggingAttrs = append(cmdLoggingAttrs, slog.String("cli_args", strings.Join(args, ", ")))
		}

		json, err := cmd.Root().Flags().GetBool("json")
		if err != nil {
			return fmt.Errorf("failed to get flag for 'json' logging: %w", err)
		}

		// This validation must be done explicitly here without
		level, valid := logLevels[v]
		if !valid {
			return fmt.Errorf("invalid log level: %v", v)
		}

		hopts := &slog.HandlerOptions{
			Level: level,
		}

		var h slog.Handler
		if json {
			h = slog.NewJSONHandler(os.Stderr, hopts)
		} else {
			h = slog.NewTextHandler(os.Stderr, hopts)
		}

		h = h.WithAttrs(cmdLoggingAttrs)
		slog.SetDefault(slog.New(h))
		return nil
	}

	root.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		slog.DebugContext(cmd.Context(), "Command complete", "command", cmd.CommandPath())

	}

	return root
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

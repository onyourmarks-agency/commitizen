package cz

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/onyourmarks-agency/commitizen/internal/git"
)

func NewInitCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "init",
		Aliases: []string{"install"},
		Short:   "Install this tool to git-core as git-cz.",
		RunE: func(_ *cobra.Command, _ []string) error {
			src, err := exec.LookPath(os.Args[0])
			if err != nil {
				return err
			}
			dst, err := git.InitSubCmd(src, "cz")
			if err != nil {
				return err
			}
			fmt.Printf("Install commitizen to %s\n", dst)
			return nil
		},
	}

	return c
}

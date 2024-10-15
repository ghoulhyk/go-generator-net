package generate

import (
	"github.com/spf13/cobra"
	"log"
)

func Cmd() *cobra.Command {
	var target string
	cmd := &cobra.Command{
		Use:     "generate srcDir [flags]",
		Short:   "生成代码",
		Example: "go-generator-net generate ./templ --target ./remoteServ",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, srcDir []string) {
			err := run(srcDir[0], target)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmd.Flags().StringVar(&target, "target", "", "生成目录")
	return cmd
}

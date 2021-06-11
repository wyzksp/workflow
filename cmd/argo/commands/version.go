package commands

import (
	"os"

	"github.com/argoproj/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/wyzksp/workflow"
	"github.com/wyzksp/workflow/cmd/argo/commands/client"
	infopkg "github.com/wyzksp/workflow/pkg/apiclient/info"
	cmdutil "github.com/wyzksp/workflow/util/cmd"
)

// NewVersionCmd returns a new `version` command to be used as a sub-command to root
func NewVersionCommand() *cobra.Command {
	var short bool
	cmd := cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.PrintVersion(CLIName, argo.GetVersion(), short)
			if _, ok := os.LookupEnv("ARGO_SERVER"); ok {
				ctx, apiClient := client.NewAPIClient()
				serviceClient, err := apiClient.NewInfoServiceClient()
				errors.CheckError(err)
				serverVersion, err := serviceClient.GetVersion(ctx, &infopkg.GetVersionRequest{})
				errors.CheckError(err)
				cmdutil.PrintVersion("argo-server", *serverVersion, short)
			}
		},
	}
	cmd.Flags().BoolVar(&short, "short", false, "print just the version number")
	return &cmd
}

package cmd

import (
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/create"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/destroy"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/run"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/secrets"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/show"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/step"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/upgrade"
	"github.com/jenkins-x-labs/helmboot/pkg/cmd/verify"
	"github.com/jenkins-x-labs/helmboot/pkg/common"
	"github.com/jenkins-x/jx/v2/pkg/log"
	"github.com/spf13/cobra"
)

// Main creates the new command
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   common.TopLevelCommand,
		Short: "boots up Jenkins and/or Jenkins X in a Kubernetes cluster using GitOps",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	cmd.AddCommand(run.NewCmdRun())
	cmd.AddCommand(secrets.NewCmdSecrets())
	cmd.AddCommand(step.NewCmdStep())
	cmd.AddCommand(destroy.NewCmdDestroy())

	cmd.AddCommand(common.SplitCommand(create.NewCmdCreate()))
	cmd.AddCommand(common.SplitCommand(upgrade.NewCmdUpgrade()))
	cmd.AddCommand(verify.NewCmdVerify())
	cmd.AddCommand(common.SplitCommand(show.NewCmdShow()))
	return cmd
}

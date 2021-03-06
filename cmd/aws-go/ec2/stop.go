package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"

	"github.com/bharath-srinivas/aws-go/cmd/aws-go/command"
	"github.com/bharath-srinivas/aws-go/function"
	"github.com/bharath-srinivas/aws-go/internal/spinner"
)

// stop instance command.
var stopCmd = &cobra.Command{
	Use:     "stop [instance id]",
	Short:   "Stop the specified EC2 instance",
	Args:    cobra.ExactArgs(1),
	Example: "aws-go stop i-0a12b345c678de",
	PreRun:  command.PreRun,
	RunE:    stopInstance,
}

func init() {
	ec2Cmd.AddCommand(stopCmd)
	stopCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "perform the operation with dry run enabled")
}

// run command.
func stopInstance(cmd *cobra.Command, args []string) error {
	sp := spinner.Default(spinner.Prefix[2])
	sp.Start()
	sess := ec2.New(command.Session)

	instanceId := function.EC2{
		ID: args[0],
	}

	ec2Service := &function.EC2Service{
		EC2:     instanceId,
		Service: sess,
	}

	resp, err := ec2Service.StopInstance(dryRun)
	if err != nil {
		sp.Stop()
		return err
	}

	previousState := *resp.StoppingInstances[0].PreviousState.Name
	currentState := *resp.StoppingInstances[0].CurrentState.Name
	sp.Stop()
	fmt.Println("Previous State: " + previousState + "\nCurrent State: " + currentState)

	return nil
}

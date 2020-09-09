package nerdgraph

import (
	"os"
	"reflect"
	"strconv"

	"github.com/newrelic/newrelic-cli/internal/client"
	"github.com/newrelic/newrelic-cli/internal/output"
	"github.com/newrelic/newrelic-cli/internal/utils"
	"github.com/newrelic/newrelic-client-go/newrelic"
	"github.com/newrelic/newrelic-client-go/pkg/alerts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	accountID int
)

var cmdMutation = &cobra.Command{
	Use:   "mutation",
	Short: "Entry point for executing mutations on New Relic resources via NerdGraph API mutation endpoints",
	Long: `Entry point for executing mutations on New Relic resources in more detail...

TBD :)
`,
	Example: `newrelic nerdgraph mutation --help`,
	ValidArgs: []string{
		"alertsPolicyCreate",
	},
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.WithClient(func(nrClient *newrelic.NewRelic) {
			mutationsMap := map[string]interface{}{
				"alertsPolicyCreate": nrClient.Alerts.CreatePolicyMutation,
			}

			mutationName := args[0]
			policy := alerts.AlertsPolicyInput{
				IncidentPreference: alerts.AlertsIncidentPreferenceTypes.PER_POLICY,
				Name:               "sblue-cli-crud-test-reflect",
			}

			result, err := invokeMethod(mutationsMap[mutationName], accountID, policy)
			if err != nil {
				log.Fatal(err)
			}

			utils.LogIfFatal(output.Print(result))
		})
	},
}

func invokeMethod(method interface{}, args ...interface{}) (interface{}, error) {
	v := reflect.ValueOf(method)
	rargs := make([]reflect.Value, len(args))
	for i, a := range args {
		rargs[i] = reflect.ValueOf(a)
	}

	result := v.Call(rargs)

	var err error
	if v := result[1].Interface(); v != nil {
		err = v.(error)
	}

	return result[0].Interface(), err
}

func init() {
	Command.AddCommand(cmdMutation)

	cmdMutation.Flags().IntVar(&accountID, "accountId", envAccountID(), "the New Relic account ID to use for operations")
}

func envAccountID() int {
	a, err := strconv.Atoi(os.Getenv("NEW_RELIC_ACCOUNT_ID"))
	if err != nil {
		return 0
	}

	return a
}

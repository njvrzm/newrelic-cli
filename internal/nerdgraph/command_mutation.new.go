package nerdgraph

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/newrelic/newrelic-cli/internal/client"
// 	"github.com/newrelic/newrelic-client-go/newrelic"
// 	"github.com/newrelic/newrelic-client-go/pkg/alerts"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/spf13/cobra"
// )

// var (
// 	accountID     int
// 	mutationInput string
// )

// type alertsPolicyCreateMetadata struct {
// 	Method func(int, alerts.AlertsPolicyInput) (*alerts.AlertsPolicy, error)
// 	Input  alerts.AlertsPolicyInput
// }

// var typeRegistry = make(map[string]interface{})

// var cmdMutation = &cobra.Command{
// 	Use:     "mutation",
// 	Short:   "Entry point for executing mutations on New Relic resources via NerdGraph API mutation endpoints",
// 	Long:    `Entry point for executing mutations on New Relic resources in more detail...TBD :)`,
// 	Example: `newrelic nerdgraph mutation --help`,
// 	ValidArgs: []string{
// 		"alertsPolicyCreate",
// 	},
// 	Args: cobra.OnlyValidArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// Register type in map
// 		typeRegistry["alerts.AlertsPolicyInput"] = alerts.AlertsPolicyInput{}

// 		client.WithClient(func(nrClient *newrelic.NewRelic) {

// 			var input alerts.AlertsPolicyInput
// 			err := json.Unmarshal([]byte(mutationInput), &input)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			fmt.Printf("\n Unmarshal input:  %+v \n\n", input)

// 			// meta := alertsPolicyCreateMetadata{
// 			// 	Method: nrClient.Alerts.CreatePolicyMutation,
// 			// 	Input:  alerts.AlertsPolicyInput{},
// 			// }
// 		})
// 	},
// }

// func init() {
// 	Command.AddCommand(cmdMutation)

// 	cmdMutation.Flags().IntVar(&accountID, "accountId", envAccountID(), "the New Relic account ID to use for operations")
// 	cmdMutation.Flags().StringVar(&mutationInput, "input", "", "the input data to pass to the mutation request")
// }

// func envAccountID() int {
// 	a, err := strconv.Atoi(os.Getenv("NEW_RELIC_ACCOUNT_ID"))
// 	if err != nil {
// 		return 0
// 	}

// 	return a
// }

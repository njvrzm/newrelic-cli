package nerdgraph

import (
  "encoding/json"
  "fmt"
  "os"
  "reflect"
  "strconv"
  "strings"

  "github.com/newrelic/newrelic-cli/internal/client"
  "github.com/newrelic/newrelic-cli/internal/output"
  "github.com/newrelic/newrelic-cli/internal/utils"
  "github.com/newrelic/newrelic-client-go/newrelic"
  "github.com/newrelic/newrelic-client-go/pkg/alerts"
  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
)

// var typeRegistry = make(map[string]reflect.Type)

// func registerType(t interface{}) {
// 	tName := strings.ReplaceAll(reflect.TypeOf(t).String(), "*", "")
// 	typeRegistry[tName] = reflect.TypeOf(t)
// }

// func newInstance(name string) (interface{}, error) {
// 	typ, ok := typeRegistry[name]
// 	if !ok {
// 		return nil, fmt.Errorf("type %s is not registered in type registry", name)
// 	}

// 	return reflect.New(typ).Elem().Interface(), nil
// }

mutationsMap := map[string]endpointMetadata{
  "alertsPolicyCreate": {
    Method: nrClient.Alerts.CreatePolicyMutation,
    Input:  reflect.TypeOf(alerts.AlertsPolicyInput{}),
  },
}

var (
  accountID     int
  mutationInput string
)

type endpointMetadata struct {
  Method    interface{}  `json:"method"`
  Input     reflect.Type `json:"input"`
  InputType string       `json:"inputType"`
}

var cmdMutation = &cobra.Command{
  Use:     "mutation",
  Short:   "Entry point for executing mutations on New Relic resources via NerdGraph API mutation endpoints",
  Long:    `Entry point for executing mutations on New Relic resources in more detail...TBD :)`,
  Example: `newrelic nerdgraph mutation --help`,
  ValidArgs: []string{
    "alertsPolicyCreate",
  },
  Args: cobra.OnlyValidArgs,
  Run: func(cmd *cobra.Command, args []string) {

    client.WithClient(func(nrClient *newrelic.NewRelic) {

      // a := alerts.AlertsPolicyInput{}
      mutationsMap := map[string]endpointMetadata{
        "alertsPolicyCreate": {
          Method: nrClient.Alerts.CreatePolicyMutation,
          Input:  reflect.TypeOf(alerts.AlertsPolicyInput{}),
        },
      }

      mutationName := args[0]

      fmt.Print("\n****************************\n")

      t := mutationsMap[mutationName].Input
      instance := reflect.Zero(t)
      interf := instance.Interface()

      fmt.Printf("\n instance type: %T \n", instance)
      fmt.Printf("\n instance type: %T \n", instance.Convert(t))
      fmt.Printf("\n instance value: %v \n", instance)
      fmt.Printf("\n instance interf: %v \n", interf)

      method := mutationsMap[mutationName].Method
      // input := mutationsMap[mutationName].Input
      err := json.Unmarshal([]byte(mutationInput), &instance)
      if err != nil {
        log.Fatal(err)
      }

      fmt.Printf("\n Unmarshal input:  %+v \n\n", instance)

      result, err := invokeMethod(method, accountID, instance)
      if err != nil {
        log.Fatal(err)
      }

      fmt.Print("\n****************************\n")

      utils.LogIfFatal(output.Print(result))
    })
  },
}

func init() {
  Command.AddCommand(cmdMutation)

  cmdMutation.Flags().IntVar(&accountID, "accountId", envAccountID(), "the New Relic account ID to use for operations")
  cmdMutation.Flags().StringVar(&mutationInput, "input", "", "the input data to pass to the mutation request")
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

func envAccountID() int {
  a, err := strconv.Atoi(os.Getenv("NEW_RELIC_ACCOUNT_ID"))
  if err != nil {
    return 0
  }

  return a
}

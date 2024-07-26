package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/ethpandaops/goomy-blob/scenarios"
	"github.com/ethpandaops/goomy-blob/scenariotypes"
	"github.com/ethpandaops/goomy-blob/tester"
	"github.com/ethpandaops/goomy-blob/utils"
)

type CliArgs struct {
	verbose      bool
	trace        bool
	rpchosts     []string
	rpchostsFile string
	privkey      string
	seed         string
}

func mainArgs(flags *pflag.FlagSet, cliArgs *CliArgs) {

}

func main() {
	cliArgs := CliArgs{}
	flags := pflag.NewFlagSet("main", pflag.ContinueOnError)

	flags.BoolVarP(&cliArgs.verbose, "verbose", "v", false, "Run the script with verbose output")
	flags.BoolVar(&cliArgs.trace, "trace", false, "Run the script with tracing output")
	flags.StringArrayVarP(&cliArgs.rpchosts, "rpchost", "h", []string{}, "The RPC host to send transactions to.")
	flags.StringVar(&cliArgs.rpchostsFile, "rpchost-file", "", "File with a list of RPC hosts to send transactions to.")
	flags.StringVarP(&cliArgs.privkey, "privkey", "p", "", "The private key of the wallet to send funds from.")
	flags.StringVarP(&cliArgs.seed, "seed", "s", "", "The child wallet seed.")

	flags.Parse(os.Args)

	invalidScenario := false
	var scenarioName string
	var scenarioBuilder func() scenariotypes.Scenario
	if flags.NArg() < 2 {
		invalidScenario = true
	} else {
		scenarioName = flags.Args()[1]
		scenarioBuilder = scenarios.Scenarios[scenarioName]
		if scenarioBuilder == nil {
			invalidScenario = true
		}
	}
	if invalidScenario {
		fmt.Printf("invalid or missing scenario\n\n")
		fmt.Printf("implemented scenarios:\n")
		scenarioNames := []string{}
		for sn := range scenarios.Scenarios {
			scenarioNames = append(scenarioNames, sn)
		}
		sort.Slice(scenarioNames, func(a int, b int) bool {
			return strings.Compare(scenarioNames[a], scenarioNames[b]) > 0
		})
		for _, name := range scenarioNames {
			fmt.Printf("  %v\n", name)
		}
		return
	}

	scenario := scenarioBuilder()
	if scenario == nil {
		panic("could not create scenario instance")
	}

	flags.Init(fmt.Sprintf("%v %v", flags.Args()[0], scenarioName), pflag.ExitOnError)
	scenario.Flags(flags)
	cliArgs = CliArgs{}
	flags.Parse(os.Args)

	if cliArgs.trace {
		logrus.SetLevel(logrus.TraceLevel)
	} else if cliArgs.verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	rpcHosts := []string{}
	for _, rpcHost := range strings.Split(strings.Join(cliArgs.rpchosts, ","), ",") {
		if rpcHost != "" {
			rpcHosts = append(rpcHosts, rpcHost)
		}
	}

	if cliArgs.rpchostsFile != "" {
		fileLines, err := utils.ReadFileLinesTrimmed(cliArgs.rpchostsFile)
		if err != nil {
			panic(err)
		}
		rpcHosts = append(rpcHosts, fileLines...)
	}

	testerConfig := &tester.TesterConfig{
		RpcHosts:      rpcHosts,
		WalletPrivkey: cliArgs.privkey,
		WalletCount:   100,
		WalletPrefund: utils.EtherToWei(uint256.NewInt(5)),
		WalletMinfund: utils.EtherToWei(uint256.NewInt(2)),
	}
	err := scenario.Init(testerConfig)
	if err != nil {
		panic(err)
	}

	retry := 0
	for {
		err := runScenario(testerConfig, scenario, &cliArgs)
		if err == nil {
			break
		}

		logrus.Errorf("error running scenario: %v", err)
		retry++
		if retry > 10 {
			break
		}

		time.Sleep(10 * time.Second)
	}
}

func runScenario(testerConfig *tester.TesterConfig, scenario scenariotypes.Scenario, cliArgs *CliArgs) error {
	tester := tester.NewTester(testerConfig)
	err := tester.Start(cliArgs.seed)
	if err != nil {
		return err
	}
	defer tester.Stop()

	return scenario.Run(tester)
}

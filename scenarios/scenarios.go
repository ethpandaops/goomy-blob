package scenarios

import (
	"github.com/ethpandaops/blob-spammer/scenariotypes"

	"github.com/ethpandaops/blob-spammer/scenarios/combined"
	"github.com/ethpandaops/blob-spammer/scenarios/normal"
	"github.com/ethpandaops/blob-spammer/scenarios/replacements"
	"github.com/ethpandaops/blob-spammer/scenarios/wallets"
)

var Scenarios map[string]func() scenariotypes.Scenario = map[string]func() scenariotypes.Scenario{
	"combined":     combined.NewScenario,
	"normal":       normal.NewScenario,
	"replacements": replacements.NewScenario,

	"wallets": wallets.NewScenario,
}

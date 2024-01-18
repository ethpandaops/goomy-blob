package scenarios

import (
	"github.com/ethpandaops/goomy-blob/scenariotypes"

	"github.com/ethpandaops/goomy-blob/scenarios/combined"
	"github.com/ethpandaops/goomy-blob/scenarios/conflicting"
	"github.com/ethpandaops/goomy-blob/scenarios/normal"
	"github.com/ethpandaops/goomy-blob/scenarios/replacements"
	"github.com/ethpandaops/goomy-blob/scenarios/wallets"
)

var Scenarios map[string]func() scenariotypes.Scenario = map[string]func() scenariotypes.Scenario{
	"combined":     combined.NewScenario,
	"conflicting":  conflicting.NewScenario,
	"normal":       normal.NewScenario,
	"replacements": replacements.NewScenario,

	"wallets": wallets.NewScenario,
}

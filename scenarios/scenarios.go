package scenarios

import (
	"github.com/ethpandaops/blob-spammer/scenariotypes"

	"github.com/ethpandaops/blob-spammer/scenarios/combined"
	"github.com/ethpandaops/blob-spammer/scenarios/conflicting"
	"github.com/ethpandaops/blob-spammer/scenarios/deploytx"
	"github.com/ethpandaops/blob-spammer/scenarios/eoatx"
	"github.com/ethpandaops/blob-spammer/scenarios/normal"
	"github.com/ethpandaops/blob-spammer/scenarios/replacements"
	"github.com/ethpandaops/blob-spammer/scenarios/wallets"
)

var Scenarios map[string]func() scenariotypes.Scenario = map[string]func() scenariotypes.Scenario{
	"combined":     combined.NewScenario,
	"conflicting":  conflicting.NewScenario,
	"normal":       normal.NewScenario,
	"replacements": replacements.NewScenario,

	"eoatx":    eoatx.NewScenario,
	"deploytx": deploytx.NewScenario,

	"wallets": wallets.NewScenario,
}

package scenarios

import (
	"github.com/pk910/blob-spammer/scenariotypes"

	"github.com/pk910/blob-spammer/scenarios/combined"
	"github.com/pk910/blob-spammer/scenarios/normal"
	"github.com/pk910/blob-spammer/scenarios/replacements"
	"github.com/pk910/blob-spammer/scenarios/wallets"
)

var Scenarios map[string]func() scenariotypes.Scenario = map[string]func() scenariotypes.Scenario{
	"combined":     combined.NewScenario,
	"normal":       normal.NewScenario,
	"replacements": replacements.NewScenario,

	"wallets": wallets.NewScenario,
}

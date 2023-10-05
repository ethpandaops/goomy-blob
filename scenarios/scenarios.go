package scenarios

import (
	"github.com/pk910/blob-spammer/scenariotypes"

	bloball "github.com/pk910/blob-spammer/scenarios/blob-all"
	blobreplace "github.com/pk910/blob-spammer/scenarios/blob-replace"
	blobspam "github.com/pk910/blob-spammer/scenarios/blob-spam"
	"github.com/pk910/blob-spammer/scenarios/wallets"
)

var Scenarios map[string]func() scenariotypes.Scenario = map[string]func() scenariotypes.Scenario{
	"blob-all":     bloball.NewScenario,
	"blob-spam":    blobspam.NewScenario,
	"blob-replace": blobreplace.NewScenario,

	"wallets": wallets.NewScenario,
}

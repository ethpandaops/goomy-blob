package scenarios

import (
	"github.com/pk910/blob-sender/scenariotypes"

	bloball "github.com/pk910/blob-sender/scenarios/blob-all"
	blobreplace "github.com/pk910/blob-sender/scenarios/blob-replace"
	blobspam "github.com/pk910/blob-sender/scenarios/blob-spam"
	"github.com/pk910/blob-sender/scenarios/wallets"
)

var Scenarios map[string]func() scenariotypes.Scenario = map[string]func() scenariotypes.Scenario{
	"blob-all":     bloball.NewScenario,
	"blob-spam":    blobspam.NewScenario,
	"blob-replace": blobreplace.NewScenario,

	"wallets": wallets.NewScenario,
}

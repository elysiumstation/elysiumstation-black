package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/elysiumstation/blackfury/x/bank/client/cli"
	"github.com/elysiumstation/blackfury/x/bank/client/rest"
)

var (
	SetDenomMetaDataProposalHandler = govclient.NewProposalHandler(cli.NewSetDenomMetaDataProposalCmd, rest.SetDenomMetadataProposalRESTHandler)
)

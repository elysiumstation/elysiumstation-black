package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elysiumstation/blackfury/app"
	blacktypes "github.com/elysiumstation/blackfury/types"
	"github.com/elysiumstation/blackfury/x/ve/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreate_ValidateBasic(t *testing.T) {
	app.Setup(false)
	for _, tc := range []struct {
		desc         string
		sender       string
		to           string
		amount       sdk.Coin
		lockDuration uint64
		valid        bool
	}{
		{
			desc:   "invalid sender address",
			sender: "",
		},
		{
			desc:   "invalid receiver address",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			to:     "xxx",
		},
		{
			desc:   "ErrAmountNotPositive",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			to:     "did:fury:black1353a4uac03etdylz86tyq9ssm3x2704jlshe67",
			amount: sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(0)),
		},
		{
			desc:         "ErrPastLockTime",
			sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			to:           "did:fury:black1353a4uac03etdylz86tyq9ssm3x2704jlshe67",
			amount:       sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
			lockDuration: 0,
		},
		{
			desc:         "ErrTooLongLockTime",
			sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			to:           "did:fury:black1353a4uac03etdylz86tyq9ssm3x2704jlshe67",
			amount:       sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
			lockDuration: types.MaxLockTime + 1,
		},
		{
			desc:         "valid",
			sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			to:           "did:fury:black1353a4uac03etdylz86tyq9ssm3x2704jlshe67",
			amount:       sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
			lockDuration: types.MaxLockTime - 1,
			valid:        true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg := &types.MsgCreate{
				Sender:       tc.sender,
				To:           tc.to,
				Amount:       tc.amount,
				LockDuration: tc.lockDuration,
			}
			err := msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgCreate_GetSigners(t *testing.T) {
	app.Setup(false)
	msg := &types.MsgCreate{
		Sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
		To:           "did:fury:black1353a4uac03etdylz86tyq9ssm3x2704jlshe67",
		Amount:       sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
		LockDuration: uint64(100000),
	}
	signers := msg.GetSigners()
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	require.NoError(t, err)
	require.Equal(t, sender, signers[0])
}

func TestMsgDeposit_ValidateBasic(t *testing.T) {
	app.Setup(false)
	for _, tc := range []struct {
		desc   string
		sender string
		veId   string
		amount sdk.Coin
		valid  bool
	}{
		{
			desc:   "invalid sender address",
			sender: "",
		},
		{
			desc:   "invalid veId",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:   "xxx",
		},
		{
			desc:   "ErrAmountNotPositive",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:   "ve-100",
			amount: sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(0)),
		},
		{
			desc:   "valid",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:   "ve-100",
			amount: sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
			valid:  true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg := &types.MsgDeposit{
				Sender: tc.sender,
				VeId:   tc.veId,
				Amount: tc.amount,
			}
			err := msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgDeposit_GetSigners(t *testing.T) {
	app.Setup(false)
	msg := &types.MsgDeposit{
		Sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
		VeId:   "ve-100",
		Amount: sdk.NewCoin(blacktypes.AttoFuryDenom, sdk.NewInt(1)),
	}
	signers := msg.GetSigners()
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	require.NoError(t, err)
	require.Equal(t, sender, signers[0])
}

func TestMsgExtendTime_ValidateBasic(t *testing.T) {
	app.Setup(false)
	for _, tc := range []struct {
		desc         string
		sender       string
		veId         string
		lockDuration uint64
		valid        bool
	}{
		{
			desc:   "invalid sender address",
			sender: "",
		},
		{
			desc:   "invalid veId",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:   "xxx",
		},
		{
			desc:         "ErrPastLockTime",
			sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:         "ve-100",
			lockDuration: 0,
		},
		{
			desc:         "ErrTooLongLockTime",
			sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:         "ve-100",
			lockDuration: types.MaxLockTime + 1,
		},
		{
			desc:         "valid",
			sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:         "ve-100",
			lockDuration: types.MaxLockTime - 1,
			valid:        true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg := &types.MsgExtendTime{
				Sender:       tc.sender,
				VeId:         tc.veId,
				LockDuration: tc.lockDuration,
			}
			err := msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgExtendTime_GetSigners(t *testing.T) {
	app.Setup(false)
	msg := &types.MsgExtendTime{
		Sender:       "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
		VeId:         "ve-100",
		LockDuration: types.MaxLockTime - 1,
	}
	signers := msg.GetSigners()
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	require.NoError(t, err)
	require.Equal(t, sender, signers[0])
}

func TestMsgMerge_ValidateBasic(t *testing.T) {
	app.Setup(false)
	for _, tc := range []struct {
		desc     string
		sender   string
		fromVeId string
		toVeId   string
		valid    bool
	}{
		{
			desc:   "invalid sender address",
			sender: "",
		},
		{
			desc:     "invalid fromVeId",
			sender:   "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			fromVeId: "xxx",
		},
		{
			desc:     "invalid toVeId",
			sender:   "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			fromVeId: "ve-100",
			toVeId:   "xxx",
		},
		{
			desc:     "fromVeId != toVeId",
			sender:   "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			fromVeId: "ve-100",
			toVeId:   "ve-100",
		},
		{
			desc:     "valid",
			sender:   "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			fromVeId: "ve-100",
			toVeId:   "ve-200",
			valid:    true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg := &types.MsgMerge{
				Sender:   tc.sender,
				FromVeId: tc.fromVeId,
				ToVeId:   tc.toVeId,
			}
			err := msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgMerge_GetSigners(t *testing.T) {
	app.Setup(false)
	msg := &types.MsgMerge{
		Sender:   "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
		FromVeId: "ve-100",
		ToVeId:   "ve-200",
	}
	signers := msg.GetSigners()
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	require.NoError(t, err)
	require.Equal(t, sender, signers[0])
}

func TestMsgWithdraw_ValidateBasic(t *testing.T) {
	app.Setup(false)
	for _, tc := range []struct {
		desc   string
		sender string
		veId   string
		valid  bool
	}{
		{
			desc:   "invalid sender address",
			sender: "",
		},
		{
			desc:   "invalid veId",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:   "xxx",
		},
		{
			desc:   "valid",
			sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
			veId:   "ve-100",
			valid:  true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg := &types.MsgWithdraw{
				Sender: tc.sender,
				VeId:   tc.veId,
			}
			err := msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgWithdraw_GetSigners(t *testing.T) {
	app.Setup(false)
	msg := &types.MsgWithdraw{
		Sender: "did:fury:black1mnfm9c7cdgqnkk66sganp78m0ydmcr4panm2dm",
		VeId:   "ve-100",
	}
	signers := msg.GetSigners()
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	require.NoError(t, err)
	require.Equal(t, sender, signers[0])
}

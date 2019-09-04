package address

import (
	"github.com/mailchain/mailchain/internal/chains"
	"github.com/mailchain/mailchain/internal/encoding"
	"github.com/pkg/errors"
)

func DecodeByProtocol(in, protocol string) ([]byte, error) {
	switch protocol {
	case chains.Ethereum:
		return encoding.DecodeZeroX(in)
	default:
		return nil, errors.Errorf("%q unsupported protocol", protocol)
	}
}
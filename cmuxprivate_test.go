package cmuxprivate

import (
	"testing"

	"github.com/getlantern/cmux/v2"
	"github.com/getlantern/psmux"
)

func TestPsmux(t *testing.T) {
	proto := NewPsmuxProtocol(nil)
	cmux.RunAllProtocolTests(proto, t)

	config := psmux.DefaultConfig()
	config.Version = 2
	proto = NewPsmuxProtocol(config)
	cmux.RunAllProtocolTests(proto, t)
}

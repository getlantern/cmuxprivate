cmuxprivate
===========

Originally used for plugins https://github.com/getlantern/cmux that reference
private repositories (cmux itself is a public repository), but should now just be cleaned up 
and removed.


psmux
-----

example:

```
import (
	"github.com/getlantern/cmux"
	"github.com/getlantern/cmuxprivate"
	"github.com/getlantern/psmux"
)

config := psmux.DefaultConfig()
config.Version = 2
...
protocol := cmuxprivate.NewPsmuxProtocol(config)

dialer := cmux.Dialer(&cmux.DialerOpts{Protocol: protocol, ...}
```

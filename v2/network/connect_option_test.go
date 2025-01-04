package network

import (
	"testing"

	"github.com/sijms/go-ora/v2/configurations"
)

func TestExtractServers(t *testing.T) {
	text := `(DESCRIPTION=
(ADDRESS_LIST=(LOAD_BALANCE=OFF)(FAILOVER=ON)
(ADDRESS=(PROTOCOL=tcp)(HOST=host_dguard)(PORT=1521))
(ADDRESS=(PROTOCOL=tcp)(HOST=host_active)(PORT=1521))
)
(CONNECT_DATA=(SERVICE_NAME=service)(SERVER=DEDICATED))
)`
	t.Log(configurations.ExtractServers(text))
	text = `(DESCRIPTION_LIST=(LOAD_BALANCE=off)(FAILOVER=on)
(DESCRIPTION=(CONNECT_TIMEOUT=5)
(ADDRESS=(PROTOCOL=TCP)(HOST=host_dguard)(PORT=1521))
(CONNECT_DATA=(SERVICE_NAME=service)(SERVER=DEDICATED))
)
(DESCRIPTION=(CONNECT_TIMEOUT=5)
(ADDRESS=(PROTOCOL=TCP)(HOST=host_active)(PORT=1521))
(CONNECT_DATA=(SERVICE_NAME=service)(SERVER=DEDICATED))
)
)`
	t.Log(configurations.ExtractServers(text))
}

func TestUpdateDatabaseInfo(t *testing.T) {
	text := `(DESCRIPTION_LIST=(LOAD_BALANCE=off)(FAILOVER=on)
 (DESCRIPTION=(CONNECT_TIMEOUT=5)(ADDRESS=(PROTOCOL=TCP)
  (HOST=dataguard_host)(PORT=1521))
  (CONNECT_DATA=(SERVICE_NAME=SERVICE_RO)(SERVER=DEDICATED)))
 (DESCRIPTION=(CONNECT_TIMEOUT=5)(ADDRESS=(PROTOCOL=TCP)
  (HOST=active_instance)(PORT=1521))
  (CONNECT_DATA=(SERVICE_NAME=SERVICE)(SERVER=DEDICATED))))`

	text = `DESCRIPTION=(ADDRESS=(PROTOCOL=TCPS)(HOST=host.com)(PORT=1521))(CONNECT_DATA=(SERVER=DEDICATED)(SERVICE_NAME=service))(SECURITY=(SSL_SERVER_CERT_DN="CN=cname,O=org,L=location")))`
	text = `(DESCRIPTION_LIST=(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=host1.domain.com)(PORT=1521))(CONNECT_DATA=(SERVICE_NAME=ServiceName)))(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=host2.domain.com)(PORT=1521))(CONNECT_DATA=(SERVICE_NAME=ServiceName))))`
	op := &configurations.ConnectionConfig{}
	err := op.UpdateDatabaseInfo(text)
	if err != nil {
		t.Error(err)
	}
	t.Log(op)
}

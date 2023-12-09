package route

import (
	"ims-server/device/internal/api"
	ioginx "ims-server/pkg/ginx"
	"net/http"
)

var Routes = []ioginx.Route{
	// TODO：考虑权限问题
	{Func: api.GetDataByID(), FuncName: "GetDataByID", Methods: []string{http.MethodGet, http.MethodPost}, Permission: nil},
}

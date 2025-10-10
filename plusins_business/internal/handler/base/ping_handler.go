package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/logic"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
)

func PingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), ctx)
		result := l.Ping()
		httpx.OkJson(w, result)
	}
}

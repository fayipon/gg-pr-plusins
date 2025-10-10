package handler

import (
	"net/http"

	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/logic"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
	"github.com/fayipon/gg-pr-plusins/plusins_business/desc/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterBusinessHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterBusinessReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterBusinessLogic(r.Context(), ctx)
		resp, err := l.RegisterBusiness(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

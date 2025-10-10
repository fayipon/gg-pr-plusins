package merchant

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
    "github.com/fayipon/gg-pr-plusins/merchant-service/internal/logic/merchant"
    "github.com/fayipon/gg-pr-plusins/merchant-service/internal/svc"
    "github.com/fayipon/gg-pr-plusins/merchant-service/internal/types"
)

func CreateMerchantHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.CreateMerchantReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }

        l := merchant.NewCreateMerchantLogic(r.Context(), ctx)
        if err := l.CreateMerchant(&req); err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, map[string]string{"message": "Merchant created successfully"})
        }
    }
}

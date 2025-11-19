package errorx

import (
    "embed"

    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"
    "github.com/zeromicro/go-zero/core/logx"
)

//go:embed lang/*.yaml
var localeFS embed.FS

var Bundle *i18n.Bundle

func init() {
    Bundle = i18n.NewBundle(language.English)
    Bundle.RegisterUnmarshalFunc("yaml", UnmarshalYAML)

    // Load EN
    if _, err := Bundle.LoadMessageFileFS(localeFS, "lang/en.yaml"); err != nil {
        logx.Error("failed to load en.yaml:", err)
    }

    // Load ZH
    if _, err := Bundle.LoadMessageFileFS(localeFS, "lang/zh.yaml"); err != nil {
        logx.Error("failed to load zh.yaml:", err)
    }
}

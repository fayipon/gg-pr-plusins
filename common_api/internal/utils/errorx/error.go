package errorx

import (
    "context"
    "fmt"
    "time"

    "github.com/nicksnyder/go-i18n/v2/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type CodeError struct {
    Code          int    `json:"code"`
    Message       string `json:"message"`
    LocaleMessage string `json:"locale_message"`
    Timestamp     int64  `json:"timestamp"`
}

func (e *CodeError) Error() string {
    return fmt.Sprintf("code=%d message=%s", e.Code, e.Message)
}

func NewCodeError(ctx context.Context, code int, lang string) *CodeError {

    messageID := fmt.Sprintf("error.%d", code)

    // localized message
    localizer := i18n.NewLocalizer(Bundle, lang)
    localizedMsg, _ := localizer.Localize(&i18n.LocalizeConfig{
        MessageID: messageID,
    })

    // English message
    baseLocalizer := i18n.NewLocalizer(Bundle, "en")
    baseMsg, _ := baseLocalizer.Localize(&i18n.LocalizeConfig{
        MessageID: messageID,
    })

    // Log
    logx.Errorf("[ERROR] code=%d base=%s localized=%s", code, baseMsg, localizedMsg)

    return &CodeError{
        Code:          code,
        Message:       baseMsg,
        LocaleMessage: localizedMsg,
        Timestamp:     time.Now().Unix(),
    }
}

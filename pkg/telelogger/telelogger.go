package telelogger

import (
	"fmt"
	"net/url"
	"skeleton/pkg/rest"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tianrosandhy/goconfigloader"
)

var ERROR_WHITELIST = []string{
	// fill the error whitelist substring that you want to ignore
}

type TeleLogger struct {
	Config *goconfigloader.Config
	Log    *logrus.Logger
}

func NewTeleLogger(logger *logrus.Logger, cfg *goconfigloader.Config) *TeleLogger {
	return &TeleLogger{
		Log:    logger,
		Config: cfg,
	}
}

func (t *TeleLogger) PushError(err error) {
	errData := fmt.Sprintf("SERVICE NAME : *%s* (ENV=%s)", t.Config.GetString("APP_NAME"), t.Config.GetString("ENVIRONMENT"))
	errData = errData + "\n\n" + err.Error()
	t.PushString(errData)
}

func (t *TeleLogger) PushString(stringData string) {
	if !t.Config.GetBool("ENABLE_TELEGRAM_LOG") {
		// telegram log is disabled. no action
		return
	}

	for _, e := range ERROR_WHITELIST {
		if strings.Contains(stringData, e) {
			t.Log.Printf("TeleLogger IGNORED (Whitelist=%s)", e)
			return
		}
	}

	// push to tele logger service
	go func() {
		// url encode stringData
		stringData = url.QueryEscape(stringData)
		teleURLTarget := t.Config.GetString("TELEGRAM_BOT_ENDPOINT") +
			t.Config.GetString("TELEGRAM_BOT_TOKEN") +
			"/sendMessage?"

		teleURL := teleURLTarget +
			"chat_id=" + t.Config.GetString("TELEGRAM_BOT_CHATID") +
			"&text=" + (stringData) +
			"&parse_mode=html"

		cli, err := rest.NewRestClient(t.Config.GetString("TELEGRAM_BOT_LOG_PATH"))
		if err != nil {
			t.Log.Fatal(err)
		}
		cli.SetURL(teleURL).
			SetMethod("POST").
			SetTimeout(5).
			SetDebug(false).
			SetHeaders(map[string]string{
				"Accept": "application/json",
			})

		_, httpCode := cli.Execute()
		if httpCode >= 400 || httpCode < 200 {
			t.Log.Printf("TeleLogger FAIL (%d) : %s", httpCode, teleURLTarget)
		}
	}()

}

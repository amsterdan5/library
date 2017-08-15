package logs

import (
	"testing"

	"github.com/amsterdan/logs"
)

func TestLog(t *testing.T) {
	logs.GetOption().Init()
	logs.Info("111")
	logs.Waring("2222")
}

func TestSetLog(t *testing.T) {
	logOption := logs.GetOption()
	logArr := map[string]interface{}{}
	logArr["filename"] = "log/setTest.log"
	logArr["level"] = logs.L_WARING

	logsJson, err := json.Marshal(logArr)
	if err != nil {
		fmt.Println(err.Error())
	}
	logOption.SetLogger(string(logsJson))
	logOption.Init()

	logs.Info("111")
	logs.Waring("2222")

	str := "hello"
	secondStr := []string{"word", "yes", "goods"}
	logs.Waring(str, secondStr)
}

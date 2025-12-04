package loguru

import "testing"
import "loguru/logger"

func TestAll(t *testing.T) {
	logger.Init(true, "test.log")

	logger.Info("aa", map[string]string{
		"hello": "world",
	}, "xiaohe", []int{1, 2, 3})
	logger.Info("login", "user", "xiaohe", "ip", "1.1.1.1")

}

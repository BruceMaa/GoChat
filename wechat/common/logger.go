package common

import (
	"io"
	"os"
)

var WechatLoggerWriter io.Writer = os.Stdout
var WechatErrorLoggerWriter io.Writer = os.Stderr
package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const RequestBodyMaxSize = 1024 * 1024 * 1024
const LogLevel = hlog.LevelInfo
const CdnServerHost = ":8000"
const ResourceDir = "./res/"
const LogDir = "./log/"

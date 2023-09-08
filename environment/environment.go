package environment

import "strings"

type envVar string

const (
	Prefix = "GR_"
)

const (
	LogLevel    envVar = "LOG_LEVEL"
	LicenseFile envVar = "LICENSE"
)

func (ev envVar) Env() string {
	return Prefix + string(ev)
}

func (ev envVar) Flag() string {
	return strings.ToLower(string(ev))
}

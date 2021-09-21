package internal

import (
	"fmt"
	"runtime/debug"
)

const website = "\n\nhttps://github.com/vandmo/hju"

func Version() string {
	result := version
	result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	result = fmt.Sprintf("%s\nreleased by: %s", result, builder)
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
	}
	return result + website
}

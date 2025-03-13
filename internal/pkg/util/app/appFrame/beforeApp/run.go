package beforeApp

import (
	"gamePractice/internal/pkg/util/app/appFrame/options"
	logCustom "gamePractice/internal/pkg/util/log"
)

func Run(Options *options.Options) {
	logCustom.Init(Options.LogOptions.FilePath)
}

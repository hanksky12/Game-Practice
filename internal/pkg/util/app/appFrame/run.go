package appFrame

import (
	"gamePractice/internal/pkg/util/app/appFrame/afterApp"
	"gamePractice/internal/pkg/util/app/appFrame/beforeApp"
	"gamePractice/internal/pkg/util/app/appFrame/options"
)

func Run(options *options.Options, appFunc func()) {
	beforeApp.Run(options)
	appFunc()
	afterApp.Run(options)
}

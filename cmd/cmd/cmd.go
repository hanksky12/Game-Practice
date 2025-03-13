package main

import (
	"gamePractice/cmd/cmd/app/controller"
	"gamePractice/internal/pkg/util/app/appFrame"
	"gamePractice/internal/pkg/util/app/appFrame/options"
	"gamePractice/internal/pkg/util/app/cmd"
)

func main() {
	options := &options.Options{
		LogOptions: &options.LogOptions{
			FilePath: "log/cmd.log",
		},
	}
	appFrame.Run(options, func() { cmd.Execute(&controller.Job{}) })
	/*
		在主目錄gamePractice下。 (不要直接build cmd.go =>no use)
		go run ./cmd/cmd/cmd.go -c LineGame -p false         (讀取文件)
		go run ./cmd/cmd/cmd.go -c LineGame -p true          (讀取mock)
	*/
}

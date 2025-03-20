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
			Level:    "warn", //warn
		},
	}
	appFrame.Run(options, func() { cmd.Execute(&controller.Job{}) })
	/*
		主目錄下 (不要直接build cmd.go =>no use)
		go run ./cmd/cmd/cmd.go -c Practice1
		go run ./cmd/cmd/cmd.go -c Practice2
		go run ./cmd/cmd/cmd.go -c Practice3  .....
	*/
}

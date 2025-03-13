package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"reflect"
)

var instance any // 存儲外部傳入的 Job 類

// Execute 用於啟動命令解析
func Execute(i any) {
	instance = i // 接收外部傳入的 Job 類
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 定義根命令
var rootCmd = &cobra.Command{
	Use:   "command_app",
	Short: "Command line application for executing job tasks",
	Long:  `This is a command line application for executing various job tasks dynamically.`,
	Run:   run,
}

// 初始化
func init() {
	rootCmd.PersistentFlags().StringP("command", "c", "", "Job command to run (method name)")
	rootCmd.PersistentFlags().StringArrayP("parameter", "p", []string{}, "Parameters for the job method")
}

// run 通用的執行邏輯，動態執行 Job 類中的方法
func run(cmd *cobra.Command, args []string) {
	// 獲取命令和參數
	methodName, _ := cmd.Flags().GetString("command")
	parameters, _ := cmd.Flags().GetStringArray("parameter")

	// 確保命令有效
	if methodName == "" {
		log.Println("Please specify a method name using the -c flag.")
		return
	}

	// 使用反射動態調用外部傳入的 instance 類中的方法
	//jobType := reflect.TypeOf(instance)
	method := reflect.ValueOf(instance).MethodByName(methodName)

	// 如果找不到指定的方法，則返回錯誤
	if !method.IsValid() {
		log.Fatalf("Method %s not found in the provided Job instance\n", methodName)
		return
	}

	// 獲取該方法的參數數量
	numParams := method.Type().NumIn()
	// 如果傳入的參數數量與方法要求的數量相匹配
	if numParams == len(parameters) {
		// 構建反射所需的參數
		inputs := make([]reflect.Value, len(parameters))
		for i, param := range parameters {
			inputs[i] = reflect.ValueOf(param)
		}
		// 調用方法
		method.Call(inputs)
	} else if numParams == 0 {
		// 如果方法不需要參數，直接調用
		method.Call(nil)
	} else {
		log.Fatalf("Method %s expects %d parameters, but got %d\n", methodName, numParams, len(parameters))
	}
}

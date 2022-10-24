package main

import (
	"fmt"
	"os"

	"github.com/nivin-studio/juejin-auto/dingtalk"
	"github.com/nivin-studio/juejin-auto/juejin"
)

func main() {
	message := juejin.New().
		SetCookie(os.Getenv("JUEJIN_COOKIE")).
		CheckIn().
		Lottery().
		DipLucky().
		CollectBug().
		GetResult()

	robot := dingtalk.NewRobot().
		SetWebhook(os.Getenv("DINGTALK_WEBHOOK")).
		SetSecret(os.Getenv("DINGTALK_SECRET"))

	err := robot.SendMessage(dingtalk.NewTextMessage(message))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hello, 世界!")
}

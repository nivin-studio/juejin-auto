package main

import (
	"fmt"
	"juejin-checkin/dingtalk"
	"juejin-checkin/juejin"
	"os"
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

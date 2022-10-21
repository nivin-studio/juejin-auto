package main

import (
	"fmt"
	"juejin-checkin/dingtalk"
	"juejin-checkin/juejin"
	"os"
)

func main() {
	juejin := juejin.New(os.Getenv("JUEJIN_COOKIE"))

	message := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s",
		juejin.CheckIn(),
		juejin.Lottery(),
		juejin.DipLucky(),
		juejin.CollectBug(),
	)

	robot := dingtalk.NewRobot().
		SetWebhook(os.Getenv("DINGTALK_WEBHOOK")).
		SetSecret(os.Getenv("DINGTALK_SECRET"))

	err := robot.SendMessage(dingtalk.NewTextMessage(message))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hello, 世界")
}

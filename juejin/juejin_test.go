package juejin

import (
	"fmt"
	"testing"

	"github.com/nivin-studio/juejin-auto/utils"
)

func TestCheckIn(t *testing.T) {
	juejin := New().SetCookie(utils.Env("JUEJIN_COOKIE", ``))

	result := juejin.CheckIn().GetResult()

	fmt.Print(result)
}

func TestLotteryDraw(t *testing.T) {
	juejin := New().SetCookie(utils.Env("JUEJIN_COOKIE", ``))

	result := juejin.LotteryDraw().GetResult()

	fmt.Print(result)
}

func TestLotteryDip(t *testing.T) {
	juejin := New().SetCookie(utils.Env("JUEJIN_COOKIE", ``))

	result := juejin.LotteryDip().GetResult()

	fmt.Print(result)
}

func TestCollectBugs(t *testing.T) {
	juejin := New().SetCookie(utils.Env("JUEJIN_COOKIE", ``))

	result := juejin.CollectBugs().GetResult()
	fmt.Print(result)
}

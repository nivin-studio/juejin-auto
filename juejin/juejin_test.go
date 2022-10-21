package juejin

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckIn(t *testing.T) {
	juejin := New().SetCookie(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.CheckIn().GetResult()

	fmt.Print(result)
}

func TestLottery(t *testing.T) {
	juejin := New().SetCookie(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.Lottery().GetResult()

	fmt.Print(result)
}

func TestDipLucky(t *testing.T) {
	juejin := New().SetCookie(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.DipLucky().GetResult()

	fmt.Print(result)
}

func TestCollectBug(t *testing.T) {
	juejin := New().SetCookie(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.CollectBug().GetResult()
	fmt.Print(result)
}

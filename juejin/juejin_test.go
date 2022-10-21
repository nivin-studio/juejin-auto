package juejin

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckIn(t *testing.T) {
	juejin := New(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.CheckIn()

	fmt.Print(result)
}

func TestLottery(t *testing.T) {
	juejin := New(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.Lottery()

	fmt.Print(result)
}

func TestDipLucky(t *testing.T) {
	juejin := New(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.DipLucky()

	fmt.Print(result)
}

func TestCollectBug(t *testing.T) {
	juejin := New(os.Getenv("JUEJIN_COOKIE"))

	result := juejin.CollectBug()
	fmt.Print(result)
}

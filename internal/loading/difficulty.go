package loading

import (
	"fmt"
	"strings"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
)

func SelectDif() {
	ClearScreen()

	fmt.Println(ColorCyan + "=================================================" + ColorReset)
	fmt.Println(ColorYellow + "  [ SYS_INIT: ALLOCATING CORE LIMITS (LUA_ENV) ]" + ColorReset)
	fmt.Println(ColorCyan + "=================================================" + ColorReset)

	fmt.Println("[ MEMORY SECTOR ] System integrity severely damaged.")
	fmt.Println("[ DIRECTORY RES ] Select available operational limits:")

	fmt.Printf(" [%s1%s] %sEASY  %s| HDD: 1024KB | LUA_API: Basic   | Clock: 50Hz\n", ColorGreen, ColorReset, ColorGreen, ColorReset)
	fmt.Printf(" [%s2%s] %sNORMAL%s| HDD: 512KB  | LUA_API: Limited | Clock: 25Hz\n", ColorCyan, ColorReset, ColorCyan, ColorReset)
	fmt.Printf(" [%s3%s] %sHARD  %s| HDD: 256KB  | LUA_API: Minimal | Clock: 10Hz\n\n", ColorRed, ColorReset, ColorRed, ColorReset)

	fmt.Print(ColorCyan + "allocation_vector_> " + ColorReset)

	var dif string
	fmt.Scanln(&dif)

	choice := strings.TrimSpace(dif)

	switch choice {
	case "1":
		fmt.Println(ColorGreen + "\n[ OK ] Sector allocated. Environment: EASY." + ColorReset)
	case "2":
		fmt.Println(ColorCyan + "\n[ OK ] Sector allocated. Environment: NORMAL." + ColorReset)
	case "3":
		fmt.Println(ColorRed + "\n[ WARNING ] CRITICAL: Minimal sector allocated (HARD)." + ColorReset)
	default:
		fmt.Println(ColorRed + "\n[ CRITICAL ERROR ] Allocation vector invalid. Rebooting core..." + ColorReset)
		ConsoleBeep(1500 * time.Millisecond)
		time.Sleep(2 * time.Second)
		StartLoading()
		return
	}

	time.Sleep(5 * time.Second)
	ClearScreen()
	fmt.Println("BOOT SEQUENCE COMPLETED. AWAITING LUA SCRIPT...")
}

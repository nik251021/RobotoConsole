package loading

import (
	"fmt"
	"strings"
	"time"

	"RobotoConsole/internal/config"
	"RobotoConsole/internal/other"
)

func SelectDif() {
	ClearScreen()

	fmt.Println(other.ColorCyan + "=================================================" + other.ColorReset)
	fmt.Println(other.ColorYellow + "  [ SYS_INIT: ALLOCATING CORE LIMITS (LUA_ENV) ]" + other.ColorReset)
	fmt.Println(other.ColorCyan + "=================================================" + other.ColorReset)

	fmt.Println("[ MEMORY SECTOR ] System integrity severely damaged.")
	fmt.Println("[ DIRECTORY RES ] Select available operational limits:")

	fmt.Printf(" [%s1%s] %sEASY  %s| HDD: 1024KB | LUA_API: Basic   | Clock: 50Hz\n", other.ColorGreen, other.ColorReset, other.ColorGreen, other.ColorReset)
	fmt.Printf(" [%s2%s] %sNORMAL%s| HDD: 512KB  | LUA_API: Limited | Clock: 25Hz\n", other.ColorCyan, other.ColorReset, other.ColorCyan, other.ColorReset)
	fmt.Printf(" [%s3%s] %sHARD  %s| HDD: 256KB  | LUA_API: Minimal | Clock: 10Hz\n\n", other.ColorRed, other.ColorReset, other.ColorRed, other.ColorReset)

	fmt.Print(other.ColorCyan + "allocation_vector_> " + other.ColorReset)

	var dif string
	fmt.Scanln(&dif)

	choice := strings.TrimSpace(dif)

	switch choice {
	case "1":
		config.SetDifficulty(1)
		fmt.Println(other.ColorGreen + "\n[ OK ] Sector allocated. Environment: EASY." + other.ColorReset)
	case "2":
		config.SetDifficulty(2)
		fmt.Println(other.ColorCyan + "\n[ OK ] Sector allocated. Environment: NORMAL." + other.ColorReset)
	case "3":
		config.SetDifficulty(3)
		fmt.Println(other.ColorRed + "\n[ WARNING ] CRITICAL: Minimal sector allocated (HARD)." + other.ColorReset)
	default:
		fmt.Println(other.ColorRed + "\n[ CRITICAL ERROR ] Allocation vector invalid. Rebooting core..." + other.ColorReset)
		ConsoleBeep(1500 * time.Millisecond)
		time.Sleep(2 * time.Second)

		StartLoading()
		return
	}

	time.Sleep(5 * time.Second)
	ClearScreen()
	fmt.Println("BOOT SEQUENCE COMPLETED. AWAITING LUA SCRIPT...")
}

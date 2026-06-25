package loading

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func ShowBootLog() {
	logLines := []string{
		"[    0.000000] Initializing cgroup subsys cpuset",
		"[    0.012455] CPU0: Intel(R) Core(TM) i7 CPU @ 2.60GHz",
		"[    0.854662] ACPI: Fan [FAN0] (on)",
		"[    2.054124] [  OK  ] Mounted /dev/sda1 to /boot",
		"[    2.412565] [  OK  ] Starting systemd-udevd service...",
		"[    2.854125] [  OK  ] Starting Network Time Synchronization...",
		"[    3.124854] [  OK  ] Enabling swap space...",
		"[    3.541254] [ INFO ] Loading kernel modules...",
		"[    4.214555] [  OK  ] Started RobotoCore OS v4.2.1-RC1.",
		"[    4.912454] READY. Please enter command 'login.'",
	}

	delay := time.Duration(2000/len(logLines)) * time.Millisecond

	for _, line := range logLines {
		fmt.Println(line)
		// В цикле вызываем короткий "щелчок" (80 миллисекунд)
		ConsoleBeep(80 * time.Millisecond)
		time.Sleep(delay)
	}
}

func ConsoleBeep(duration time.Duration) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "rundll32 user32.MessageBeep,0")
	case "darwin":
		cmd = exec.Command("osascript", "-e", "beep 1")
	default:
		cmd = exec.Command("speaker-test", "-t", "pink", "-l", "1")
	}

	if runtime.GOOS == "linux" {
		done := make(chan error)
		go func() {
			done <- cmd.Run()
		}()
		select {
		case <-time.After(duration):
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
		case <-done:
		}
	} else {
		_ = cmd.Run()
	}
}

func StartLoading() {
	ClearScreen()

	ConsoleBeep(750 * time.Millisecond)

	time.Sleep(1 * time.Second)
	ShowBootLog()
	//Sellecting difficulty
	//Initializing game
	var login string
	fmt.Scanln(&login)
	if login != "login" {
		fmt.Println("Access not granded.")
		time.Sleep(1 * time.Second)
		ClearScreen()
		ConsoleBeep(1 * time.Second)
		return
	}
	fmt.Println("Access granted")
	time.Sleep(1 * time.Second)
	ClearScreen()
	ConsoleBeep(1 * time.Second)
	SelectDif()
}

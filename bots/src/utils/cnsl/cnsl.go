package cnsl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func hexToRGB(hex uint32) (uint8, uint8, uint8) {
	return uint8(hex >> 16), uint8(hex >> 8), uint8(hex)
}

func ClearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func osColorFromHex(hex uint32) *string {
	r, g, b := hexToRGB(hex)
	result := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return &result
}

func getTime() *string {
	time := fmt.Sprintf("%02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	return &time
}

func GetUserInput() *string {
	reader := bufio.NewReader(os.Stdin)
	// trim any trailing whitespace
	read, _ := reader.ReadString('\n')
	text := strings.TrimSpace(read)
	return &text
}

func format(color uint32, text *string) string {
	DONE := "\x1b[0m\n"
	time := getTime()
	paint := osColorFromHex(color)
	white := osColorFromHex(0xffffff)
	ret := fmt.Sprintf("%s%s %s#%s %s%s",
		*paint, *time, *white, *paint, *text, DONE)
	return ret

}

func Ok(text string) {
	fmt.Print(format(0x00dead, &text))
}

func Error(text string) {
	fmt.Print(format(0xff1f7a, &text))
}

func Log(text string) {
	fmt.Print(format(0x3d1e6d, &text))
}

func LogLowPrio(text string) {
	fmt.Print(format(0x3b3b3b, &text))
}

func LogHighPrio(text string) {
	fmt.Print(format(0xffabff, &text))
}

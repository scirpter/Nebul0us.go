package tests

import (
	"fmt"
	"io"
	"neb/src/common"
	"os"
	"testing"
)

func TestEncrypt(t *testing.T) {
	file, err := os.Open("../../accounts.txt")
	if err != nil {
		t.Error("Failed to open file.")
	}
	defer file.Close()

	body, _ := io.ReadAll(file)
	str := string(body)
	authData := fmt.Sprintf(str, common.GLOBAL_ACCESS, common.VERSION)
	encrypted, _ := common.Encrypt(&authData)
	fmt.Println(*encrypted)
}

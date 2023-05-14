package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
	"math"
	"neb/src/utils/cnsl"
	"net/http"
	"os"
	"time"
)

type Undefined uint8
type NET_ID uint8

type MapCornerDirection float32

const (
	TOPRIGHT    MapCornerDirection = 0.25 * math.Pi
	TOPLEFT     MapCornerDirection = 0.75 * math.Pi
	BOTTOMLEFT  MapCornerDirection = 1.25 * math.Pi
	BOTTOMRIGHT MapCornerDirection = 1.75 * math.Pi
)

func CornerDirectionStringToFloat(direction *string) MapCornerDirection {
	switch *direction {
	case "tr":
		return TOPRIGHT
	case "tl":
		return TOPLEFT
	case "bl":
		return BOTTOMLEFT
	case "br":
		return BOTTOMRIGHT
	default:
		return BOTTOMLEFT
	}
}

func GetAllCornerDirections() []string {
	return []string{"tr", "tl", "bl", "br"}
}

const (
	ValuedUndefined     Undefined = 0
	ValuedUndefinedHigh Undefined = 0xff
	SERVER_PORT         uint16    = 27900
	GLOBAL_ACCESS                 = 0
	VERSION                       = 107
	ENCRYPTION_KEY                = "YOUR AUTHENTICATION ENCRYPTION KEY"
	KEY_SIZE                      = 32
	MAX_CLIENTS_PER_IP            = 4
)

func GetCommonRequestStamp() uint64 {
	return uint64(time.Now().UnixNano() / 1000000)
}

func GetIP() *string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		cnsl.Error("failed to send ipify request. make sure your internet is working and try again. if it still does not work, use vpn.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	defer resp.Body.Close()

	ip := make([]byte, 16)
	n, err := resp.Body.Read(ip)
	if err != nil {
		cnsl.Error("failed to read ipify response. please contact the bot developer if this keeps happening.")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	ipStr := string(ip[:n])
	return &ipStr
}

func IsDigit(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func Encrypt(plaintext *string) (*string, error) {
	def := ""
	key, _ := hex.DecodeString(ENCRYPTION_KEY)
	keyBytes := sha512.Sum512(key)
	block, err := aes.NewCipher(keyBytes[:KEY_SIZE])
	if err != nil {
		return &def, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(*plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return &def, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(*plaintext))

	res := hex.EncodeToString(ciphertext)
	return &res, nil
}

func Decrypt(ciphertext *string) (*string, error) {
	def := ""
	key, _ := hex.DecodeString(ENCRYPTION_KEY)
	keyBytes := sha512.Sum512(key)
	block, err := aes.NewCipher(keyBytes[:KEY_SIZE])
	if err != nil {
		return &def, err
	}

	ciphertextBytes, err := hex.DecodeString(*ciphertext)
	if err != nil {
		return &def, err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return &def, errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	res := string(ciphertextBytes)
	return &res, nil
}

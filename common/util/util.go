package util

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"
)

func StrToJson(str string) map[string]interface{} {
	var tempMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		// panic(err)
		// logger.Error("StrToJson Error:", err)
		return nil
	}
	return tempMap
}

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func Sha512(data string) string {
	dataSha := sha512.Sum512([]byte(data))
	return hex.EncodeToString(dataSha[:])
}

// 現在時間 timestamp
func TimeNow() int64 {
	// return time.Now().Unix()
	return time.Now().UnixNano() / 1e6
}

func StrToUint16(str string) (uint16, error) {
	num, err := strconv.ParseUint(str, 10, 16)
	return uint16(num), err
}

func StrToUint32(str string) (uint32, error) {
	num, err := strconv.ParseUint(str, 10, 32)
	return uint32(num), err
}

func StrToChar(str string) (byte, error) {
	tmpByte := []byte(str)

	if len(tmpByte) != 1 {
		return byte(0), errors.New("not valid string")
	}

	return tmpByte[0], nil
}

func StrToIntMust(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

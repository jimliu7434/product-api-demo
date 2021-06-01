package log

import (
	"fmt"
	"os"
	"path/filepath"

	"product-api-demo/common/util"
)

func DelLog() {
	diffTime := int64(1000 * 3600 * 24 * keepDay)
	nowTime := util.TimeNow()

	err := filepath.Walk(path, func(file string, f os.FileInfo, err error) error {
		if file == path {
			return nil
		}
		if f == nil {
			return err
		}
		fileTime := f.ModTime().UnixNano() / 1e6
		if (nowTime - fileTime) > diffTime { // 判斷文件是否超過保留天數
			fmt.Printf("Delete file %v !\r\n", file)
			os.RemoveAll(file)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\r\n", err)
	}
}

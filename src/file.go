package src

import (
	"os"
)

// SaveAsMarkdown 检查文件是否存在，如果不存在则创建文件，存在则打开文件并写入内容
// 存储方式不确定，暂时写不了
func SaveAsMarkdown(markdown, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(markdown)
	if err != nil {
		return err
	}

	return nil
}

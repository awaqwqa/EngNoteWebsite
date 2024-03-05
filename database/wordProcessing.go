package database

import (
	"EngNoteWebsite"
	"github.com/gin-gonic/gin"
	"log"
)

/*
该部分代码无法使用，占位置用的
*/

// CreateWord 创建单词
func CreateWord(c *gin.Context) {
	var word EngNoteWebsite.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// 在数据库中插入单词
	_, err := EngNoteWebsite.Db.Exec("INSERT INTO words (text, details, count) VALUES (?, ?, ?)", word.Text, word.Details, word.Count)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Failed to create word"})
		return
	}

	c.JSON(200, gin.H{"message": "Word created successfully"})
}

// SearchWord 搜索单词
func SearchWord(c *gin.Context) {
	text := c.Param("text")

	// 查询单词
	var word EngNoteWebsite.Word
	err := EngNoteWebsite.Db.QueryRow("SELECT * FROM words WHERE text = ?", text).Scan(&word.ID, &word.Text, &word.Details, &word.Count)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": "Word not found"})
		return
	}

	c.JSON(200, word)
}

// IncrementWordCount 增加单词计数
func IncrementWordCount(c *gin.Context) {
	text := c.Param("text")

	// 更新单词计数
	result, err := EngNoteWebsite.Db.Exec("UPDATE words SET count = count + 1 WHERE text = ?", text)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Failed to increment word count"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Failed to increment word count"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Word not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Word count incremented successfully"})
}

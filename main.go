package EngNoteWebsite

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

/*
1.文件读取没做
2.存储方式还没确定
3.所有代码现在都还是一坨，大概看个样子，慢慢提高健壮性
4.数据库代码仅占位，不可用（想好了在说，现在用啥都不知道）

ps：学的go怎么感觉忘得差不多了，又得慢慢学 6666

循环依赖怎么解决啊？？
*/

// Word 结构体用于映射单词表
type Word struct {
	ID      int         `json:"id"`
	Text    string      `json:"text"`
	Details interface{} `json:"meaning"` //该项仍未进行处理
	Count   int         `json:"count"`
}

var Db *sql.DB

func main() {
	var err error
	// 连接数据库
	Db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname") //后续修改
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	// 设置连接池(优化待定
	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)

	// 确保连接正常
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// // 路由设置（待完善
	// router.POST("/words", database.CreateWord)
	// router.GET("/words/:text", database.SearchWord)
	// router.POST("/words/:text/encounter", database.IncrementWordCount)

	// 启动服务器
	router.Run(":8888")
}

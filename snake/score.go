package snake

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Score struct {
	score int     // 分数
	db    *sql.DB // 数据库
}

func NewScore() *Score {
	// 创建 data文件夹 并忽略返回值err
	os.Mkdir("data", os.ModePerm)
	// 打开 SQLite 数据库文件
	db, err := sql.Open("sqlite3", "data\\score.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// 创建数据表
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS scores (
        id INTEGER PRIMARY KEY,
        score INTEGER,
        time INTEGER
    )`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
	return &Score{
		score: 0,
		db:    db,
	}
}

func (s *Score) Save() {
	_, err := s.db.Exec("INSERT INTO scores(score, time) VALUES(?, ?)", s.score, time.Now().Unix())
	if err != nil {
		panic(err)
	}
	log.Println("数据存储成功！", s.score)
}

// HighestScore 查询最高分数
func (s *Score) HighestScore() int {
	// 查询最高分数
	var maxScore int
	err := s.db.QueryRow("SELECT MAX(score) FROM scores").Scan(&maxScore)
	if err != nil {
		maxScore = 0
	}
	return maxScore
}

// Close 关闭数据库
func (s *Score) Close() {
	s.db.Close()
}

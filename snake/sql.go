package snake

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Save struct {
	*sql.DB
}

func NewSave() *Save {
	// 打开 SQLite 数据库文件
	db, err := sql.Open("sqlite3", "assets\\score.db")
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
	return &Save{db}
}

func (s *Save) StoreScore(score int) {
	_, err := s.Exec("INSERT INTO scores(score, time) VALUES(?, ?)", score, time.Now().Unix())
	if err != nil {
		panic(err)
	}
	log.Println("数据存储成功！", score)
}

func (s *Save) HighestScore() int {
	// 查询最高分数
	var maxScore int
	err := s.QueryRow("SELECT MAX(score) FROM scores").Scan(&maxScore)
	if err != nil {
		panic(err)
	}
	return maxScore
}

package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Questions struct {
	Id        int    `db:"id"`
	Question  string `db:"question"`
	CreatedAt string `db:"created_at"`
}

type QAndA struct {
	Id        int     `db:"id"`
	Question  string  `db:"question"`
	Answered  bool    `db:"answered"`
	Answer    *string `db:"answer"`
	CreatedAt string  `db:"created_at"`
}

type Conn struct {
	conn *sqlx.DB
}

func NewDB() (*Conn, error) {
	var err error
	db, err := sqlx.Open(
		"mysql",
		"aratasato:hoge@tcp(127.0.0.1:3306)/geing",
	)

	return &Conn{db}, err
}

// 質問を追加
func (db *Conn) CreateQuestion(body string) error {
	fmt.Println(body)
	tx := db.conn.MustBegin()
	tx.MustExec("INSERT INTO qandas (question) VALUES (?)", body)
	err := tx.Commit()
	return errors.WithMessage(err, "failed: add question")
}

// 質問回答セットを1件取得
func (db *Conn) GetQA(id int) (QAndA, error) {
	qa := QAndA{}
	err := db.conn.Get(&qa, "SELECT * FROM qandas WHERE id = ?", id)
	return qa, errors.WithMessage(err, "failed: get qa")
}

func (db *Conn) GetQuestions(page int) ([]Questions, error) {
	var questions []Questions
	err := db.conn.Select(&questions, "SELECT id, question, created_at FROM qandas WHERE id > ? * 10 LIMIT 20", page)
	return questions, errors.WithMessage(err, "failed: get question")
}
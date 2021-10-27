package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

type SqlLite struct {
	db *sql.DB
}

func GetSqlLite() SqlLite {
	var sqlLiteConnect SqlLite
	return sqlLiteConnect
}

func (s SqlLite) Insert(table, fields, values string) (int64, error) {
	s.connect()
	defer s.close()

	result, err := s.db.Exec(
		fmt.Sprintf(
			"insert into %s (%s) values (%s)",
			table,
			fields,
			values,
		),
	)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.LastInsertId()
}

func (s SqlLite) Update(table, id string, values map[string]string) (int64, error) {
	s.connect()
	defer s.close()

	var set []string
	for field, val := range values {
		set = append(set, fmt.Sprintf("%s = '%s'", field, val))
	}

	result, err := s.db.Exec(
		fmt.Sprintf(
			"update %s set %s where id = %s",
			table,
			strings.Join(set, ", "),
			id,
		),
	)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.RowsAffected()
}

func (s SqlLite) SelectRow(table, where string) *sql.Row {
	s.connect()
	defer s.close()

	return s.db.QueryRow(fmt.Sprintf("select * from %s where %s", table, where))
}

func (s SqlLite) DeleteById(table, id string) (int64, error) {
	s.connect()
	defer s.close()

	result, err := s.db.Exec(fmt.Sprintf("delete from %s where id = $1", table), id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.RowsAffected()
}

func (s SqlLite) connect() {
	if s.db != nil {
		return
	}

	var err error
	s.db, err = sql.Open("sqlite3", "db.db")

	if err != nil {
		log.Fatalln(err)
	}
}

func (s SqlLite) FindInSet(table, field, value string) (*sql.Rows, error) {
	s.connect()
	defer s.close()

	rows, err := s.db.Query(
		fmt.Sprintf(
			"select * from %s where %s like '%%,%s,%%' or %s like '%s,%%' or %s like '%%,%s' or %s = '%s'",
			table,
			field,
			value,
			field,
			value,
			field,
			value,
			field,
			value,
		),
	)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return rows, nil
}

func (s SqlLite) close() {
	if s.db == nil {
		return
	}
	if err := s.db.Close(); err != nil {
		log.Fatalln(err)
	}
}

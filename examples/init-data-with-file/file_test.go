package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hedon954/go-mysql-mocker/gmm"
)

func Test_InitDataWithFile(t *testing.T) {
	db, _, shutdown, err := gmm.Builder("test").
		SQLFiles("../../fixtures/sequel_ace.sql").
		Build()
	if err != nil {
		panic(err)
	}
	defer shutdown()

	rows, err := db.Query("select count(*) from config_info")
	assert.Nil(t, err)
	var count int
	if rows.Next() {
		assert.Nil(t, rows.Scan(&count))
	}
	assert.Equal(t, 45, count)
}

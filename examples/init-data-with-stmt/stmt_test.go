package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hedon954/gmm"
	"github.com/hedon954/gmm/examples"
)

func Test_InitDataWithStmt(t *testing.T) {
	_, db, shutdown, err := gmm.Builder(examples.DBName).
		CreateTable(examples.CertificationInfo{}).
		SQLStmts("insert into certification_info(id, username, password) values (1, 'hedon', 'hedon-pwd');").
		Build()
	assert.Nil(t, err)
	defer shutdown()

	var res []examples.CertificationInfo
	err = db.Select("id", "username", "password").Find(&res).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "hedon", res[0].Username)
}

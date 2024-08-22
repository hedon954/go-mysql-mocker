package main

import (
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/hedon954/go-mysql-mocker/examples"
	"github.com/hedon954/go-mysql-mocker/gmm"
)

func Test_InitDatWithModel(t *testing.T) {
	var data []*examples.CertificationInfo
	for i := 0; i < 5; i++ {
		data = append(data, &examples.CertificationInfo{ID: 10 + i, Username: "root", Password: uuid.NewString()})
	}

	_, db, shutdown, err := gmm.Builder(examples.DBName).InitData(data).Build()
	assert.Nil(t, err)
	defer shutdown()

	var result []examples.CertificationInfo
	err = db.Select("id", "username", "password").Find(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(result))
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	assert.Equal(t, 10, result[0].ID)
}

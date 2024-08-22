package hello_world

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hedon954/gmm"
)

func TestChangeUserStateToMatch(t *testing.T) {
	t.Run("no data should have no affect row", func(t *testing.T) {
		db, _, shutdown, err := gmm.Builder("db-name").Port(20201).CreateTable(UserState{}).Build()
		assert.Nil(t, err)
		defer shutdown()
		res, err := ChangeUserStateToMatch(db, "1")
		assert.Nil(t, err)
		assert.Equal(t, int64(0), res)
	})

	t.Run("has uid 1 should affect 1 row and change state to `match`", func(t *testing.T) {
		// prepare db and init data
		origin := UserState{State: "no-match", UID: "1"}
		db, gDB, shutdown, err := gmm.Builder().CreateTable(UserState{}).InitData(&origin).Build()
		assert.Nil(t, err)
		defer shutdown()

		// check before change state
		var before = UserState{}
		assert.Nil(t, gDB.Select("uid", "state").Where("uid=?", "1").Find(&before).Error)
		assert.Equal(t, origin, before)

		// run biz logic
		res, err := ChangeUserStateToMatch(db, "1")
		assert.Nil(t, err)
		assert.Equal(t, int64(1), res)

		// check after change state
		var after = UserState{}
		assert.Nil(t, gDB.Select("uid", "state").Where("uid=?", "1").Find(&after).Error)
		assert.Equal(t, UserState{
			UID:   "1",
			State: "match",
		}, after)
	})
}

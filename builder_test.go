package pgxsqlbuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBuilder(t *testing.T) {
	cmd := NewUpdateBuilder(`public."user"`,
		WithSetFields(
			NewField("name", "Alexander"),
			NewField("surname", "Teplov"),
			NewField("sex", "male"),
		),
		WithAndWhereFields(NewField("id=", 1)))
	query, err := cmd.Build()
	assert.NoError(t, err)
	assert.Equal(t, `UPDATE public."user" SET name = $1, surname = $2, sex = $3 WHERE id=$4`, query.SQL)
	assert.Equal(t, []interface{}{"Alexander", "Teplov", "male", 1}, query.Values)
}

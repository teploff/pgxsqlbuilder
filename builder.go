package pgxsqlbuilder

import (
	"strconv"
	"strings"
)

type Option func(builder *UpdateBuilder)

type UpdateBuilder struct {
	tableName    string
	setFields    []*Field
	andWhereExpr []*Field
	orWhereExpr  []*Field
}

func NewUpdateBuilder(tableName string, opts ...Option) *UpdateBuilder {
	builder := &UpdateBuilder{
		tableName:    tableName,
		setFields:    make([]*Field, 0, 1),
		andWhereExpr: make([]*Field, 0, 1),
		orWhereExpr:  make([]*Field, 0, 1),
	}

	for _, opt := range opts {
		opt(builder)
	}

	return builder
}

func (u *UpdateBuilder) Build() (*Query, error) {
	setArea := ""
	whereArea := ""
	count := 0
	values := make([]interface{}, 0, 1)

	for _, field := range u.setFields {
		count += 1
		setArea += field.name + " = $" + strconv.Itoa(count) + ", "
		values = append(values, field.value)
	}
	setArea = strings.TrimRight(setArea, ", ")

	for _, expr := range u.andWhereExpr {
		count += 1
		whereArea += "AND " + expr.name + "$" + strconv.Itoa(count)
		values = append(values, expr.value)
	}
	whereArea = strings.TrimLeft(whereArea, "AND ")

	for _, expr := range u.orWhereExpr {
		count += 1
		whereArea += "OR " + expr.name + "$" + strconv.Itoa(count)
		values = append(values, expr.value)
	}

	return &Query{
		sql:    "UPDATE " + u.tableName + " SET " + setArea + " WHERE " + whereArea,
		values: values,
	}, nil
}

func WithSetFields(fields ...*Field) Option {
	return func(builder *UpdateBuilder) {
		builder.setFields = append(builder.setFields, fields...)
	}
}

func WithAndWhereFields(fields ...*Field) Option {
	return func(builder *UpdateBuilder) {
		builder.andWhereExpr = append(builder.andWhereExpr, fields...)
	}
}

func WithOrWhereFields(fields ...*Field) Option {
	return func(builder *UpdateBuilder) {
		builder.orWhereExpr = append(builder.orWhereExpr, fields...)
	}
}

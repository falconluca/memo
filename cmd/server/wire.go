//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/shaohsiung/memo/api/protobuf"
	"github.com/shaohsiung/memo/internal/memo"
	"github.com/shaohsiung/memo/internal/pkg/dbcontext"
)

func InitMemoServer(DSN string) (protobuf.MemoServer, error) {
	wire.Build(memo.NewServer, memo.NewRepository, dbcontext.NewDB)
	return memo.Server{}, nil
}

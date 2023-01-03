package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"mt-api-go/drivers"
)

func main() {

	// db-migrationの実行
	// コマンドラインの場合
	//  migrate -database 'postgres://mtadmin:mtadmin@localhost:5433/muscletracking?sslmode=disable' -path tools/db/migrations up 実行
	//  migrate -database 'postgres://mtadmin:mtadmin@localhost:5433/muscletracking?sslmode=disable' -path tools/db/migrations down 取消
	m, err := migrate.New(
		"file://tools/db/migrations",
		"postgres://mtadmin:mtadmin@localhost:5433/muscletracking?sslmode=disable")

	if err != nil {
		fmt.Println("err1")
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			fmt.Println("err2")
			panic(err)
		}
	}

	// SQLデバッグ
	boil.DebugMode = true

	engine := drivers.InitRouter()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	err = engine.Run(":8081")
	if err != nil {
		fmt.Println("ERROR!!!")
	}
}

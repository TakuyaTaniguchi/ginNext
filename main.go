package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

// "github.com/joho/godotenv"

// env := os.Getenv("ENV")

// if "production" == env {
// 	env = "production"
// } else {
// 	env = "development"
// }

// // 環境変数取得
// godotenv.Load(".env." + env)
// godotenv.Load()

func main() {
	engine := gin.Default()

	dbInit()

	//Index
	engine.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()

		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"todos":   todos,
		})
	})

	//Create
	engine.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	engine.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "details",
			"todo":    todo,
		})
	})

	//Update
	engine.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	//削除確認
	engine.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dbGetOne(id)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "DelteCheck",
			"todo":    todo,
		})
	})

	//Delete
	engine.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbDelete(id)
		ctx.Redirect(302, "/")

	})

	engine.Run(":4000")

}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "todos"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

//DB初期化
func dbInit() {

	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//DB追加
func dbInsert(text string, status string) {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func dbDelete(id int) {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic("データベース開けず！（dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func dbGetAll() []Todo {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

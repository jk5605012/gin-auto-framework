// +build windows
package create

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateFolder(name string) {
	base := fmt.Sprintf(".\\%s", name)
	if err := os.MkdirAll(base, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	if err := createSubfolder(base, "routers"); err != nil {
		log.Fatalln(err)
	}
	if err := createSubfolder(base, "db"); err != nil {
		log.Fatal(err)
	}
	if err := createSubfolder(base, "conf"); err != nil {
		log.Fatal(err)
	}
	pkg := fmt.Sprintf("%s\\%s", base, "pkg")
	if err := createSubfolder(pkg, "e"); err != nil {
		log.Fatalln(err)
	}
	if err := createSubfolder(pkg, "setting"); err != nil {
		log.Fatal(err)
	}
	app := fmt.Sprintf("%s\\%s", base, "app")
	if err := createSubfolder(app, "controllers"); err != nil {
		log.Fatalln(err)
	}
	if err := createSubfolder(app, "models"); err != nil {
		log.Fatalln(err)
	}
	if err := createSubfolder(app, "middlewares"); err != nil {
		log.Fatalln(err)
	}
}

func CreateAllFile(name string, dbType string) {
	base := fmt.Sprintf(".\\%s", name)

	db := dbFormat(dbType)
	f, err := createFile(base, "\\db\\db.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(db); err != nil {
		log.Fatal(err)
	}
}

func createFile(path, fileName string) (*os.File, error) {
	filePath := fmt.Sprintf("%s\\%s", path, fileName)
	f, err := os.Create(filePath)
	return f, err
}

func createSubfolder(path string, folderName string) error {
	folder := fmt.Sprintf("%s\\%s", path, folderName)
	err := os.MkdirAll(folder, os.ModePerm)
	return err
}

func dbFormat(dbType string) (result string) {
	db := `
package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/{type}"
)

// test
var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("{type}", {conn})
	if err != nil {
		log.Fatal(err)
		return
	}
}`
	if dbType == "mysql" {
		conn := `fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DbSet.User,
		setting.DbSet.Password,
		setting.DbSet.Host,
		setting.DbSet.Name))`
		result = strings.ReplaceAll(db, "{type}", "mysql")
		result = strings.ReplaceAll(result, "{conn}", conn)
	}
	return
}

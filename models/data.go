package models

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	//postgresql driver
	_ "github.com/lib/pq"
)

//DB is a pg connection
var db *gorm.DB
var err error
var connStr string

var (
	// host = "192.168.31.137"
	// host     = "huanyu0w0.cn"
	host     = "xiaoheidui.cn"
	port     = "65432"
	dbname   = "postgres"
	user     = "postgres"
	password = "postgres"
)

func init() {
	// initEnv()
	initDB()
	initSchema()
}

func initEnv() {
	if err := env(); err != nil {
		panic(err)
	}
}

func initDB() {
	connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, dbname, password)
	if db, err = gorm.Open("postgres", connStr); err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	// db.LogMode(true)
}

func initSchema() {
	db.AutoMigrate(new(User), new(Timespace), new(Album), new(Chat), new(Favour),
		new(Label), new(Picture), new(Tips))
}

func env() error {
	host = os.Getenv("HOST")
	if len(host) == 0 {
		return fmt.Errorf("缺少环境变量HOST ")
	}
	port = os.Getenv("PORT")
	if len(port) == 0 {
		return fmt.Errorf("缺少环境变量PORT. ")
	}
	dbname = os.Getenv("DBNAME")
	if len(dbname) == 0 {
		return fmt.Errorf("缺少环境变量DBNAME. ")
	}
	user = os.Getenv("USER")
	if len(user) == 0 {
		return fmt.Errorf("缺少环境变量USER. ")
	}
	password = os.Getenv("PASSWORD")
	if len(password) == 0 {
		return fmt.Errorf("缺少环境变量PASSWORD. ")
	}
	return nil
}

//refresh the database connection
func refresh() {
	if db, err = gorm.Open("postgres", connStr); err != nil {
		panic("Failed to connect database: " + err.Error())
	}
}

func errCheck(err error) {
	if err != nil {

		refresh()
	}
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

//Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

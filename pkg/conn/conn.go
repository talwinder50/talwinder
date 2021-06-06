/*
	Package conn handles sophonic connections to trisolan resources
*/
package conn

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SophonicConnection represents a connection to remote resource
type SophonicConnection struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config SophonicConnection) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}

var Connector *gorm.DB

// ConnectSophon establishes a so phonic connection to a remote resource and returns a
// So phonic Connection that can be used for communication

func ConnectSophon(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")

	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

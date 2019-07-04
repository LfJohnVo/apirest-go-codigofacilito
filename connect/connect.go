package connect

import (
	"../structures"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var connection *gorm.DB
//variables de entorno
const engine_sql string = "mysql"
const username string = "root"
const password string = ""
const database string = "taller1"

func InitializeDatabase(){
	connection = ConnectORM(CreateString())
	log.Println("La conexion con la base de datos fue exitosa!")
}

func CloseConnection(){
	connection.Close()
	log.Println("La conexion con la base de datos fue cerrada!")
}

func ConnectORM(stringConnection string) *gorm.DB{
	connection, err := gorm.Open( engine_sql, stringConnection )
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return connection
}

func CreateString() string {
	return username + ":" + password + "@/" + database
}

//funciones para consultar
func GetUser(id string) structures.User{
	user := structures.User{}
	connection.Where("id = ?" , id).First(&user)
	return user
}

func CreateUser(user structures.User) structures.User {
	connection.Create(&user)//Se asigna un id
	return user
}

func UpdateUser(id string, user structures.User) structures.User {
	currentUser := structures.User{}
	connection.Where("id = ?", id).First(&currentUser)

	currentUser.Username = user.Username
	currentUser.First_Name = user.First_Name
	currentUser.Last_Name = user.Last_Name
	connection.Save(&currentUser)

	return currentUser
}

func DeleteUser(id string){
	user := structures.User{}
	connection.Where("id = ?" , id).First(&user)
	connection.Delete(&user)
}
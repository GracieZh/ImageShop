/*
based on https://hugo-johnsson.medium.com/rest-api-with-golang-mux-mysql-c5915347fa5b

go mod init backend
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql

create table `categories` ( `id` int not null AUTO_INCREMENT, `name` varchar(100) not null, primary key(`id`)) engine=InnoDB default charset=utf8mb4;
insert into categories values ('100000', 'red');
insert into categories values (1, 'red');
*/

package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
)

type Category struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

var db *sql.DB
var err error

func WrapHandler(f http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {

      fmt.Println("http request: ", r.Method, r.URL.Path )

      f.ServeHTTP(w, r)
  }
}
func main() {

  fmt.Println("backend starting...")

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/learning")
  if err != nil {
    panic(err.Error())
  }  
  
  defer db.Close()  

  fmt.Println("backend DB OK")

  router := mux.NewRouter()  
  router.HandleFunc("/category", getCategories).Methods("GET")
  router.HandleFunc("/category/", getCategoryByName).Methods("GET")
  router.HandleFunc("/category", setOptions).Methods("OPTIONS")
  router.HandleFunc("/category", createCategory).Methods("POST")
  router.HandleFunc("/category/{id}", getCategory).Methods("GET")
  router.HandleFunc("/category/{id}", updateCategory).Methods("PUT")
  router.HandleFunc("/category/{name}", deleteCategory).Methods("DELETE")  
  router.HandleFunc("/category/{name}", setOptions).Methods("OPTIONS")  

  fmt.Println("backend Listening...\n")

  http.ListenAndServe(":8000", WrapHandler(router))
}

func getCategories(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")  
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  

  fmt.Println("getCategories\n")

  var categories []Category  
  result, err := db.Query("SELECT id, name from categories")
  if err != nil {
    panic(err.Error())
  }  
  defer result.Close()  
  for result.Next() {
    var category Category
    err := result.Scan(&category.ID, &category.Name)
    if err != nil {
      panic(err.Error())
    }
    categories = append(categories, category)
  }  
  json.NewEncoder(w).Encode(categories)
}

func setOptions(w http.ResponseWriter, r *http.Request) {

  fmt.Println("setOptions\n")

  w.Header().Set("Content-Type", "application/json")  
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")  
  w.Header().Set("Content-Length", "0")  

  fmt.Fprintf(w, " ")
}

func createCategory(w http.ResponseWriter, r *http.Request) {

  fmt.Println("createCategory\n")

  w.Header().Set("Content-Type", "application/json")  
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  
  stmt, err := db.Prepare("INSERT INTO categories(name) VALUES(?)")
  if err != nil {
    panic(err.Error())
  }  
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }  
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  name := keyVal["name"]  

  fmt.Println("createCategory name=%s\n", name)

  _, err = stmt.Exec(name)
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("New category was created\n")
  
  result, err := db.Query("SELECT id, name FROM categories WHERE name = ?", name)
  if err != nil {
    panic(err.Error())
  }  
  defer result.Close()  
  var category Category  
  for result.Next() {
    err := result.Scan(&category.ID, &category.Name)
    if err != nil {
      panic(err.Error())
    }
  }  
  json.NewEncoder(w).Encode(category)

}

func getCategoryByName(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  
  params := mux.Vars(r)  

  fmt.Println("getCategoryByName params", params)
  fmt.Println("getCategoryByName name", params["name"])
  fmt.Println("getCategoryByName name ", r.FormValue("name"))

  

  result, err := db.Query("SELECT id, name FROM categories WHERE name like '%"+r.FormValue("name")+"%'")
  if err != nil {
    panic(err.Error())
  }  
  defer result.Close()  
  var categories []Category
  var category Category  
  for result.Next() {
    err := result.Scan(&category.ID, &category.Name)
    fmt.Println("getCategoryByName,",category.Name)
    if err != nil {
      panic(err.Error())
    }
    categories = append(categories, category)
  }  
  fmt.Println("getCategoryByName2")
  json.NewEncoder(w).Encode(categories)
}

func getCategory(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  
  params := mux.Vars(r)  
  result, err := db.Query("SELECT id, name FROM categories WHERE id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }  
  defer result.Close()  
  var category Category  
  for result.Next() {
    err := result.Scan(&category.ID, &category.Name)
    if err != nil {
      panic(err.Error())
    }
  }  
  json.NewEncoder(w).Encode(category)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  
  params := mux.Vars(r)  
  stmt, err := db.Prepare("UPDATE categories SET name = ? WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }  
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }  
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  newName := keyVal["name"]  
  _, err = stmt.Exec(newName, params["id"])
  if err != nil {
    panic(err.Error())
  }  
  fmt.Fprintf(w, "Category with ID = %s was updated", params["id"])
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {

  fmt.Println("deleteCategory")
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")  
  w.Header().Set("Access-Control-Allow-Headers", "*")  
  params := mux.Vars(r)  

  fmt.Println("deleteCategory name ", r.FormValue("name"))

  stmt, err := db.Prepare("DELETE FROM categories WHERE name = ?"+r.FormValue("name"))
  if err != nil {
    panic(err.Error())
  }  
  _, err = stmt.Exec(params["name"])
  if err != nil {
    panic(err.Error())
  }  
  json.NewEncoder(w).Encode("Deleted")
}
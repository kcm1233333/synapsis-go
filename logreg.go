package main

import (
  "database/sql"
  "fmt"
  
  "log"
  "io/ioutil"
  "encoding/json"
  "net/http"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
 )

const (
  host     = "peanut.db.elephantsql.com"
  port     = 5432
  user     = "zkhnlyuj"
  password = "0uAnAFE-OGiQrTFJdhxPjdbz38Sp2xyb"
  dbname   = "zkhnlyuj"
)
type Pengguna struct{
    kodepengguna string
	namapengguna string
	alamatpengguna string
    emailpengguna string
    katasandi string
    kodeotp string	
}
type Body struct{
    body  *gorm.DB
}
func CheckError(err error) {
    if err != nil {
        panic(err)
    } else {
	fmt.Println("Successfully Inserted!")
	}
}
type Response struct {
    status  string `json:"Status"`
}
func  Registration (w http.ResponseWriter, r *http.Request) {
 psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  
    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
    kodepengguna := keyVal["kodepengguna"]
	namapengguna := keyVal["namapengguna"]
	alamatpengguna:=keyVal["alamatpengguna"]
	emailpengguna:= keyVal["emailpengguna"]
	katasandi  :=   keyVal["katasandi"]
	kodeotp    :=   keyVal["kodeotp"]
	
  insertStmt := `insert into "pengguna" values($1, $2, $3,$4,$5,$6)`
   _, e := db.Exec(insertStmt, kodepengguna, namapengguna, alamatpengguna, emailpengguna, katasandi, kodeotp)
   CheckError(e)
   	data := [] struct {
        Status string
        
    } {
        { "Successfully Inserted!" },
        
    }
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)  
}

func  Login (w http.ResponseWriter, r *http.Request) {
 psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  
    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
    kodepengguna := keyVal["kodepengguna"]
	katasandi  :=   keyVal["katasandi"]
	
	
  logStmt := `select "namapengguna" from "pengguna" where kodepengguna=$1 and katasandi=$2`
  login, e := db.Query(logStmt, kodepengguna, katasandi)
   CheckError(e)
  //defer rows.Close()
for login.Next() {
  var namapengguna string
  login.Scan(&namapengguna)
  
  data := [] struct {
        NamaPengguna string
        
    } {
        { namapengguna },
        
    }
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)  
}
}
func main() {
    log.Println("Please do something on iv t !")
    http.HandleFunc("/registration", Registration)
	http.HandleFunc("/login", Login)
    http.ListenAndServe(":5051", nil)
}
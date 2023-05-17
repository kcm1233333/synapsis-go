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
type Barang struct{
    kodebarang string
	namabarang string
	kategoribarang string
    hargabarang int64
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
func  Uang (w http.ResponseWriter, r *http.Request) {
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
	kodebayar := keyVal["kodebayar"]
	metodebayar := keyVal["metodebayar"]
	namapembayar:=keyVal["namapembayar"]
	namabank:=keyVal["namabank"]
	nomorrek:=keyVal["nomorrek"]
		
  uangStmt := `insert into "uang" values($1, $2, $3,$4,$5)`
   _, e := db.Exec(insertCartStmt, kodebarang, metodebayar, namapembayar,namabank,nomorrek)
   CheckError(e)
   	data := [] struct {
        Status string
        
    } {
        { "Successfully Inserted Money!" },
        
    }
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)  
}
func main() {
    log.Println("Please do something on iv t !")
    http.HandleFunc("/money", Uang)
	http.ListenAndServe(":5051", nil)
}
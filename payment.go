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
func  Payment (w http.ResponseWriter, r *http.Request) {
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
	kodetransaksi := keyVal["kodetransaksi"]
	kodekeranjang:=keyVal["kodekeranjang"]
    kodebayar:=keyVal["kodebayar"]		
  bayarStmt := `insert into "transaksi" values($1, $2, $3)`
   _, e := db.Exec(insertCartStmt, kodetransaksi, kodekeranjang, kodebayar)
   CheckError(e)
   	data := [] struct {
        Status string
        
    } {
        { "Successfully Paid!" },
        
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
    http.HandleFunc("/payment", Payment)
	http.ListenAndServe(":5051", nil)
}
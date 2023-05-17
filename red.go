package main

import (
	"encoding/json"
	//"io/ioutil"
	"log"
	//"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// curl localhost:8000 -d '{"name":"Hello"}'
func Cleaner(w http.ResponseWriter, r *http.Request) {
	// Read body
	/*b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}*/
	//user:=Message{}
	//var msg Message
	//mess:="Succesfully Ended!"
	//keyVal := make(map[string]string)
	//keyValInt := make(map[string]int)
    //json.Unmarshal(b, &keyValInt)
	//json.Unmarshal(b, &keyVal)
    //title := keyVal["Name"]
	//id    := keyValInt["Id"]
	//err = json.Unmarshal(b, &mess)
	//if err != nil {
		//http.Error(w, err.Error(), 500)
		//return
	//}
	//params := mux.Vars(r)
    //fmt.Println(title)
	//fmt.Println(id)
	data := [] struct {
        Status string
        
    } {
        { "Successfully Ended with JSON!" },
        
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
	http.HandleFunc("/", Cleaner)
	address := ":8000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

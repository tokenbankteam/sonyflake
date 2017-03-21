package main

import (
	"encoding/json"
	"net/http"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

// config machine id here
func LocalMachineID() (uint16, error) {
	return uint16(1), nil
}

func init() {
	var st sonyflake.Settings
	st.MachineID = LocalMachineID
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := sf.NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(sonyflake.Decompose(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
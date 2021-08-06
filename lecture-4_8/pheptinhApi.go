package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type calculator struct {
	Variable1  float32 `json:"variable1"`
	Variable2  float32 `json:"variable2"`
	Expression string  `json:"expression"`
}

func SetJsonResponse(res http.ResponseWriter, mes []byte, httpCode int) {
	res.Header().Set("Content", "application/json")
	res.WriteHeader(httpCode)
	res.Write(mes)
	fmt.Println("Phan hoi thanh cong")
}
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		mes := []byte(`{"message":"server is running"}`)
		SetJsonResponse(rw, mes, http.StatusOK)
	})

	http.HandleFunc("/calculator", func(rw http.ResponseWriter, r *http.Request) {
		var cal calculator
		data := r.Body
		fmt.Println(data)
		defer r.Body.Close()
		err := json.NewDecoder(data).Decode(&cal)
		if err != nil {
			mes := []byte(`{"message":"err get data"}`)
			SetJsonResponse(rw, mes, http.StatusInternalServerError)
			return
		}
		var ans float32
		switch cal.Expression {
		case "+":
			ans = cal.Variable1 + cal.Variable2
		case "-":
			ans = cal.Variable1 - cal.Variable2
		case "*":
			ans = cal.Variable1 * cal.Variable2
		case "/":
			ans = cal.Variable1 / cal.Variable2
		default:
			ans = 0
		}
		fmt.Println(ans)
		str := fmt.Sprintf("%f", cal.Variable1) + cal.Expression + fmt.Sprintf("%f", cal.Variable2) + "= " + fmt.Sprintf("%f", ans)
		mes := []byte(str)
		SetJsonResponse(rw, mes, http.StatusOK)

	})
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

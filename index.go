package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
)

const Token  = "TellMeK"
func main()  {
	openHttpListen()
}

func openHttpListen()  {
	http.HandleFunc("/wechat", receiveClientRequest)
	err := http.ListenAndServe("127.0.0.1:8081", nil)

	if err != nil {
		fmt.Sprintf("Listen Error : %s", err)
	}
}

func receiveClientRequest(w http.ResponseWriter, r *http.Request){

	r.ParseForm()

	signature := r.Form["signature"]
	timestamp := r.Form["timestamp"]
	nonce := r.Form["nonce"]

	tempArr := []string{signature[0],timestamp[0],nonce[0]}

	sort.Strings(tempArr)


	var tmpStr string

	for _,v :=range tempArr{
		tmpStr += v
	}

	tmpStr = Sha1String(tmpStr)

	if tmpStr == signature[0] {
		w.Write([]byte("true"))
	}else{
		w.Write([]byte("false"))
	}

	return
}

func Sha1String(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}
package main

import ("fmt","log","net/http");

func formHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"PaseForm() err: %v",err)
		return 
	}
	fmt.Fprintf(w, "Post request was successfully")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}


func helloHandle( w http.RepsonseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found -3-", http.StatusNotFound)
		return 
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"Hello I am Mikado Nakamoto")
}



func main() {

	fileServer := http.FileServer(http.Dir('./static')) // Open files from the static folder -3-
	http.Handle("/",fileServer)
	http.HandlerFunc("/form",formHandle)
	http.HandlerFunc("/hello",helloHandle)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListerAndServer(":8000",nil); err != nil {
		log.Fatal(err)
	}

}

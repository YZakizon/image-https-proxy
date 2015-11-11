/** Copyright 2015 Yeffry Zakizon **/

package main

import (
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"time"
	"fmt"
	"os"
)

func main() {

	http.HandleFunc("/", index);
	port := os.Getenv("PROXY_PORT");
	if (port == "") {
		port = "8080"
	}
	log.Println("Listening to port ", port);
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

var etag []map[string]string;


/* Only support path /url
 * example http://www.example.com/proxy?d=http://www.google.com
 */
func index(w http.ResponseWriter, r *http.Request) {
		
	qtmp := r.URL.Query()["d"];
	if (len(qtmp) == 0) {
		fmt.Fprintf(w,"");
		return;
	}
 	
	dest_url, err := url.Parse(qtmp[0]);

	if (err != nil) {
		fmt.Fprintf(w,"");
		return;
		
	}

 	if (dest_url.Scheme == "") {
		dest_url.Scheme = "http";
	}

	start := time.Now();
	client := &http.Client{};
	
	/* Only support GET method */
	req, err := http.NewRequest("GET", dest_url.String(), nil); 
	if (err!=nil){
		log.Println("Create client error ", err);
		fmt.Fprintf(w, "");
		return;
	}
	

	req.Header = r.Header;
	resp, err2 := client.Do(req);
	if (err2 != nil) {
		log.Printf("GET ERROR: %s  err: %s", dest_url.String(), err2);
		fmt.Fprintf(w, "");
		return;
	}

	defer resp.Body.Close();
	body,_ := ioutil.ReadAll(resp.Body);
	for k,v := range resp.Header {
		w.Header().Set(k, v[0]);
	}
	w.Write(body);
	stop := time.Now();
	diff := stop.Sub(start);

	log.Printf("GET %s %d ms\n", dest_url.String(), (diff.Nanoseconds() % 1e9)/1e6);
}

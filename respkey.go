package main

//Importing the required packages
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

//main function starts here
func main() {

	// Variable declartion
	var Url, Srh string
	var Slt bool
	//Usage of flags
	flag.StringVar(&Url, "u", "http://webcode.me", "Entered the target url")
	flag.StringVar(&Srh, "s", "", "Enter the keyword to search")
	flag.BoolVar(&Slt, "silent", false, "To hide the response")
	flag.Parse()

	//Get Request
	//start----->
	//Get function used to get the requested domain
	resp, err := http.Get(Url)
	//handling error
	if err != nil {
		log.Fatal(err)
	}
	//end of Get Request<-------

	//Defer workd in LIFO,Defer statements allow us to think about closing each file right after opening it, guaranteeing that, regardless of the number of return statements in the function, the files will be closed.
	defer resp.Body.Close()

	//ReadAll function
	//start----->
	//ReadAll reads from resp until an error or EOF and returns the data it read.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	//end of ReadAll Request<-------

	//Create function
	//start----->
	//Create creates the file named as out.html.
	f, err := os.Create("out.html")

	if err != nil {
		log.Fatal(err)
	}
	//end of ReadAll Request<-------
	defer f.Close()

	fmt.Println("")
	fmt.Println("===================================================================")
	fmt.Println("[+] Fetching Response: ", Url)
	fmt.Println("====================================================================")
	//f.writeString==(os.file).WriteString.... writes the content in out.html file
	_, err2 := f.WriteString(string(body))

	if err2 != nil {
		log.Fatal(err2)
	}

	//Silent flag condition
	//Silent flag used to not show the source code if not mentioned by user
	//start------>
	if Slt == false {
		fmt.Println(string(body))
	} else {
		fmt.Println("")
	}
	//end<-----------

	fmt.Println("====================================================================")
	//regexp helps to search a word with both lowercase and upprcase
	reg, _ := regexp.MatchString("(?i)"+Srh, string(body))
	// if..else loop if reg is found the search results will be shown
	if reg == true {
		fmt.Println("[+] Search Results:", Srh)
	} else {
		fmt.Println("[-] Search Not Found: ", Srh)
	}

	fmt.Println("====================================================================")
}

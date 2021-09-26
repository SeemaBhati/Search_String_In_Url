package main

//Importing the required packages
import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Hunting(Url, Srh string, Slt bool) {
	//
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
	Fout := strings.Trim(Url, "https://")
	//f, err := os.Create(Fout + ".html")

	f, err := os.Create(filepath.Join("output", filepath.Base(Fout+".txt")))

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
		fmt.Println("[+] Found:", Srh)
	} else {
		fmt.Println("[-] Not Found:", Srh)
	}

	fmt.Println("====================================================================")
}

//main function starts here
func main() {

	// Variable declartion
	var Url, Srh, Flist string
	var Slt bool
	//found := false
	//Usage of flags
	flag.StringVar(&Url, "u", "", "Entered the target url")
	flag.StringVar(&Srh, "s", "", "Enter the keyword to search")
	flag.BoolVar(&Slt, "silent", false, "To hide the response")
	flag.StringVar(&Flist, "l", "", "Enter the list of Urls via file like urls.txt")
	flag.Parse()

	err := os.MkdirAll("output", 0755)
	if err != nil {
		log.Fatal(err)
	}

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name]=true } )

	if flagset["l"] {
		// os.Open() opens specific file in
		// read-only mode and this return
		// a pointer of type os.
		file, err := os.Open(Flist)

		if err != nil {
			log.Fatalf("failed to open")

		}

		// The bufio.NewScanner() function is called in which the
		// object os.File passed as its parameter and this returns a
		// object bufio.Scanner which is further used on the
		// bufio.Scanner.Split() method.
		scanner := bufio.NewScanner(file)

		// The bufio.ScanLines is used as an
		// input to the method bufio.Scanner.Split()
		// and then the scanning forwards to each
		// new line using the bufio.Scanner.Scan()
		// method.
		scanner.Split(bufio.ScanLines)
		var text []string

		for scanner.Scan() {
			text = append(text, scanner.Text())
		}

		// The method os.File.Close() is called
		// on the os.File object to close the file
		file.Close()

		// and then a loop iterates through
		// and prints each of the slice values.
		for _, Lurls := range text {
			Hunting(Lurls, Srh, Slt)

		}
	} else {
		Hunting(Url, Srh, Slt)
	}

}

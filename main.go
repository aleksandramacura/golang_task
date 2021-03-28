
package main

import(
  "fmt"
  "os"
  "regexp"
  "net/http"
  "bufio"
  "log"
)

var name = "Simple CLI"
var usage = "A command line interface which can start HTTP web server"
var author = "Aleksandra Macura"
var version = "1.0.0"


func server(str string){
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("Server reading ...")

        file, err := os.Open(str)

        if err != nil {
          fmt.Println("Wrong file name! Try different name!")
           os.Exit(1)
        }
        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanLines)
        var text []string

        for scanner.Scan() {
            text = append(text, scanner.Text())
        }
        file.Close()

        for _, each_ln := range text {
          fmt.Fprintf(w, each_ln)
        }
    })

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }

}

func main(){

  if (len(os.Args) == 2){

    var s = os.Args[1]

    if (s == "-h" || s == "--help"){
        fmt.Println("\n")
        fmt.Println("NAME: ", name)
        fmt.Println("USAGE: ", usage)
        fmt.Println("AUTHOR: ", author)

        fmt.Println("\n")
        fmt.Println("HELP POSSIBLE COMMANDS: ")
        fmt.Println("--help, -h          Information about program, possible commands. ")
        fmt.Println("--version, -v       Version of the program. ")
        fmt.Println("<index.html>        Run the <index.html> file in the HTTP web server.")
      } else if (s == "-v" || s == "--version"){
        fmt.Println("The version of the program: ", version)
      } else {
        match, _ := regexp.MatchString("\\.html$", s)

        if (match == true){
            server(s)
        }else{
            fmt.Println("Something went wrong")
        }
      }

  }else{
    fmt.Println("\nWrong number of arguments! \nType '-h' for more information")
  }
}

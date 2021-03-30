
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
        scan := bufio.NewScanner(file)
        scan.Split(bufio.ScanLines)
        var text []string

        for scan.Scan() {
            text = append(text, scan.Text())
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

func choice(str string) int {
  if (str == "-h" || str == "--help"){
    //  fmt.Println("\n")
      fmt.Println("NAME: ", name)
      fmt.Println("USAGE: ", usage)
      fmt.Println("AUTHOR: ", author)

    //  fmt.Println("\n")
      fmt.Println("HELP POSSIBLE COMMANDS: ")
      fmt.Println("--help, -h          Information about program, possible commands. ")
      fmt.Println("--version, -v       Version of the program. ")
      fmt.Println("<index.html>        Run the <index.html> file in the HTTP web server.")
      return 0
  } else if (str == "-v" || str == "--version"){
      fmt.Println("The version of the program: ", version)
      return 1
  } else {
      match, _ := regexp.MatchString("\\.html$", str)

      if (match == true){
          server(str)
          return 2
      }else{
          fmt.Println("Something went wrong")
          return 3
      }
  }
}
func main(){

  if (len(os.Args) == 2){

    var s = os.Args[1]
    choice(s)


  }else{
    fmt.Println("\nWrong number of arguments! \nType '-h' for more information")
  }
}

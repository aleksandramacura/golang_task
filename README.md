### golang_task
## Table of contents
* [General info](#general-info)
* [About the code](#about-the-code)
* [More info](#more-info)

## General info
The goal of this task was to prepare a simple command line interface which provides the following commands:

-- version

![obraz](https://user-images.githubusercontent.com/45381820/112763671-b15f9680-9005-11eb-92df-fc88b52a1752.png)

--help

![help](https://user-images.githubusercontent.com/45381820/112763645-9b51d600-9005-11eb-9352-95249c17c473.png)

<index.html>

![obraz](https://user-images.githubusercontent.com/45381820/112763695-cf2cfb80-9005-11eb-81af-0bee24aa7c12.png)

The last command starts the HTTP server which serves the provided HTML file.
	
## About the code
The code uses the following libraries:

"fmt"

"os"

"regexp"

"net/http"

"bufio"

"log"

The code consists of two functions: 

- main - responsible for handling the arguments passes to the console
- server - responsible for starting the HTTP server and sending the HTML file to it 

## More info

For the proper working of the program the HTML file of the given name is needed! 



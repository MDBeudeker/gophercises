package main

import (
  "fmt"
  "os"
  "strings"
  "bufio"
  "flag"
)


func check( e error){
  if e != nil {
    panic(e)
  }
}

func main() {

  // fileReading w flags
  fileNamePtr := flag.String("filename","problems.csv","The quiz file")
  flag.Parse()       // needed to execute CLI parsing
  dat, err := os.ReadFile(*fileNamePtr)
  
  score := 0
  check(err)

  lines := strings.Split(string(dat), "\n")

  for i := 0; i< len(lines)-1; i++ {
    value := strings.Split(lines[i],",")
    fmt.Print(value[0], "\n")
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    if (strings.TrimSpace(input.Text()) == strings.TrimSpace(value[1])){ //compare the user input to the correct answer, Trimspace is to remove the windows line terminator
      score ++
      fmt.Print("that's Correct!!")
    }   else{
      fmt.Print("Wrong answer")
    }
      fmt.Print("\nYour Score is now ",score," out of ", len(lines)-1,"\n")
  }
}

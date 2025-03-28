package main

import (
  "fmt"
  "os"
  "strings"
  "bufio"
  "flag"
  "encoding/csv"
  "log"
  "math/rand"
  "time"
)


func check( e error){
  if e != nil {
    panic(e)
  }
}

func main() {

  // random seed
  rand.Seed(time.Now().Unix())
  // flags for file reading, randomizing the quiz order
  fileNamePtr := flag.String("filename","problems.csv","The quiz file")
  randomPtr := flag.Bool("random",false,"shuffle the quiz order; 0 = regular order (default), 1= shuffle")
  secondsPtr := flag.Uint("timer",30,"allowed time in seconds")

  // once all flags are declared, call flag.Parse() to execute the command-line parsing.
  flag.Parse()
  dat, err := os.Open(*fileNamePtr)
  check(err)
  
  score := 0 // initial score == 0

  filedata, err := csv.NewReader(dat).ReadAll()
  if err != nil {
    log.Println(err)
    return
  }

  // timer function with "time up" text
  timer := time.AfterFunc(time.Duration(*secondsPtr) * time.Second, func() {
    fmt.Println("Time's up!\nYou Scored", score,"out of",len(filedata),"points!")
    os.Exit(4)
  })
  defer timer.Stop()

  // if user wants random order, randomize the file name
  if *randomPtr == true {
    rand.Shuffle(len(filedata), func(i, j int) {
      filedata[i], filedata[j] = filedata[j], filedata[i]
    })
  }

  for _, value := range filedata{
    fmt.Print(value[0], "\n")
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    if (strings.TrimSpace(input.Text()) == strings.TrimSpace(value[1])){ //compare the user input to the correct answer, Trimspace is to remove the windows line terminator
      score ++

    }
  }
  fmt.Print("Done!\nYour Score is ",score," out of ", len(filedata),"\n")
}

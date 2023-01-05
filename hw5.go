//Andrea Labra Orozco CS311 hw5 

package main

import (
  "bufio" 
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "log"
)


//this function searches the lines in a single file for the string needed 
func LinesInFile(fileName string)[] string{
  f, _:=os.Open(fileName) 
        // Create new Scanner.
  scanner:= bufio.NewScanner(f)
  result:= [] string {}
        // Use Scan.
  for scanner.Scan() {
    line:= scanner.Text()
    // Append line to result.
    result = append(result, line)
  }
  return result
}

func main() {
  var text string
  //var res string
  //ask user which string they want to find in the files 
  fmt.Print("Enter text: ")
  // get the sub string to search from the user
  fmt.Scanln( & text)
  //this will go into the directory called files and load them 
  files, err:= ioutil.ReadDir("files") //"files", directory path
  if err != nil {
    log.Fatal(err) //error catching 
  }
  sub:= string(text)
  for _, file:= range files {
    filename:= file.Name()
    //this will open all of the files 
    fmt.Printf("reading from file: %s%s", filename, "\n")
    for index,
    //this will get the line from filename 
    //schedule files 
    line:= range LinesInFile("files/" + filename) {
      stringArr:= [] string {line}
      str:= strings.Join(stringArr, " ")
      //set it back to false 
      res:= false;
      res = strings.Contains(str, sub)
      if res == true {
        //this will print the actual string right under the text file that is found 
        fmt.Printf("line = (%v), %v\n", index + 1, line)
      }
    }
  }
}
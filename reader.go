//I am not a genius. I seek help. This is a cargo-cult code.
//Source http://stackoverflow.com/a/18479916
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

func main() {
  dict := make(map[string]uint)
  lines, err := readLines("foo.txt")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }else{

	  for _, line := range lines {
	  	//TODO Start indexing
      words := strings.Split(line," ")
	    for _  ,word := range words{
        if dict[word] >0 {
            dict[word] += 1;
          }else{
            dict[word] = 1;
          }
      }
	  }
    fmt.Println()
    for key,value := range dict{
      fmt.Println(key, value)  
    }
    
	    // if err := writeLines(lines, "foo.out.txt"); err != nil {
		 //   log.Fatalf("writeLines: %s", err)
		// }

  }



}
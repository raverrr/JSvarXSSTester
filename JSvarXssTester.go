package main

import (
  "strings"
  "bufio"
  "fmt"
  "net/http"
  "os"
  "regexp"
  "time"
)

func main() {
  // Compile regular expression to match variable declarations
  r := regexp.MustCompile(`var ([a-zA-Z0-9_]+)`)

  // Create a map to store unique variable names
  names := make(map[string]bool)

  // Read URLs from standard input
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    url := scanner.Text()

    // Make GET request to URL
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      continue
    }
    defer resp.Body.Close()

    // Scan response body for variable declarations
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
      line := scanner.Text()
      matches := r.FindAllStringSubmatch(line, -1)
      if matches != nil {
        for _, match := range matches {
          name := match[1]
          if len(name) >= 2 && !names[name] {
            names[name] = true
          }
        }
      }
    }

    // Divide variable names into groups of no more than 5 names
var nameGroups [][]string
for len(names) > 0 {
  group := make([]string, 0, 5)
  for name := range names {
    group = append(group, name)
    if len(group) == 5 || len(names) == 0 {
      // Append the group to the nameGroups slice and break out of the inner loop
      nameGroups = append(nameGroups, group)
      break
    }
  }
  for _, name := range group {
    delete(names, name)
  }
}

// Create a variable to keep track of the count
count := 0

    // Make GET request for each group of names
    for _, group := range nameGroups {
      // Build new URL with group of names as query parameters
      newURL := url + "?"
      for i, name := range group {
        if i > 0 {
          newURL += "&"
        }
        newURL += fmt.Sprintf("%s=zzzz1", name)
      }

// Print the URL with added parameters
  fmt.Printf("\x1b[32m%s\x1b[0m", "Requesting URL:")
  fmt.Println(newURL)

      // Make GET request to new URL
      resp, err := http.Get(newURL)
      if err != nil {
        fmt.Fprintln(os.Stderr, err)
        continue
      }
      defer resp.Body.Close()

  // Check if "zzzz1" is present in response body
  scanner := bufio.NewScanner(resp.Body)
  for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, "zzzz1") {
      // Increment the count if "zzzz1" is found
      count++
    }
  }

fmt.Printf("\x1b[94m%s\x1b[0m", "zzzz1 found ") // Prints "zzzz1 found" in a brighter blue color
fmt.Printf("\x1b[31m%d\x1b[0m ", count) // Prints the count in red
fmt.Printf("\x1b[94m%s\x1b[0m", "times in response body of: ") // Prints "times in response body of" in a brighter blue color
fmt.Printf("%s\n", newURL) // Prints newURL in the default color (not blue or red)

                // Sleep for specified time before making another request
          time.Sleep(100 * time.Millisecond)
        }
  }
 }

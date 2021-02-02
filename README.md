# Go Webhook Receiver Example

A general purpose Golang application that receives POST requests, performs logical tests on the incoming payload's parameters and makes OS executions.

## Example
```
package main

import wh "github.com/corinz/go-webhook-receiver"

func  main() {
  ...
  wh.ExecuteThisWhen("./slack-notify.sh", "author eq corinz")
  wh.Startup("/") // http://localhost:8080
  
}
```


## Logical Tests

This project has basic logical tests that are implemented using native logical operators
|                |User Syntax                          |Go Operator                        |
|----------------|-------------------------------|-----------------------------|
|Equals|`eq`            |`==`           |
|Not equal          |`ne`            |`!=`         |
|Greater than         |`gt`|`>`|
|Less than         |`lt`|`<`|


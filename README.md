# GOLANG Easy Logging
## Why?
I made this because I'm bad at ~~coding~~ golang and just wanted a way to easily log messages to a text file.
Added benefit of making any projects look cleaner with this repo I guess idk.

## Install
Like any other go package run `go get "github.com/Github-Reneon/easy_logging/"`

## Usage
### Create a FileWriter object:
>`logExample := FileWriter {"nameoffile.txt"}`

### Write to a file:
>`logExample.Write("Message", false)`

### Write to a file and display the message in console:
>`logExample.Write("Message", true)`


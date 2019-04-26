# Moody

## Used Technologies
- Go (server)
- MongoDB (database)
- JavaScript, HTML5, CSS 

## User Guide
https://docs.google.com/document/d/1EO76c7XOyqi8w4lr9s9Od6gZDPFmanHWzKvHjmsQWNc

## Installation
$```go get github.com/mongodb/mongo-go-driver```<br>
The output of this may look like a warning stating something like _package github.com/mongodb/mongo-go-driver: no Go files in (...)._<br>
This is expected output. If you are using the dep package manager, you can install the main mongo package as well as the bson and mongo/options package using this command:<br>
$```dep ensure --add github.com/mongodb/mongo-go-driver/mongo github.com/mongodb/mongo-go-driver/bson github.com/mongodb/mongo go-driver/mongo/options```<br>
And then again,<br>
$```go get github.com/mongodb/mongo-go-driver```<br>


## Usage
$ ```export db_uri=your_mongo_db_uri```<br>
$ ```go run main.go```<br>
Site is now available at http:localhost:8080

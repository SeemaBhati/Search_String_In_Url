# RespKey
RespKey is a tool through which you can search keywords in the response body.

## Installation
```
go get github.com/SeemaBhati/respkey
```

## Usage
```
./respkey.go -u http://webcode.me -s Hello 
```
Use ``` -silent ``` flag to hide the response body
```
./respkey.go -u http://webcode.me -s Hello -silent
```

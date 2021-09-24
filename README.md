# RespKey
RespKey is a tool through which you can search keywords in the response body.

## Installation
```
git clone https://github.com/SeemaBhati/respkey.git
```

## Usage
```
go run respkey.go -u http://webcode.me -s Hello 
```
Use ``` -silent ``` flag to hide the response body
```
go run respkey.go -u http://webcode.me -s Hello -silent
```

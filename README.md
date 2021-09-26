# RespKey
RespKey is a tool through which you can search keywords in the response body.

## Installation
```
go get -u github.com/SeemaBhati/respkey
```

## Usage
```
./respkey -u http://webcode.me -s Hello
```
```
./respkey -l .\urls.txt -s Hello -silent
```
Use ``` -silent ``` flag to hide the response body
```
./respkey -u http://webcode.me -s Hello -silent
```

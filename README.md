# RespKey
RespKey is a tool through which you can search keywords in the response body.

## Installation
```
go get -u github.com/SeemaBhati/respkey
```

## Usage
```
./respkey -u http://webcode.me -s Hello 
go run .\perfect_respkey.go -u https://webcode.me -s Hello
```
Use ``` -silent ``` flag to hide the response body
```
./respkey -u http://webcode.me -s Hello -silent
go run .\perfect_respkey.go -u https://dell.com -s google -silent 
```
Use ''' -l ''' flag to search for list of response body
'''
  go run .\perfect_respkey.go -l .\urls.txt -s admin -silent
'''

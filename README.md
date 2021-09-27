# RespKey
RespKey is a tool through which you can search keywords in the response body.

## Description 

RespKey is used to find sensitive keyword in the response body or a list of urls endpoints. 
If the keyword is found it is displayed as "Found" if the keyword is not found it is displayed as " Not Found".

RespKey can be used in testing, cyber-security.

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





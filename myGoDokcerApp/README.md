docker build -t abc:v2  .



docker run -it -p 8090:8090 -v /Users/wei/importgo/src/myGoDokcerApp/app:/go/src/github.com/user/myProject/app  abc:v2



# go-modules
trying out go.mod go modules dependency manager


when imports are from private repositories `go build` after `go mod init` will not work

needs to set env variable, either on system level
``` 
export GOPRIVATE=github.com/privateRepoName/*
```
or go environment level
```
go env -w GOPRIVATE=github.com/privateRepoName/*
```

both should work

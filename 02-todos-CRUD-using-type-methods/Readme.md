
## GO.MOD 
```go mod init github.com/USERNAME/REPONAME```sh
```go mod init github.com/USERNAME/REPONAME/FOLDER```sh

### Running the project
```go env -w GO111MODULE=off```sh
Using local mod file
```go run -modfile=./local.go.mod ./...```sh


## Guides
- [set GO111MODULE=off](https://stackoverflow.com/a/67598174/15039733)
- [how-to-import-local-files-packages-in-golang](https://linguinecode.com/post/how-to-import-local-files-packages-in-golang)

### doc references: 
-    [dot]: (https://golang.org/ref/mod#go-mod-file-ident)
-   [private repos]: (https://golang.org/ref/mod#private-module-proxy-direct_)
-   [branches]: (https://golang.org/ref/mod#vcs-branch_)

also see: [possible since 1.11](https://stackoverflow.com/a/55302537/15039733) but this method not worked for me

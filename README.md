# emediams
Microservices to manage media entretaiment

# How to create a cli in go with cobra

Create the go project 
```shell
go mod init <my project>
```

Install cobra
```shell
go get -u github.com/spf13/cobra/cobra
```

Create the cobra boilerplate

```shell
cobra init --pkg-name <my project>
```

# Troubleshooting

## cobra: command not found

```shell
vi ~/.zshrc
```

```shell
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=${PATH}:$GOBIN
```

```shell
source ~/.zshrc
```
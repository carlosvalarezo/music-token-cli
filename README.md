# emediams
Microservices to manage media entretaiment

# How to create a cli in go with cobra

Create the go project 
```shell
go mod init my project
```

Install cobra
```shell
go get -u github.com/spf13/cobra/cobra
```

Create the cobra boilerplate

```shell
cobra init --pkg-name my project
```
# Add commands to the cli
```shell
cobra add my-new-command
```

# Adding subcommands

To add a subcommand execute the same cobra command to add a new command.
Then go to the file you want to be the primary command of this new command.
Edit the init method with the primary command. 

```go
func init() {
   artistCmd.AddCommand(bynameCmd)
}
```

The easy way:
```shell
cobra add name -p "artistCmd"
```

# Adding flags

Global flags are in root.go

Local flags go in the respective command


# Running the cli

To build a particular command use
```shell
go run . my-project subcommand1 subcommand2
```

Build the project & running as binary 
```shell
go build -o my-project
```

```shell
./my-project
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

**NOTE**: Always use camel case rather than snake case because that may lead to errors
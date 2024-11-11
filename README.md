
# Go (Golang)

## Environment

* Golang  
```console
sudo apt install golang-go
```

* Visual Studio Code - the preferred code editor in this course
* Go - a Visual Studio Code extension to help with Go development
* Prettier - a code formatter for Visual Studio Code. Note: This screencast demos in this course leverage the automatic "Format on Save" feature of Visual Studio Code, with Prettier as the default formatter.

* Set your GOPATH dir (optional)  
```
go env -w GOPATH=$HOME/git/golang
```

## 
O primeiro passo é criar a pasta hello, que irá conter o nosso primeiro script em Go. Crie a pasta/hello dentro da pasta src do seu Go Workspace, que criamos anteriormente no exercício de Preparando o Ambiente.

```
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
```
<h1 align="center"> Go-Logger </h1>

<p align="center">A logging library for Go applications.</p>

<p align="center">
 <a href="#technologies">Technologies</a> •
 <a href="#running">Running</a> •
 <a href="#author">Author</a>
</p>

![image](https://github.com/jonh-dev/go-logger/assets/101439670/1c8a636c-a580-4b47-9b55-9739442b48df)

##

### Technologies

The following tools were used in building the project:

- Go
- IDE Visual Studio Code

##

### Running

**1.** First, you need to install Go on your system. You can do this by following the instructions at the following link: https://golang.org/dl/

**2.** Choose an IDE of your choice, in this case we will use Visual Studio Code. To download it follow the link: https://code.visualstudio.com/download

**3.** Open your terminal and use `go get` to download and install the `go-logger` library. Replace `github.com/jonh-dev/go-logger/logger` with the path to your `go-logger` library:

```bash
$ go get github.com/jonh-dev/go-logger
```

**4.** Now you can import the `go-logger` library into your Go project. Here's an example of how you can do this:

```go
import (
    "github.com/jonh-dev/go-logger"
)
```

**5.** Now you can use the `go-logger`library in your code. Here’s an example of how you can do this:

```go
func main() {
    log := logger.NewLogger()

    log.Info("This is an info message")
    log.Success("This is a success message")
    log.Warning("This is a warning message")
    log.Error("This is an error message")
    log.Panic("This is a panic message")
    log.Fatal("This is a fatal message")
}
```

**6.** Finally, run the project using the command go run main.go to run the application.

##

### Author

![avatar](https://user-images.githubusercontent.com/101439670/181940218-4f68ffb9-0d35-40df-b8e9-86629333d244.png)

Made by Jonh Dev 🙏

[![LinkedIn Badge](https://img.shields.io/badge/-LINKEDIN-blue?style=flat-square&logo=Linkedin&logoColor=white&link="https://www.linkedin.com/in/jo%C3%A3o-carlos-schwab-zanardi-752591213/)](https://www.linkedin.com/in/jo%C3%A3o-carlos-schwab-zanardi-752591213/)


# Learn Go

This repo is made with the purpose of learning Go programming from scratch with a practical approach


## Installation

Open the MSI file you downloaded and follow the prompts to install Go.
- By default, the installer will install Go to Program Files or Program Files (x86). 
- You can change the location as needed. 
- After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

Verify that you've installed Go.
- In Windows, click the Start menu.
- In the menu's search box, type cmd, then press the Enter key.
- In the Command Prompt window that appears, type the following command:
 ```bash
 $ go version
 ```
 Confirm that the command prints the installed version of Go.
    
## VS Code extension

Make sure to also install the Go extension which makes it easier to work with Go in VS Code.
https://code.visualstudio.com/docs/languages/go

## Features

- Go is a statically typed, compiled language designed for simplicity and efficiency.
- It is known for its strong concurrency support, garbage collection, and built-in testing framework.
- Its often used for building web servers, networking tools, and cloud services.
- Go's syntax is similar to C
- Is designed to be simple and efficient, with a focus on concurrency and parallelism.
- It has a rich standard library, which includes packages for I/O, networking, and cryptography.
- Supports multiple programming paradigms, including procedural, object-oriented, and functional programming.
- Has a strong emphasis on simplicity and clarity, which makes it easy to read and maintain code.
- Has a built-in testing framework, which makes it easy to write and run tests for your code.
- Has a strong community and a large ecosystem of third-party libraries and frameworks.
- Its a compiled language, which means that it is translated into machine code before it is executed


## Run Locally

Go to the project directory

```bash
  cd my-project
```

Install dependencies and check for version

```bash
  go version
```

Must display the message
```bash
  go version go1.24.5 windows/amd64
```

Build the project

```bash
  go build <your_project_name>.go
```

Run your program
```bash
  go run <your_project_name>.go
```

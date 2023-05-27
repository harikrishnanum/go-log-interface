# go-log-interface

go-log-interface is a Go project that demonstrates how to use interfaces to switch between different logging libraries seamlessly. By utilizing an interface for logging and a configuration file, you can easily change the underlying logging library without modifying log statements scattered throughout your codebase.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher) installed on your machine

### Run Instructions

1. Clone the repository:

```shell
git clone https://github.com/your-username/go-log-interface.git
```

2. Navigate to the project directory:

```shell
cd go-log-interface
```

3. Use `go mod` to download the project dependencies:

```shell
go mod tidy
```

### Configuration

Before running the project, make sure you have a `config.json` file in the root directory with the following structure:

```json
{
  "logger": "logrus"
}
```

Replace the value of `"logger"` with the desired logging library (`zap`, `logrus`.). 
### Running the Project

Execute the following command to run the project:

```shell
go run main.go
```

This will initialize the logger based on the configuration specified in `config.json` and demonstrate how log statements can be switched between different logging libraries by utilizing interfaces.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

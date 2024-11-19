# Termec Chat Server

This is a simple chat server that uses the TCP protocol. It is written in Golang and uses the standard library for networking. The server is capable of handling multiple clients at once and can send messages to all connected clients. The server also has a command line interface for sending messages to all clients.

The server can be consumed by any client that can connect to a TCP server or using the provided client written in Golang, [Term Chat application](https://github.com/soupaulodev/termec-chat-client).

## Usage

To run the server, you need to have Golang 1.23.2 installed on your machine. You can download Golang from the [official website](https://golang.org/).

To verify that Golang is installed on your machine, run the following command:

```sh
go version
```

To run the server, clone this repository and navigate to the root directory of the project. Then run the following command:

```sh
go run main.go
```

## License

This project is licensed under the MIT License - see the [GPL-2.0 license](https://github.com/soupaulodev/termec-chat-server/blob/main/LICENSE) file for details.

## Contributing

First of all, thank you for considering contributing to this project. Any help is appreciated. If you want to contribute, follow these steps:

1. Fork the project
2. Create a new branch (`git checkout -b feature/feature-name`)
3. Make the changes
4. Commit the changes
5. Push to the branch (`git push origin feature/feature-name`)
6. Create a new Pull Request

## Author

Paulo Marques - [soupaulodev](https://soupaulodev.com.br)

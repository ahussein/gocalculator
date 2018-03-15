[![Build Status](https://travis-ci.org/ahussein/gocalculator.png)](https://travis-ci.org/ahussein/gocalculator)

# Go Calculator Service
This is a simple calculator implemented as a service with a server side and a client side. The communication between the client and server is implemented via RPC using gRPC.

## Structure
The code is structed in three packages
- pb: the folder contain the protocol buffer service definitions and the generated go code.
- server: contains the server package that run the service and listending for requests.
- client: contians the client package that provide a CLI to interact with the calculator service.
- calculator: contians the bussiness logic of the calculator and unittests for it.

## Requirements
To run the service you will need to have the following
- golang version >= 1.9

## Installtion
To install the service you will need to run the following command (syntax in Linux system bash terminal)
```bash
go get -u github.com/ahussein/gocalculator/...
```
This will install the service and its dependencies, then the srever, client commands will be installed into your $GOPATH/bin directory.

## Running the service
To run the service you need to run the following command
```bash
$GOPATH/bin/server 
2018/03/15 13:21:27 Starting server on port :9000
```
This will start the service and listen on port 9000

## Using the client
To use the service you can use the client CLI, the service has been implemented to be very simple and offer the following functionalities
- Add two integers
- Subtract two intgers
- Multiply two integers
- Divide two integers

To use the cli you can run the following example to add two integers
```bash
client -Operation=add -Param1=10 -Param2=5
15
```
To check the help of the CLI, you can type
```bash
client --help
Usage of client:
  -Address string
        Address to the calculator server (default "localhost:9000")
  -Operation string
        Operation to execute, one of (add, subtract, multiply, divide)
  -Param1 int
        First parameter
  -Param2 int
        Second parameter
```
You can always run the client CLI agianst a remote server on different machine providing the ipaddress:port
```bash
client -Address=localhost:9000 -Operation=add -Param1=10 -Param2=5
15
```

## Run the tests
You can run the unittests using the following command
```bash
go test -v github.com/ahussein/gocalculator/calculator
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
--- PASS: TestSubtract (0.00s)
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
=== RUN   TestDivideOK
--- PASS: TestDivideOK (0.00s)
=== RUN   TestDivideByZero
--- PASS: TestDivideByZero (0.00s)
PASS
ok      github.com/ahussein/gocalculator/calculator     0.011s
```

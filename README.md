# gRPC-Protobuf
A simple application showcasing a golang based server and a python based client, communicating using gRPC

## Genereating gRPC files


### Golang

Install protoc on Linux 

```bash
sudo apt install protobuf-compiler
protoc --go_out=<path_to_project> --go-grpc_out=<path_to_project> api.proto
```

### Python 

You can generate these with protoc but I believe this way is easier to get started.

```bash
python -m pip install grpcio
python -m pip install grpcio-tools
python -m grpc_tools.protoc -I . --python_out=<path_to_project> --pyi_out=<path_to_project> --grpc_python_out=<path_to_project> api.proto
```

## Running this project
Start the golang server

```bash
# Assuming you are in the root of this project
cd golang
go run main.go
```

finally start the client in a new terminal

```bash
# Assuming you are in the root of this project
cd python
python client.py
```

## NOTE
Feel free to modify or use this porject as you like. Thanks.

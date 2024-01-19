#!/bin/bash

# Open a new terminal and run server.go
gnome-terminal -- go run server.go

# Open another terminal and run client.go
sleep 1 # Give some time for the server to start before running the client
gnome-terminal -- go run client.go

#! /bin/bash

# if go is installed, build the binary and then run it. If go is not installed, run the currenty buit binary. If there is no current build let the user know they need go installed to build, then exit
if ! command -v go &> /dev/null
then
    if [ -f ./build/sac ]; then
        echo "Running..."
        ./build/sac
    else
        echo "Go is not installed. Please install Go to build the project."
        exit 1
    fi
else
    echo "Building..."
    go build -o ./build/sac main.go
    echo "Running..."
    ./build/sac
fi

## Install Go
    sudo tar -xvf go1.16.3.linux-amd64.tar.gz
    sudo mv go /usr/local

    sudo nano ~/.bashrc
    // add to end of file
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

    source ~/.profile

    go version

## Clone Repo
    cd $GOPATH
    mkdir src && cd src
    mkdir github.com && cd github.com
    mkdir fathimtiaz && cd fathimtiaz

    git clone git@github.com:fathimtiaz/galaxy-merchant.git
    cd galaxy-merchant
    mkdir logs

## Config
    fill out config/config.json values

## Run App
    go mod init

    go build main.go

    ./main

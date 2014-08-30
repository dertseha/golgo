#!/bin/bash  
# Update system to allow further installations
sudo apt-get update
# Install dependencies for go-gl/glew3
sudo apt-get install xorg-dev 
sudo apt-get install libglew-dev
sudo apt-get install libglfw-dev
# Install dependencies for gvm
sudo apt-get install curl git mercurial make binutils bison gcc build-essential

# Install gvm to manage Go versions
curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer >gvm-installer.sh
bash ./gvm-installer.sh
source /home/ubuntu/.gvm/scripts/gvm
# Select go1.3 for gopherjs
gvm install go1.3
gvm list
gvm use go1.3
# Revert GOPATH to be the original value drone.io had
export GOPATH=~

# Retrieve and start Selenium server. Do this early in order to have it running.
sudo start xvfb
wget http://selenium.googlecode.com/files/selenium-server-standalone-2.35.0.jar --quiet
java -jar selenium-server-standalone-2.35.0.jar &

# Clone, compile and install glfw3 -- note special flags
pushd ~
git clone https://github.com/glfw/glfw.git
cd glfw
cmake -DCMAKE_INSTALL_PREFIX:PATH=/usr -DBUILD_SHARED_LIBS=on -G "Unix Makefiles"
make
sudo make install
popd

# Get all dependencies
go get -u github.com/gopherjs/gopherjs
go get -u github.com/gopherjs/webgl
go get

# Build both native and browser binaries
go build
$GOPATH/bin/gopherjs build -o ./build/golgo-js.js golgo-js/golgo-js.go

# Get further dependencies for tests
go get gopkg.in/check.v1
go get bitbucket.org/tebeka/selenium

# Run all tests
go test ./...

# Run coverage analysis
go get github.com/axw/gocov/gocov
go get github.com/mattn/goveralls
goveralls -v -service drone.io $COVERALLS_TOKEN

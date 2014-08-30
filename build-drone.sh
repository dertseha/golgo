# Print current values
echo "GOROOT: " $GOROOT
echo "GOPATH: " $GOPATH

# Update system and install development packages
sudo apt-get update
sudo apt-get install xorg-dev 
sudo apt-get install libglew-dev
sudo apt-get install libglfw-dev
sudo apt-get install curl git mercurial make binutils bison gcc build-essential

# Install gvm to manage Go versions
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source /home/ubuntu/.gvm/scripts/gvm
gvm install go1.3
gvm list
gvm use go1.3
echo "GOROOT: " $GOROOT
echo "GOPATH: " $GOPATH
export GOPATH=~
echo "GOPATH: " $GOPATH

# Retrieve and start Selenium server. Do this early in order to have it running.
sudo start xvfb
wget http://selenium.googlecode.com/files/selenium-server-standalone-2.35.0.jar --quiet
java -jar selenium-server-standalone-2.35.0.jar &

pushd ~
git clone https://github.com/glfw/glfw.git
cd glfw
cmake -DCMAKE_INSTALL_PREFIX:PATH=/usr -DBUILD_SHARED_LIBS=on -G "Unix Makefiles"
make
sudo make install
popd

go get -u github.com/gopherjs/gopherjs
go get -u github.com/gopherjs/webgl
# CGO_CFLAGS="-I/usr/local/include" CGO_LDFLAGS="-L/usr/local/lib" go get 
go get

go build
$GOPATH/bin/gopherjs build -o ./build/golgo-js.js golgo-js/golgo-js.go

go get gopkg.in/check.v1
go get bitbucket.org/tebeka/selenium

go test ./...

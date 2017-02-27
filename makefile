# Replace demo with your desired executable name
appname := clusterioIRCbridge
sources := $(wildcard *.go)
build = GOOS=$(1) GOARCH=$(2) go build -o build/$(appname)_$(1)$(3)
.PHONY: all windows darwin linux clean
all: windows darwin linux

clean:
	rm -rf build/

##### LINUX BUILDS #####
linux: build/linux_386.tar.gz
build/linux_386.tar.gz: $(sources)
	$(call build,linux,386,)

##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_amd64.tar.gz
build/darwin_amd64.tar.gz: $(sources)
	$(call build,darwin,amd64,)

##### WINDOWS BUILDS #####
windows: build/windows_386.zip
build/windows_386.zip: $(sources)
	$(call build,windows,386,.exe)
# We don't need 64 bit or ARM builds except for mac

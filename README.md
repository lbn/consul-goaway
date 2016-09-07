# consul-goaway
Force-leave failed nodes which have "-group-" in their names every n seconds

## Usage
Edit Consul address and interval in  `Makefile`, then:
```sh
go get github.com/constabulary/gb/...

make run
```

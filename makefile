DIST:=greetrpcdist

$(DIST)/greet.pb.go: greet.proto $(DIST)
	protoc -I=. --go_out=$(@D) $<

$(DIST):
	mkdir $@

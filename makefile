DIST:=greetrpcdist

$(DIST)/greet.pb.go: greet.proto $(DIST)
	protoc -I=. --go_out=$(@D) --go_opt=paths=source_relative $<

$(DIST):
	mkdir $@

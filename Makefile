build:
	docker build -t s4cat .
	docker run --rm s4cat cat /go/bin/s4cat > s4cat
	chmod u+x s4cat

clean:
	rm s4cat

.PHONY: clean

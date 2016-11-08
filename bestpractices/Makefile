testapi:
	dredd api.apib http://localhost:8080/

genapidoc:
	rm -rf docs
	mkdir docs
	aglio -i api.apib -o docs/index.html -t flatly
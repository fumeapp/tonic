PACKAGE=tonic

ray:
	@go get github.com/octoper/go-ray


unray:
	@go mod tidy -e github.com/octoper/go-ray

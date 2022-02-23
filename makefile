test-dev:
	APP_ENV=dev go test -v ./tests
test:
	go test -v ./tests
pkg: test
	newversion.py version
	jfrog "rt" "go-publish" "go-pl" $$(cat version) "--url=$$GOPROXY_API" --user=$$GOPROXY_USER --apikey=$$GOPROXY_PASS
	git tag $$(cat version)
push:
	git push origin $$(cat version)


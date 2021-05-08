VERSION=`grep "Version" textwrap.go | grep -o -E '[0-9]\.[0-9]{1,2}\.[0-9]{1,2}'`

test:
	go test -v ./...

release: check-version check-master
	git tag v${VERSION} && \
	git push origin v${VERSION}

#
# Checking rules
# https://stackoverflow.com/a/4731504
#

check-version:
ifdef VERSION
	@echo Current version: $(VERSION)
else
	$(error VERSION is not set)
endif

check-master:
ifneq ($(shell git rev-parse --abbrev-ref HEAD),main)
	$(error You're not on the "main" branch)
endif


# Setup actions
# https://github.com/mvdan/github-actions-golang
# https://github.com/goreleaser/goreleaser-action/blob/50de962f8481ce3586a634faaa1ef61523dcb5e4/README.md
#

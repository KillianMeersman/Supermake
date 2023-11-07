package parse_test

import (
	"testing"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
)

const Makefile = `
help_attach = "attach - Attach a tty to stdin/stdout of a running container for debug sessions. Stop detachment with Ctrl+p Ctrl+q" # testing... sss
help_build = "build - Build the production justrussel api service docker image"
help_console = "console - Start a console running bash"
help_deploy = "deploy - Deploy the service to Google Cloudrun"
help_publish = "publish - Tag and push the docker image"
help_server = "server - Start a development server"
help_start = "start - Start all services in production mode"
help_test = "test - Run all unit and integration tests"

export ENVIRONMENT ?= staging
export API_VERSION ?= local
export API_VERSION_CLEAN = $(shell echo $(API_VERSION) | sed 's/\./-/g')
export API_IMAGE := registry.test/backend/api:$(API_VERSION)
export AUXILIARY_VERSION ?= $(API_VERSION)
export NAMESPACE := $(ENVIRONMENT)
export SERVICE_ACCOUNT := $(ENVIRONMENT)
TAG = $(shell cat /dev/urandom | tr -dc 'a-z0-9' | fold -w 32 | head -n 1)

export SENTRY_ORG := my-org

STAGING_DIR = $(shell mktemp -d)

# This help target is first because it is the default when
# not specifying one.
help: anotherhelp
	@python:3 sh -c
	# echo "$(ls -l /etc)" > test.txt
	ls -l
	pwd
	echo $(STAGING_DIR)
	echo $(help_attach)
	echo $(help_build)
	echo $(help_console)
	echo $(help_deploy)
	echo $(help_publish)
	echo $(help_server)
	echo $(help_start)
	echo $(help_test)

anotherhelp: anotherhelp2
	echo testing!!!

anotherhelp2:
	echo "Hello from local shell $NAMESPACE"

	@local python3 -c
	import os
	print("Hello from local python", os.environ["NAMESPACE"])

	@alpine:3 sh -c
	echo "Hello from docker shell $NAMESPACE"
	echo 'docker shell' > test-docker-sh.txt

	@python:3 python -c
	import os
	print("Hello from docker python", os.environ["NAMESPACE"])
	open("test-docker-python.txt", "w+").write("docker python")

	@alpine:3 sh -c
	echo "hello $NAMESPACE!"

	test_a: test_b
		echo "test_a is running!"
		echo "test_a done"

	test_b:
		echo "test_b is running!"
		echo "test_b done"

	test_c:
		echo "test_c is running!"
		echo "test_c done"

deploy:
	ssh killy@killy.space "HOST=killy.space TAG=$(TAG) docker-compose --profile local_registry up -d"
	ssh killy@killy.space "HOST=killy.space TAG=$(TAG) docker system prune -af"
`

func TestParsing(t *testing.T) {
	file, err := parse.ParseSupermakeString(".", Makefile)
	if err != nil {
		t.FailNow()
	}

	// Do some sampling
	target, ok := file.Targets["anotherhelp2::test_a"]
	if !ok {
		t.FailNow()
	}

	// Test dependencies
	wantedTargets := []string{"anotherhelp2::test_b", "anotherhelp2::0", "anotherhelp2::1", "anotherhelp2::2", "anotherhelp2::3", "anotherhelp2::4", "anotherhelp2::test_a::0"}
	for i, dep := range target.Dependencies {
		if dep != wantedTargets[i] {
			t.Fail()
		}
	}
	if len(target.Steps) != 0 {
		t.FailNow()
	}

	target, ok = file.Targets["anotherhelp2::test_a::0"]
	if !ok {
		t.FailNow()
	}
	if len(target.Steps) != 1 {
		t.FailNow()
	}

	deploy, ok := file.Targets["deploy::0"]
	if !ok {
		t.FailNow()
	}
	if len(deploy.Steps) != 1 && len(deploy.Steps[0].GetShellCommands()) != 2 {
		t.FailNow()
	}
}

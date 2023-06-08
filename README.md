# Supermake

Supermake is a superset of Makefiles that serves as a modern CI pipeline. Supermake can be ran locally as well as on remote CI server(s).

Features:
1. Nested targets & automatic parallelisation.
2. Easy-to-use containerization of steps.
3. Compatibility with existing Makefiles.

## User guide
This user guide will guide you through the Supermake featureset by starting with an existing makefile, then improving it step-by-step.

Target dependencies can be expressed in the usual Makefile form.

```Makefile
build: test
	docker build -t myimage .

test:
	python manage.py test
	echo 'All tests passed :)'
```

### Using containers
It can be useful to run certain steps inside a Docker container. This can be achieved by using the following notation on its own line:

```
@image:[tag] [entrypoint]
```

These are called environments. All following steps, up until another environment or the end of the target, will be ran inside this container.

When a container is started, the working directory is copied over into the container's working directory. After the container-specific steps are done, all files are copied out of the container's working directory and back into the original directory. This way, users need not fiddle with manual extraction of files.

The `@local` keyword tells Supermake to go back to an uncontainerized shell. All commands use this environment by default.

```Supermake
build: test
	docker build -t myimage .

test:
	@python:3.11 bash
	python manage.py test

	@local
	echo 'All tests passed :)'
```


### Subtargets and parallelisation
One issue with traditional Makefiles is the target graph runs in series. Often leading to situations where they can't be used in CI pipelines as-is, as some steps need parallelism for it to reach acceptable speeds.

Supermake targets can be nested, allowing them to run in parallel automatically. Nested targets can also be used in conjunction with other steps.
```Supermake
build: test
	docker build -t myimage .

test:
	echo 'Starting tests'

	test_a:
		@python:3.11
		python manage.py test a

	test_b:
		@python:3.11
		python manage.py test b

	echo 'All tests passed :)'
```

In this case, test will run all steps normally until it encounters test_a, it then check for other nested targets directly following test_a and runs them all in parallel.

Supermake will wait until all running nested targets have completed before evaluating the next step, or exiting the target.

It is possible to run subtargets by using `supermake test.test_a`. Subtargets can be nested up to three levels deep, though this is discouraged.

### Multi-node pipelines
Targets can be ran on specific nodes via an environment-like directive on directly after the target name, a la email..

```Supermake
build@linux: test
	docker build -t myimage .

test:
	echo 'Starting tests'

	test_a:
		@python:3.11
		python manage.py test a

	test_b:
		@python:3.11
		python manage.py test b

	echo 'All tests passed :)'
```


### Node groups

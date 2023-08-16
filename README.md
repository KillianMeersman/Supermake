# Supermake
## Description
Supermake is a superset of Makefiles that serves as a modern CI pipeline. Supermake can be ran locally as well as on remote CI server(s).

**Features**:
1. Automatic parallelisation.
2. Nested targets.
3. Easy containerization of steps.
4. Compatible with most existing Makefiles.

## User guide
This user guide will guide you through the Supermake featureset by starting with an simple, existing makefile. We'll improve it by adding Supermake features.

Supermake uses the target syntax you're already familiar with, our Makefile starts off as:

```Makefile
build: test
	docker build -t myimage .

test:
	pytest -x tests/
	echo 'All tests passed :)'
```

### Using environments
It can be useful to run certain steps inside a certain shell, interpreter or container. This can be achieved by using the following notation on its own line:

```
[name]@image:[tag] [entrypoint]
```
or
```
[name]@local [entrypoint]
```

All following steps, up until another environment or the end of the target, will be ran in this shell or environment. They can be given a custom name, so as to make it easier to follow logs.

#### Using containers

When a container is started, the working directory is mounted into the container's working directory. This way, you do not need to fiddle with manual extraction of files.

**Notes on containers**
All containers:
- Run with the calling user's UID and GUID.
- Run with host networking mode.
- Are automatically removed after running.
- Use 'sh -ce' as entrypoint unless explicitly provided, irrespective of the image's original entrypoint.

#### Using local shell

The `@local` keyword tells Supermake to go back to an uncontainerized shell, with an optional custom entrypoint. All commands use this environment by default, with "sh -ce" as the def ault entrypoint.

Let's augment our Makefile with environments.
```Makefile
build: test
	docker build -t myimage .

test:
	@python:3 bash
	python manage.py test

	@local
	echo 'All tests passed :)'
```


### Subtargets and parallelisation
One issue with traditional Makefiles is the lack of built-in parallel processing. Often leading to situations where they can't be used in CI pipelines as-is, as some steps need parallelism for it to reach acceptable speeds.

Supermake targets run in parallel whenever possible. They may also be nested to increase readability; Nested targets can also be used in conjunction with other steps.

Let's speed up our testing step by splitting it up into three parallel testing runs. In this example, our end-to-end tests will also wait for both unit and integrations tests to pass first.
```Makefile
build: test
	docker build -t myimage .

test:
	echo 'Starting tests'

	unit:
		@python:3.11
		pytest -x tests/unit/

	integration:
		@python:3.11
		pytest -x tests/integration/

	e2e: unit integration
		@python:3.11
		pytest -x tests/e2e/

	echo 'All tests passed :)'
```

In this example, the `test` target will run steps normally until it encounters `unit`, `integration` and `e2e`, which it will run in parallel.

It is possible to run subtargets by using `supermake test::unit`. Subtargets can be nested as deep as you want, though this is discouraged.

### Distributed pipelines & runner selectors
When using distributed pipelines, targets can be ran on specific runners via an environment-like directive on the target line. It takes the following form:

```
[name]@[node]: [dependencies...]
```

When runners are registered, they are given a name depending on their os & capabilities. Many runners can share the same name, if they share the same OS & capabilities. It is up to the runner provider to establish a naming scheme.

Let's make sure our build step only runs on linux runners with the docker driver installed. In our example, this takes the form of `debian-docker`, which should be fairly self-explanatory.

```Makefile
build@debian-docker: test
	docker build -t myimage .

test:
	echo 'Starting tests'

	unit:
		@python:3.11
		pytest -x tests/unit/

	integration:
		@python:3.11
		pytest -x tests/integration/

	e2e: unit integration
		@python:3.11
		pytest -x tests/e2e/

	echo 'All tests passed :)'
```

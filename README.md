# Supermake
## Description
Supermake is a variant of Make that aims to be a modern CI pipelining tool. Supermake can be ran locally as well as on remote machines.

### Features
1. Automatic parallelisation.
2. Nested targets.
3. Easy containerization of recipes.
4. Easily switch between interpreters for different recipes or even within the same recipe.
5. Compatible with most existing Makefiles (given some easy changes).


## User guide
This user guide will guide you through the Supermake featureset by starting with a simple, existing makefile which we'll augment with Supermake features.

Supermake uses the target syntax you're already familiar with, our Makefile starts off as follows:

```Makefile
IMAGE_TAG ?= latest
IMAGE_URL = "docker.com/myimage:$(IMAGE_TAG)"

.ONESHELL:

build: test
	docker build -t $(IMAGE_URL) .

test:
	pytest -x tests/
	echo 'All tests passed :)'
```

As you can see, we have three targets (.ONESHELL is a "pseudo-target", we'll come back to that) of which two have recipes (the actual scripts under the target names).

### First, a note on Supermake recipes
Unlike the default Make behaviour, which will pass every line of your recipe into a new shell. Supermake passes the whole recipe to the chosen (or default) interpreter. This functionality can be replicated in Make by defining the `.ONESHELL` target. We'll remove this target in the following steps as it's now redundant.

It's for this reason that the default shell in Supermake is `sh -ce`, the `-e` flag makes the recipe fail when a line returns a non-zero exit code, as you would expect in Make. Talking about shells...

### Using different executors
It can be useful to run certain steps of a recipe inside a certain shell, interpreter or container. You can, at any point in a recipe, specify an **executor**. All following recipe lines, up until the next executor line or the end of the recipe, will be passed into said executor.

#### Using a local shell/interpreter
To switch to a different shell or interpreter, with optional arguments (e.g. `-ce`), use the following notation:

```
[name]@local [entrypoint]
```

Note that such executors always begin with `@local`, followed by the entrypoint. Some examples:
- `@local sh -ce`
- `@local python3 -c`
- `@local bash -cex`


#### Using containers

Supermake can also start a container, mount the working directory into it and then run the recipe with the given entrypoint **inside** the container. Use the following notation to do so:

```
[name]@image:[tag] [entrypoint]
```

These executors always begin with the image url, followed by an optional tag (otherwise `latest`) and an optional entrypoint (see chapter above).

Some examples:
- `@python:3 python3 -c`
- `@debian:12 bash -ce`

As stated: When a container is started, the working directory is mounted into the container's working directory. This way, you do not need to fiddle with manual extraction of files.

Please note that all containers:
- Run with the calling user's UID and GUID.
- Use 'host' networking mode.
- Are automatically removed after running.
- Use `sh -ce` as entrypoint unless another entrypoint is explicitly provided, **irrespective of the image's original entrypoint!**

Moving on with our example. Let's augment our Makefile with executors so that our testing step always runs inside a Python container.
```Makefile
IMAGE_TAG ?= latest
IMAGE_URL = "docker.com/myimage:$(IMAGE_TAG)"

build: test
	docker build -t $(IMAGE_URL) .

test:
	@python:3.11 bash -ce
	python manage.py test

	@local
	echo 'All tests passed :)'
```

As an example, we switch back to the local default shell when letting the user know that all tests passed. Had we not put the `@local` line there, the echo command would run inside the Python container as well.

### Subtargets and parallelisation
One issue with traditional Makefiles is the lack of built-in parallel processing. Often leading to situations where they can't be used in CI pipelines as-is because they'd be too slow.

Supermake targets run in parallel whenever possible. They may also be nested to increase readability; Nested targets can also be used in conjunction with other steps.

Let's speed up our testing step by splitting it up into three parallel testing runs. In this example, our end-to-end tests will also wait for both unit and integrations tests to pass first.
```Makefile
IMAGE_TAG ?= latest
IMAGE_URL = "docker.com/myimage:$(IMAGE_TAG)"

build: test
	docker build -t $(IMAGE_URL) .

test:
	echo 'Starting tests'

	unit:
		echo 'Running unit tests'
		@python:3.11 bash -ce
		pytest -x tests/unit/

	integration:
		echo 'Running integration tests'
		@python:3.11 bash -ce
		pytest -x tests/integration/

	e2e: unit integration
		echo 'Running end-to-end tests'
		@python:3.11 bash -ce
		pytest -x tests/e2e/

	echo 'All tests passed :)'
```

In this example, the `test` target will run steps normally until it encounters `unit`, `integration` and `e2e`, which it will run in parallel.

It is possible to run subtargets by using `supermake run test::unit`.

### Distributed pipelines & runner selectors
By default, Supermake runs all targets locally. Ignoring runner directives.

When using distributed pipelines, targets can be ran on specific runners via an environment-like directive on the target line. It takes the following form:

```
[name]@[node]: [dependencies...]
```

When runners are registered, they are given groups depending on their os & capabilities. Many runners can share the same group(s). It is up individual choice on how to establish a naming scheme.

Let's make sure our build step only runs on linux runners with the docker driver installed. In our example, this means the runner in in the `debian-docker` group.

```Makefile
IMAGE_TAG ?= latest
IMAGE_URL = "docker.com/myimage:$(IMAGE_TAG)"

build@debian-docker: test
	docker build -t $(IMAGE_URL) .

test:
	echo 'Starting tests'

	unit:
		echo 'Running unit tests'
		@python:3.11 bash -ce
		pytest -x tests/unit/

	integration:
		echo 'Running integration tests'
		@python:3.11 bash -ce
		pytest -x tests/integration/

	e2e: unit integration
		echo 'Running end-to-end tests'
		@python:3.11 bash -ce
		pytest -x tests/e2e/

	echo 'All tests passed :)'
```

Now, supermake will look for runners with the `debian-docker` tag when running the `build` step. It can still run the `test` step on any runner.

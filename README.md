# Supermake
Supermake is an alternative to Make, the popular build automation tool.
It offers builtin parallelism and container support, all using the a Makefile-like format developers know and love.

> [!Note]  
> Supermake is very much still in-development. It has not yet reached feature-parity. Bugs are to be expected.
> 
> Contributions in the form of PR's, documentation improvements, bug reports and feature requests are all welcome.

## Features 💡
1. Automatic parallelisation of targets whenever possible, increasing speed by utilizing all CPU cores.
2. Nested targets, increasing readability.
3. Run (part of) your recipes in different shells, interpreters and even containers 🐳.
5. Make-like syntax. Re-purpose your Makefiles with a few tweaks.

## Example
```Makefile
TAG ?= latest  # Inherit from environment variable, using 'latest' as fallback
BASE_IMAGE = ghcr.io/killianmeersman/supermake/example/base:${{ TAG }}
CI_IMAGE = ghcr.io/killianmeersman/supermake/example/ci${{ TAG }}

build:  # build will run both the base and ci target
	base:
		docker build --target base -t ${{ BASE_IMAGE }}

	ci: base  # ci will wait until base is done
		docker build --target ci -t ${{ CI_IMAGE }}


test: build::ci  # Wait until the ci subtarget is done, both subtargets will run in parallel
	test_module_a:
		@python:3.11  # Will run inside the python:3.11 image, with the current directory mounted inside
		pytest a/
	
	test_module_b:
		@${{ CI_IMAGE }}  # Will run inside the newly built ci image, with the current directory mounted inside
		pytest b/
```

## Differences with Make ⚠️
- Supermake has no conditional syntax (yet).
- Every Supermake recipe runs in the same shell, making `.ONESHELL` redundant. Make will, by default, run every recipe line in a new shell, .
- Supermake uses `${{ IDENTIFIER }}` variable notation (spaces are optional), instead of Make's `$(IDENTIFIER)` notation.


## Further reading
Check out out `USERGUIDE.md` for a quickstart or `DOCS.md` for a deep-dive into Supermake.

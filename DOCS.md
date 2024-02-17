# Docs
## Variables
Every line is evaluated for variable references, these are replaces ('expanded') before evaluation.
Variables are usually defined at the top of your Supermake file. And have the following syntax:
```
[export ]IDENFITIER (=|:=|?=) VALUE
```



#### export
Prepending the variable with `export` will make it available in all target processes. Otherwise it's only used for expansion.

#### IDENTIFIER
The variable identifier may be composed of any character but cannot include spaces.

#### assignment operator (=|:=|?=)
`=` and `:=` will both assign the provided value to the idenfitier. This will overwrite any value inherited from the parent process.

`?=` is fallback assignment, using the provided value if the variable wasn't inherited from the parent process.

#### VALUE
The value may be any character, except spaces. Variable lines are expanded before being evaluated, allowing you to reference other variables defined before the current line.

### Using variables
You can reference variables on all lines using the `${{ IDENTIFIER }}` syntax. e.g.
```Makefile
MESSAGE = test
export EXPORTED_MESSAGE = ${{ MESSAGE }}-exported

expanded:
	echo ${{ MESSAGE }} # Prints 'test'

not_expanded:
	echo ${MESSAGE} # This will not work, as the MESSAGE variable is not exported, the shell will not interpolate it.

exported:
	echo ${EXPORTED_MESSAGE} # Prints 'test-exported'
```
Be default, referencing a non-existent variable will throw an error.

Variables are recursively expanded, meaning you can construct variable references from other variables. e.g.

```Makefile
VERSION = v1
MESSAGE_v1 = Hello world

recursive_expansion:
    echo ${{MESSAGE_${{VERSION}}}} # Prints 'Hello world'
```

## Targets
Targets are the bread and butter of your Supermake file. They are units of work and can depend on other targets. Dependency targets are automatically scheduled and ran by Supermake. Targets are alway ran only once in a Supermake run.

Targets have the following syntax:
```Makefile
IDENTIFIER[@NODE]: [DEPENDENCIES...]
    (EXECUTOR|RECIPE)...
```

#### IDENTIFIER
The target idenfifier may consist of any character, except space.

#### @NODE
Optionally, you may specify a node selector to run this target on (more on that later).

#### DEPENDENCIES
Zero or more target names that have to run before this target can run. separated by spaces.

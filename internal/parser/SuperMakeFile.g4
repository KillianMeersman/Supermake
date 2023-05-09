grammar SuperMakeFile;

options {
	language = Go;
}

// Lexer rules
INDENT: '\t';
SPACES: ' '+ -> skip; // Skip whitespace
NEWLINE: [\r\n]+; // End-of-line character(s)

VAR: '$' [A-Za-z_][A-Za-z0-9_]*; // Variables (e.g. $FOO)
EXECUTOR: '@' [A-Za-z0-9_:]+ [A-Za-z0-9_]+;
WORD:
	~[ \t\r\n$#]+; // Words (e.g. target name, dependencies, etc.)

COMMENT: '#' ~[\r\n]*; // Comments

fragment DIGIT: [0-9];

// Parser rules
supermakefile: (line | NEWLINE)*;

line: (recipe | declaration) NEWLINE;

declaration: variable | target;

variable: VAR '=' WORD;

target: WORD ':' dependencies? NEWLINE recipe;

dependencies: WORD (SPACES WORD)*;

recipe: (INDENT (executor_line | command_line | INDENT target))*;

executor_line: EXECUTOR NEWLINE;
command_line: (WORD+ SPACES*)+ NEWLINE;

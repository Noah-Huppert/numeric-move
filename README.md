# Numeric Move
Tool for moving files with numeric prefixes.

# Table Of Contents
- [Overview](#overview)
- [Usage](#usage)
- [Build](#build)

# Overview
Tool which helps move files who's names are prefixed with numbers.

*Example of numerically prefixed files:*  

```
# ls
001-foo
002-bazz
003-bar
```

Renaming files with this pattern can be painful if you want to preserve the
numeric ordering.

Numeric move helps move these types of files while preserving the numeric 
order in the file names.

# Usage
The tool takes the current numeric prefix and the new numeric prefix 
as arguments:

```
# nmv FROM TO
```

The tool will rename all files with this prefix with a new prefix.  

If a files exists with the `TO` prefix those file's prefixes will be 
incremented until they no longer conflict. If in this process more conflicts 
occur the same strategy will be used.

## Simple Move
Using the simple example from above:

```
# ls
001-foo
002-bazz
003-bar
# nmv 001 002
# ls
002-foo
003-bazz
004-bar
```

## Difference Move
The numeric move tool can also ensure that there is a numeric difference 
between each prefix using the `--diff,-d` option:

```
# ls
001-foo
010-bazz
020-bar
# nmv 001 010 -d 10
# ls
010-foo
020-bazz
030-bar
```

## Resize Move
If the `TO` the numeric prefix cannot be expressed in the number of digits 
which files currently have the tool will fail. The `--resize,-r` option can be 
used to force the tool to add additional digits to all files numeric prefixes.

```
# ls
001-foo
002-bazz
003-bar
# nmv 001 1000 -r
# ls
1000-foo
1001-bazz
1002-bar
```

# Build
Build with Make:

```
make
```

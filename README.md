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
```
nmv FROM TO
```

This will rename a file with the numeric prefix `FROM` to have numeric prefix 
of `TO`.

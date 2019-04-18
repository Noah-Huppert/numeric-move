# Design
Numeric move design.

# Table Of Contents
- [Overview](#overview)
- [Spaced List](#spaced-list)

# Overview
Numeric move keeps the ordering of files intact when the target file is moved.  

It tried to move as few files as possible.  

Only files after the new target file's position are moved.

# Spaced List
Spaced list is a special data structure used by the numeric move tool.  

It stores names of files and the spaces between file's numeric prefixes.  
This allows files to be easily moved around and have their numeric 
prefixes recalculated.

## Example
Given the files:

```
00-aaa
05-bbb
10-ccc
```

The following spaced list exists:

```
key:
(file name)
[space size]
```
```
(aaa) [4] (bbb) [4] (ccc)
```

Inserting the file `07-iii` gives us:

```
(aaa) [4] (bbb) [1] (iii) [2] (ccc)
```

Inserting duplicate prefixed files pushes existing files back.  
If we insert `05-ddd`:

```
(aaa) [4] (ddd) (bbb) (iii) [2] (ccc)
```

The list can be converted back into prefixed file names:

```
00-aaa
05-ddd
06-bbb
07-iii
10-ccc
```

## Operations
### Build
`Build([]Files)`

Arguments:

- `Files`: List of file structures order by their prefix field in 
  ascending order
  
Builds a spaced list. Inserts files in their correct positions and adds spaces 
when required.  

The difference between the build and insert operation is that build ensures the 
spaced list represents the provided files exactly. While insert takes the 
liberty of shuffling files around, which may result in different prefixes than 
they had originally.

### Insert
`Insert(Prefix, Name)`  

Arguments:

- `Prefix`: Numeric prefix
- `Name`: File name

Insert file at position in list which will give file the provided prefix.  
Attempts to remove 1 unit from the closest space after the inserted file.

### Get Prefixes
`GetPrefixes() map[string]number`

Get the prefixes of each file as represented by the spaced list.

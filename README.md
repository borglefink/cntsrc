[![MIT License](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)
[![works badge](https://cdn.rawgit.com/nikku/works-on-my-machine/v0.2.0/badge.svg)](https://github.com/nikku/works-on-my-machine)

## Description

*cntsrc* is a small command line utility for counting source code lines. 
It can also count binaries (number of files and filesize only).
A config file configures what files to count (see config section below).
When counting source code lines, newline will be determined as type windows (CRLF) or unix/osx (LF) for each file.

The result will look along the lines of this:
```
Directory processed:
c:\projectdirectory
---------------------------------------------------------------
filetype        #files       #lines  line%          size  size%
---------------------------------------------------------------
.css                 9        3 512   42.3        92 168   12.0
.html                1          229    2.8         7 626    1.0
.js                 22        4 563   54.9       195 256   25.5
.jpg                 7                           260 274   33.9
.png               120                           211 318   27.6
---------------------------------------------------------------
Total:             159        8 304  100.0       766 642  100.0
```

## Usage

Give a directory as a parameter. If none is given, the current directory is used.
All sub-directories will be searched as well, and included in the result.

```
countsource [directory] [-c fullpathtoconfigfile] [-debug]
```

The optional parameter *-debug* is for analysis/debug, showing which directories and files are included or excluded.

Use *cntsrc -?* to show usage.

## Config file

The configuration file can be specified with *-c "full config file name"*. 

If no config file is specified, it is read from the current directory. 
If not found in the current directory, the config is expected to be found in the same directory as the executable. 

If a config file does not exist, one is created in the current directory,
with default values similar to the following:

```JSON
/*
 * Config file for example project
 * For binary files, only file size is counted
 */
{
  "Extensions": [
    ".css",
    ".html",
    ".js",
    ".jpg",
    ".png"
  ],
  "Exclusions": [
    "bin",
    "obj",
    "Scripts\\jquery.ui",
    "cntsrc.config",
    "jquery.min.js"
  ]
}
```

The full file path is searched for exclusion string. 
Directories can be qualified/separated with the standard operating system path separator i.e. *\* on windows, */* on unix/mac. 
Note that the windows path separator *\* needs to be escaped as *\\*.

It is possible to put comments in the config file. 
Note that comments are normally not allowed in json, so these comments are stripped from the config file before it is read.
Only Go-type comments are allowed, single line comments starting with //, or block comments enclosed by /* and */.

## Install

Clone the repository into your GOPATH somewhere and resolve the only third party dependency,
then do a **go build** or a **go install**.

Create a config file for a project you want to count source code for, and put the config file in the root of that directory.
If you have several projects using identical config files, use a single config file and refer to it with the *-c* parameter when counting.

_cntsrc_ is dependent upon Michael T Jones' fast parallel filesystem traversal package. See [github.com/MichaelTJones/walk](https://github.com/MichaelTJones/walk).


## Background

I wanted to count the number of source code lines for all the source code in an ASP.NET MVC project to keep track of the size of it. So I wrote this. 
This is the second implementation of this utility.
The first version can be found at [github.com/borglefink/countsource](https://github.com/borglefink/countsource).

## License

A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)

## Description

*cntsrc* is a command line utility for counting source code lines. It can also count binaries (number of files and filesize only). A config file configures what files to count (see config section below). When counting source code lines, newline will be determined as type windows (CRLF) or unix/osx (LF) for each file.

The result will look along the lines of this:
```
Directory processed:
/home/myuser/projectdirectory
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

Give a directory as a parameter. If none is given, the current directory is used. All sub-directories will be searched as well, and included in the result.

```
cntsrc [-c pathtoconfigfile] [-big n] [-debug] [directory] 
```

The optional parameter *-big n* can be added to get the *n* largest source code files, in terms of source code lines.

The optional parameter *-debug* is for analysis/debug, showing which files are included or excluded. This parameter makes it easier to set up the config file properly, but is normally never used after the config file is set up properly.

Use *cntsrc -h* to show usage.

## Config file

The configuration file can be specified with *-c "full config file name"*. 

If no config file is specified, it is read from the current directory. If not found in the current directory, the config is expected to be found in the same directory as the executable. 

If a config file does not exist, one is created in the current directory, with default values similar to the following:

```JSON
/*
 * Config file for example project
 * For binary files, only file size is counted
 */
{
  "FileExtensions": [
    ".css",
    ".html",
    ".js",
    ".jpg",
    ".png"
  ],
  "Exclusions": [
    ".git",
    "\\bin",
    "\\node_modules",
    "Scripts\\jquery",
    "cntsrc.config"
  ]
}
```

Create a config file for a project you want to count source code for, and put the config file in the root of that project directory. If you have several projects using identical config files, use a single config file and refer to it with the *-c* parameter when counting.

When traversing the file system, each file system entry is examined, and will be excluded if the file- or directoryname (including path) contains one of the exclusion strings. Directories can be qualified/separated with the standard operating system path separator i.e. \ on windows, / on linux and osx. Note that the windows path separator \ needs to be escaped  as \\\\ inside a json string.

It is possible to put comments in the config file. Note that comments are normally not allowed in json, these comments are stripped from the config file before it is interpreted as json. Only Go-type comments are allowed, single line comments starting with //, or block comments enclosed by /\* and \*/.

## Install

Clone the repository into your GOPATH somewhere and resolve dependencies (see below), then do a **go install**.

## Dependencies

_cntsrc_ is dependent upon Michael T Jones' fast parallel filesystem traversal package. 
See [github.com/MichaelTJones/walk](https://github.com/MichaelTJones/walk). 

## Background

I wanted to count the number of source code lines for source code in an ASP.NET MVC project, to keep track of the size of it. So I wrote _cntsrc_. This is a re-implemetation of the original utility, found at [github.com/borglefink/countsource](https://github.com/borglefink/countsource).

## License

A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)

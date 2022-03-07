# NUMCAT

This utility is similar to the venerable `cat(1)` program, in that it writes the output of a file to stdout.

The idea is to concatenate files, you can type something like this:

```text
# Concatenate the files a, b, and c into a new file, d
$ cat a b c > d
```
In practice, it can also be used to view the contents of a file. This utility just 
prepends the line number onto each line.

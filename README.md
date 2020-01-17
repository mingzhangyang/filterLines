# filterLines
A utility to do the similar job "grep -f file1 file2"

Build

- git clone to download the repo
- cd to the directory
- run `go build -o filterLines main.go filter.go`

How to use the utility

- `./filterLines -h` to see the help info
- `./filterLines file1 file2` to filter file2 using lines in file1. By default, each line in file1 is taken as pattern; each line in the file2 will be split by `\t`, the first field is used to check whether matches patterns; if matched then keep this line, otherwise ignore this line.
- `./filterLines -r file1 file2`, do the similar work to above. But it will keep the lines in file2 that don't match to the patterns in file1; matched will be ignored.
- `./filterLines -f 3 file1 file2`, specify that the third field in each line of file2 will be used to match patterns
- `./filterLines -d "," file1 file2`, specify that `,` should be used to split lines of file2

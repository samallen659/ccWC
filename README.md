# ccWC

This is my implementation of the wc project from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc)

## Building

To build ccWC use the following command

```bash
go build -o ccWC ./main.go
```

### Usage

ccWC is made to follow the basic usage of wc
  -c    Show byte count of input
  -l    Show line count of input
  -m    Show character count of input
  -w    Show word count of input

File can be read by providing the file name

```bash
./ccWC file.txt
```

Or through stdin

```bash
cat file.txt | ./ccWC
```

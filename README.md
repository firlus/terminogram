# Terminogram - Nonogram over Telnet
[This is how Terminogram looks](firl.us/terminogram-screenshot)

## Usage
Connect via telnet to a Terminogram server. Start your own by running:
```
go run cmd/nononet/nononet.go
```
The standard port is ```42002```.

You also can use my hosted version:
```
telnet terminogram.firlus.dev
```
I use the Telnet standard port 23, so you do not have to explicitly add it.

If you do not know how Nonogram works, ask Google for the rules.

Control your cursor with the arrow keys. Mark a square by pressing Space.
If you miss 3 times you have lost.

## Current state
This is the bare minimum to get it running. Code quality is meh, there is still room for improvement.
There is exactly one Nonogram puzzle in this application.

Here's a list of features I will try to implement if I have time:
* Rename the go module to Terminogram (Nononet was the first attempt of finding a name and it is terrible...)
* More than one puzzle (selected randomly, read from file)
* Mark rows and columns where nothing more is to be added green
* Mark error crosses red
* Explain usage in the app
* Make it possible to play more than one round

## Contribute
Make pull requests that add puzzles. Just a 10x10 grid of 0s and 1s in a file. If you'd like to add features, start an issue, you are welcome to make improvements to this toy. 

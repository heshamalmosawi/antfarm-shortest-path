# antfarm-shortest-path
## Description

This project is a project under-construction (incomplete as of now) simulates a digital version of an ant farm. Its main objective is to create a program that reads the input from a file that states the number of ants, and the structure for its colony (the tunnels and rooms in the ant colony). Then the program would process it and find the quickest path(s) for the ants to traverse the colony. The output would be each ant's move from room to room until they reach the ending room.

## How it Works
- An ant farm is created with tunnels and rooms.
- The ants start at the room labeled ##start and aim to reach the room labeled ##end with the fewest moves possible.
- The program must handle various invalid or poorly-formatted input scenarios and display appropriate error messages.
- Results are displayed on the standard output/terminal in.

## Instructions
- Create rooms and tunnels.
- A room's name must not start with letter 'L' or '#'.
- Rooms are connected by tunnels, Each room can only contain one ant at a time. (Except start and end)

## Usage
```
go run main.go <file_name>
```

## Input file format
- tbc

## Contributors
- Sayed Hesham Husain
- Ali Ebrahim
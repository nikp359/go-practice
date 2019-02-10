# Ghost Legs
[Ссылка на codingame](https://www.codingame.com/ide/puzzle/ghost-legs)

Input
Line 1: Integer W and H for width and height of the diagram below.
Next H lines: Containing a Ghost Legs diagram as your input.

The diagram itself is composed of characters: '|' and '-', (and space).
The top line in the diagram has a number of labels T.
The bottom line contains labels B.

Each T and B is a single ascii character that can be of any random value. Do not assume they will always be ABC or 123.

As a rule of the game, left and right horizontal connectors will never appear at the same point.
Output
List all connected pairs between top and bottom labels, TB, in the order of the top labels from Left to Right. Write each pair in a separate line.

Constraints
3 < W, H ≤ 100

Example

Input
7 7
A  B  C
|  |  |
|--|  |
|  |--|
|  |--|
|  |  |
1  2  3

Output
A2
B1
C3
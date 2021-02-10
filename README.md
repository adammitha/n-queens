# n-queens

Simple command line tool for solving the n-queens problem for arbitrary-sized boards.

## Installation instructions
1. Fetch the source code
`go get github.com/adammitha/n-queens`
2. `cd` into the root directory of the repository
3. Install the binary into your gobin
`go install .`

## How to use
`n-queens <size of board>`
E.g. for an 8 x 8 board: 
```
>n-queens 8
>  Q  ▢  ▢  ▢  ▢  ▢  ▢  ▢ 
   ▢  ▢  ▢  ▢  Q  ▢  ▢  ▢ 
   ▢  ▢  ▢  ▢  ▢  ▢  ▢  Q 
   ▢  ▢  ▢  ▢  ▢  Q  ▢  ▢ 
   ▢  ▢  Q  ▢  ▢  ▢  ▢  ▢ 
   ▢  ▢  ▢  ▢  ▢  ▢  Q  ▢ 
   ▢  Q  ▢  ▢  ▢  ▢  ▢  ▢ 
   ▢  ▢  ▢  Q  ▢  ▢  ▢  ▢ 

```
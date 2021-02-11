# N-queens solver

## About
Simple command line tool for solving the n-queens problem for arbitrary-sized boards. This repository is a companion for an upcoming blog post (link to follow).

## Installation
### Option 1: Prebuild binaries
1. Prebuilt binaries are available from the `Releases` page for GOOS=Linux, Windows, Darwin.

### Option 2: Build from source
1. Fetch the source code
`go get github.com/adammitha/n-queens`
2. `cd` into the root directory of the repository
3. Install the binary into your gobin
`go install .`

## Usage
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

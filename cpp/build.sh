#!/bin/bash
clang++ -std=c++20 -o a.out $1

# Error: Requires Linker Symbols (ld)
# $(brew --prefix gcc)/bin/gcc-14 $1

# For Testing
cmake -S . -B build
cmake --build build
./build/runTests

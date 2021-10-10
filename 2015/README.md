# Advent of Code 2015

https://adventofcode.com/2015

## Language

Implemented in [C++11](https://en.cppreference.com/w/cpp/11) with
[CMake](https://cmake.org/) and
[GoogleTest](https://google.github.io/googletest/).

## Preparation

1. Install [CMake](https://cmake.org/).
2. Install a compiler (GCC, Clang, MSVC).
3. Initialize the build system:

   ```console
   cmake -B build
   ```

## Building

Use `cmake` to build the project:

```console
cmake --build build/
```

## Running tests

Run test test using `ctest`:

```console
cd build/
ctest
```

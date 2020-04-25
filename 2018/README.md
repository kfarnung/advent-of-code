# Advent of Code 2018

https://adventofcode.com/2018

## Preparation

These solutions were originally written in Python 2, but updated to support both Python 2 and Python 3. The tests are written to use `pytest`, which needs to be installed:

```console
python -m pip install pytest
```

## Running tests

The tests can simply be run using `pytest` with the Python version of choice:

```console
python -m pytest
```

## Execution time

Some solutions do take a long time to execute (on the order of 20 minutes) using
the standard `CPython` distribution. The use of [PyPy](https://www.pypy.org/)
drastically improves the execution time.

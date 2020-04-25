# Advent of Code 2018

https://adventofcode.com/2018

## Language

This year I chose to implement in Python 2, but ported the code such that it
works in either Python 2 or Python 3.

## Preparation

1. Install your favorite version of [Python](https://www.python.org/).
2. The tests are written to use `pytest`, which needs to be installed:

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

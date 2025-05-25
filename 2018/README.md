# Advent of Code 2018

https://adventofcode.com/2018

## Language

This year I chose to implement in Python 2, but ported the code such that it
works properly in Python 3. Python 2 support has been removed.

## Preparation

Install [Python 3](https://www.python.org/).

## Running tests

The tests can be run using `pytest` within a Python 3 venv:

```console
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
pytest
```

## Execution time

Some solutions do take a long time to execute (on the order of 20 minutes) using
the standard `CPython` distribution. The use of [PyPy](https://www.pypy.org/)
drastically improves the execution time.

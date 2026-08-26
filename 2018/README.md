# Advent of Code 2018

https://adventofcode.com/2018

## Language

This year I chose to implement in Python 2, but ported the code such that it
works properly in Python 3. Python 2 support has been removed.

## Preparation

Install [Python 3](https://www.python.org/) and [uv](https://docs.astral.sh/uv/).

## Running tests

The tests can be run using `uv`:

```console
uv run pytest -n auto
```

To run including slow tests:

```console
uv run pytest -n auto --slow
```

## Linting

```console
uv run ruff check .
```

To format the code:

```console
uv run ruff format .
```

## Execution time

Some solutions do take a long time to execute (on the order of 20 minutes) using
the standard `CPython` distribution. The use of [PyPy](https://www.pypy.org/)
drastically improves the execution time.

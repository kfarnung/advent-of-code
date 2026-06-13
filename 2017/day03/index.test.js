
const Day03 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 03', () => {
  describe('Part 1', () => {
    test('Part 1', () => {
      assert.strictEqual(Day03.manhattanDistance(1), 0);
      assert.strictEqual(Day03.manhattanDistance(12), 3);
      assert.strictEqual(Day03.manhattanDistance(23), 2);
      assert.strictEqual(Day03.manhattanDistance(1024), 31);
    });

    test('Part 2', () => {
      assert.strictEqual(Day03.searchSums(0), 1);
      assert.strictEqual(Day03.searchSums(1), 2);
      assert.strictEqual(Day03.searchSums(3), 4);
      assert.strictEqual(Day03.searchSums(24), 25);
      assert.strictEqual(Day03.searchSums(350), 351);
    });
  });

  test('Input', () => {
    const [part1, part2] = Day03.run('../private/inputs/2017/day03.txt');
    assert.strictEqual(part1, 430);
    assert.strictEqual(part2, 312453);
  });
});

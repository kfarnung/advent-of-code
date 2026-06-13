
const Day11 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 11', () => {
  test('Input', () => {
    const [part1, part2] = Day11.run('../private/inputs/2017/day11.txt');
    assert.strictEqual(part1, 696);
    assert.strictEqual(part2, 1461);
  });
});

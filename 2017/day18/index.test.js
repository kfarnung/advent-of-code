
const Day18 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 18', () => {
  test('Input', () => {
    const [part1, part2] = Day18.run('../private/inputs/2017/day18.txt');
    assert.strictEqual(part1, 2951);
    assert.strictEqual(part2, 7366);
  });
});

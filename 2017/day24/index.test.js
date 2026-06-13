
const Day24 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 24', () => {
  test('Input', () => {
    const [part1, part2] = Day24.run('../private/inputs/2017/day24.txt');
    assert.strictEqual(part1, 1868);
    assert.strictEqual(part2, 1841);
  });
});

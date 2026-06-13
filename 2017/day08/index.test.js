
const Day08 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 08', () => {
  test('Input', () => {
    const [part1, part2] = Day08.run('../private/inputs/2017/day08.txt');
    assert.strictEqual(part1, 5966);
    assert.strictEqual(part2, 6347);
  });
});

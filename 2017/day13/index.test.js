
const Day13 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 13', () => {
  test('Input', () => {
    const [part1, part2] = Day13.run('../private/inputs/2017/day13.txt');
    assert.strictEqual(part1, 1640);
    assert.strictEqual(part2, 3960702);
  });
});


const Day06 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 06', () => {
  test('Input', () => {
    const [part1, part2] = Day06.run('../private/inputs/2017/day06.txt');
    assert.strictEqual(part1, 7864);
    assert.strictEqual(part2, 1695);
  });
});

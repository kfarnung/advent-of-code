
const Day12 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 12', () => {
  test('Input', () => {
    const [part1, part2] = Day12.run('../private/inputs/2017/day12.txt');
    assert.strictEqual(part1, 169);
    assert.strictEqual(part2, 179);
  });
});

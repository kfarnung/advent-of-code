
const Day22 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 22', () => {
  test('Input', () => {
    const [part1, part2] = Day22.run('../private/inputs/2017/day22.txt');
    assert.strictEqual(part1, 5261);
    assert.strictEqual(part2, 2511927);
  });
});

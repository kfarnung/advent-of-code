
const Day02 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 02', () => {
  test('Input', () => {
    const [part1, part2] = Day02.run('../private/inputs/2017/day02.txt');
    assert.strictEqual(part1, 42378);
    assert.strictEqual(part2, 246);
  });
});

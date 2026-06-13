
const Day19 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 19', () => {
  test('Input', () => {
    const [part1, part2] = Day19.run('../private/inputs/2017/day19.txt');
    assert.strictEqual(part1, 'RUEDAHWKSM');
    assert.strictEqual(part2, 17264);
  });
});

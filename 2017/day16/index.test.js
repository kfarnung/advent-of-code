const Day16 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 16', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day16.run('../private/inputs/2017/day16.txt');
      assert.strictEqual(part1, 'padheomkgjfnblic');
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day16.run('../private/inputs/2017/day16.txt');
      assert.strictEqual(part2, 'bfcdeakhijmlgopn');
    });
  });
});

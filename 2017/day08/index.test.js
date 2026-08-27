import Day08 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 08', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day08.run('../private/inputs/2017/day08.txt');
      assert.strictEqual(part1, 5966);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day08.run('../private/inputs/2017/day08.txt');
      assert.strictEqual(part2, 6347);
    });
  });
});

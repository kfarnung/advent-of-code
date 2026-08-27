import Day13 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 13', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day13.run('../private/inputs/2017/day13.txt');
      assert.strictEqual(part1, 1640);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day13.run('../private/inputs/2017/day13.txt');
      assert.strictEqual(part2, 3960702);
    });
  });
});

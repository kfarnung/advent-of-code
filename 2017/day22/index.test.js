import Day22 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 22', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day22.run('../private/inputs/2017/day22.txt');
      assert.strictEqual(part1, 5261);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day22.run('../private/inputs/2017/day22.txt');
      assert.strictEqual(part2, 2511927);
    });
  });
});

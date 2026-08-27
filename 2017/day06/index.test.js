import Day06 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 06', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day06.run('../private/inputs/2017/day06.txt');
      assert.strictEqual(part1, 7864);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day06.run('../private/inputs/2017/day06.txt');
      assert.strictEqual(part2, 1695);
    });
  });
});

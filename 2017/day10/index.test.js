import Day10 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 10', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day10.run('../private/inputs/2017/day10.txt');
      assert.strictEqual(part1, 5577);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day10.run('../private/inputs/2017/day10.txt');
      assert.strictEqual(part2, '44f4befb0f303c0bafd085f97741d51d');
    });
  });
});

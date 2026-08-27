import Day19 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 19', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day19.run('../private/inputs/2017/day19.txt');
      assert.strictEqual(part1, 'RUEDAHWKSM');
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day19.run('../private/inputs/2017/day19.txt');
      assert.strictEqual(part2, 17264);
    });
  });
});

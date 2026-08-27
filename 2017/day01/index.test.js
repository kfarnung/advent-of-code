import Day01 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 01', () => {
  describe('Part 1', () => {
    test('Examples', () => {
      assert.strictEqual(Day01.sumString('1122'), 3);
      assert.strictEqual(Day01.sumString('1111'), 4);
      assert.strictEqual(Day01.sumString('1234'), 0);
      assert.strictEqual(Day01.sumString('91212129'), 9);
    });

    test('Input', () => {
      const [part1] = Day01.run('../private/inputs/2017/day01.txt');
      assert.strictEqual(part1, 1102);
    });
  });

  describe('Part 2', () => {
    test('Examples', () => {
      assert.strictEqual(Day01.sumStringHalfway('1212'), 6);
      assert.strictEqual(Day01.sumStringHalfway('1221'), 0);
      assert.strictEqual(Day01.sumStringHalfway('123425'), 4);
      assert.strictEqual(Day01.sumStringHalfway('123123'), 12);
      assert.strictEqual(Day01.sumStringHalfway('12131415'), 4);
    });

    test('Input', () => {
      const [, part2] = Day01.run('../private/inputs/2017/day01.txt');
      assert.strictEqual(part2, 1076);
    });
  });
});


const Day01 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 01', () => {
  describe('Examples', () => {
    test('Part 1', () => {
      assert.strictEqual(Day01.sumString('1122'), 3);
      assert.strictEqual(Day01.sumString('1111'), 4);
      assert.strictEqual(Day01.sumString('1234'), 0);
      assert.strictEqual(Day01.sumString('91212129'), 9);
    });

    test('Part 2', () => {
      assert.strictEqual(Day01.sumStringHalfway('1212'), 6);
      assert.strictEqual(Day01.sumStringHalfway('1221'), 0);
      assert.strictEqual(Day01.sumStringHalfway('123425'), 4);
      assert.strictEqual(Day01.sumStringHalfway('123123'), 12);
      assert.strictEqual(Day01.sumStringHalfway('12131415'), 4);
    });
  });

  test('Input', () => {
    const [part1, part2] = Day01.run('../private/inputs/2017/day01.txt');
    assert.strictEqual(part1, 1102);
    assert.strictEqual(part2, 1076);
  });
});

/* global describe, expect, test */

const Day01 = require('.');

describe('Examples', () => {
  test('Part 1', () => {
    expect(Day01.sumString('1122')).toBe(3);
    expect(Day01.sumString('1111')).toBe(4);
    expect(Day01.sumString('1234')).toBe(0);
    expect(Day01.sumString('91212129')).toBe(9);
  });

  test('Part 2', () => {
    expect(Day01.sumStringHalfway('1212')).toBe(6);
    expect(Day01.sumStringHalfway('1221')).toBe(0);
    expect(Day01.sumStringHalfway('123425')).toBe(4);
    expect(Day01.sumStringHalfway('123123')).toBe(12);
    expect(Day01.sumStringHalfway('12131415')).toBe(4);
  });
});

test('Input', () => {
  const [part1, part2] = Day01.run('../private/inputs/2017/day01.txt');
  expect(part1).toBe(1102);
  expect(part2).toBe(1076);
});

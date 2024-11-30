/* global expect, test */

const Day09 = require('.');

test('Input', () => {
  const [part1, part2] = Day09.run('../private/inputs/2017/day09.txt');
  expect(part1).toBe(16021);
  expect(part2).toBe(7685);
});

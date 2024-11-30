/* global expect, test */

const Day14 = require('.');

test('Input', () => {
  const [part1, part2] = Day14.run('../private/inputs/2017/day14.txt');
  expect(part1).toBe(8292);
  expect(part2).toBe(1069);
});

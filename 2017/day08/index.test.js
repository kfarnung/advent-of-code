/* global expect, test */

const Day08 = require('.');

test('Input', () => {
  const [part1, part2] = Day08.run('../private/inputs/2017/day08.txt');
  expect(part1).toBe(5966);
  expect(part2).toBe(6347);
});

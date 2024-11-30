/* global expect, test */

const Day24 = require('.');

test('Input', () => {
  const [part1, part2] = Day24.run('../private/inputs/2017/day24.txt');
  expect(part1).toBe(1868);
  expect(part2).toBe(1841);
});

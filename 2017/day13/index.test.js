/* global expect, test */

const Day13 = require('.');

test('Input', () => {
  const [part1, part2] = Day13.run('../private/inputs/2017/day13.txt');
  expect(part1).toBe(1640);
  expect(part2).toBe(3960702);
});

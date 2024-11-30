/* global expect, test */

const Day17 = require('.');

test('Input', () => {
  const [part1, part2] = Day17.run('../private/inputs/2017/day17.txt');
  expect(part1).toBe(419);
  expect(part2).toBe(46038988);
});

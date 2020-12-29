/* global expect, test */

const Day17 = require('.');

test('Input', () => {
  const [part1, part2] = Day17.run('./day17/input');
  expect(part1).toBe(419);
  expect(part2).toBe(46038988);
});

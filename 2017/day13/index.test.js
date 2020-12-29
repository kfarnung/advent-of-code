/* global expect, test */

const Day13 = require('.');

test('Input', () => {
  const [part1, part2] = Day13.run('./day13/input');
  expect(part1).toBe(1640);
  expect(part2).toBe(3960702);
});

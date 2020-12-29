/* global expect, test */

const Day02 = require('.');

test('Input', () => {
  const [part1, part2] = Day02.run('./day02/input');
  expect(part1).toBe(42378);
  expect(part2).toBe(246);
});

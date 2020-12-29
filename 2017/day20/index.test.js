/* global expect, test */

const Day20 = require('.');

test('Input', () => {
  const [part1, part2] = Day20.run('./day20/input');
  expect(part1).toBe(300);
  expect(part2).toBe(502);
});

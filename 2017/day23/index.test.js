/* global expect, test */

const Day23 = require('.');

test('Input', () => {
  const [part1, part2] = Day23.run('../private/inputs/2017/day23.txt');
  expect(part1).toBe(4225);
  expect(part2).toBe(905);
});

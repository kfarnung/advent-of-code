/* global expect, test */

const Day21 = require('.');

test('Input', () => {
  const [part1, part2] = Day21.run('../private/inputs/2017/day21.txt');
  expect(part1).toBe(197);
  expect(part2).toBe(3081737);
});

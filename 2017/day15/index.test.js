/* global expect, test */

const Day15 = require('.');

test('Input', () => {
  const [part1, part2] = Day15.run('../private/inputs/2017/day15.txt');
  expect(part1).toBe(638);
  expect(part2).toBe(343);
});

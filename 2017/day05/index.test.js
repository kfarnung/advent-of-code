/* global expect, test */

const Day05 = require('.');

test('Input', () => {
  const [part1, part2] = Day05.run('../private/inputs/2017/day05.txt');
  expect(part1).toBe(339351);
  expect(part2).toBe(24315397);
});

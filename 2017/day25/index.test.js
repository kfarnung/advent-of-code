/* global expect, test */

const Day25 = require('.');

test('Input', () => {
  const [part1] = Day25.run('./day25/input');
  expect(part1).toBe(2832);
});

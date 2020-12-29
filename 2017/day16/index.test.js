/* global expect, test */

const Day16 = require('.');

test('Input', () => {
  const [part1, part2] = Day16.run('./day16/input');
  expect(part1).toBe('padheomkgjfnblic');
  expect(part2).toBe('bfcdeakhijmlgopn');
});

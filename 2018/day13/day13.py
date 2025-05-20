"""
Implementation for Advent of Code Day 13.

https://adventofcode.com/2018/day/13
"""

_TURN_ORDER = [
    -1,
    0,
    1
]

_CART_DIRECTIONS = [
    '^', '>', 'v', '<',
]


class Cart:
    """Represents a cart on the track."""

    def __init__(self, row_index, cell_index, direction):
        self.row_index = row_index
        self.cell_index = cell_index
        self.direction = _CART_DIRECTIONS.index(direction)
        self.intersection_count = 0
        self.did_collide = False

    def get_position(self):
        """Gets the cart's current position."""
        return self.row_index, self.cell_index

    def get_track_replacement(self):
        """Gets the track that's under the cart (only works at start)."""
        if self.direction % 2 == 0:
            return '|'

        return '-'

    def update_direction(self, next_track):
        """Update the direction of the cart."""
        if next_track == '+':
            self._turn(self._get_next_turn())
        elif next_track == '/':
            if self.direction % 2 == 0:
                self._turn(1)
            elif self.direction % 2 == 1:
                self._turn(-1)
        elif next_track == '\\':
            if self.direction % 2 == 0:
                self._turn(-1)
            else:
                self._turn(1)
        elif next_track in ('-', '|'):
            # Nothing to do
            pass
        else:
            raise ValueError('Unexpected track')

    def move_once(self):
        """Update the position of the cart."""
        if self.direction == 0:
            self.row_index -= 1
        elif self.direction == 1:
            self.cell_index += 1
        elif self.direction == 2:
            self.row_index += 1
        elif self.direction == 3:
            self.cell_index -= 1
        else:
            raise ValueError('Invalid direction')

    def _turn(self, delta):
        self.direction = (self.direction + delta) % len(_CART_DIRECTIONS)

    def _get_next_turn(self):
        turn = _TURN_ORDER[self.intersection_count]
        self.intersection_count = (
            self.intersection_count + 1) % len(_TURN_ORDER)
        return turn


class CartTrack:
    """Represents the current state of the cart track."""

    def __init__(self, content):
        self.grid = [list(cell for cell in row) for row in content]
        self.carts = []

        for row_index, row in enumerate(self.grid):
            for cell_index, cell in enumerate(row):
                cart = CartTrack._try_get_cart(row_index, cell_index, cell)
                if cart:
                    row[cell_index] = cart.get_track_replacement()
                    self.carts.append(cart)

    def __str__(self):
        return ''.join(''.join(row) for row in self.grid)

    def find_first_crash(self):
        """Finds the location where the first two carts crash."""
        while True:
            for cart in sorted(self.carts, key=Cart.get_position):
                cart.move_once()
                cart.update_direction(
                    self.grid[cart.row_index][cart.cell_index])

                if self._mark_collisions():
                    return cart.cell_index, cart.row_index

    def find_last_cart(self):
        """Finds the location of the last cart remaining on the track."""
        while len(self.carts) > 1:
            for cart in sorted(self.carts, key=Cart.get_position):
                if cart.did_collide:
                    continue

                cart.move_once()
                cart.update_direction(
                    self.grid[cart.row_index][cart.cell_index])
                self._mark_collisions()

            self.carts = [cart for cart in self.carts if not cart.did_collide]

        return self.carts[0].cell_index, self.carts[0].row_index

    def _mark_collisions(self):
        found_collision = False

        for cart in self.carts:
            if cart.did_collide:
                continue

            for cart2 in self.carts:
                if (not cart2.did_collide and cart != cart2 and
                        cart.get_position() == cart2.get_position()):
                    cart.did_collide = True
                    cart2.did_collide = True
                    found_collision = True
                    break

            if cart.did_collide:
                break

        return found_collision

    @staticmethod
    def _try_get_cart(row_index, cell_index, direction):
        if direction in _CART_DIRECTIONS:
            return Cart(row_index, cell_index, direction)

        return None


def run_part1(file_content):
    """Implmentation for Part 1."""
    track = CartTrack(file_content)
    return track.find_first_crash()


def run_part2(file_content):
    """Implmentation for Part 2."""
    track = CartTrack(file_content)
    return track.find_last_cart()


if __name__ == "__main__":
    import sys

    def _print_coordinates(coordinates):
        return "{},{}".format(coordinates[0], coordinates[1])

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(_print_coordinates(run_part1(file_content))))
            print("Part 2: {}".format(_print_coordinates(run_part2(file_content))))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])

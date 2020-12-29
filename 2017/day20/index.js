const fs = require('fs');

class Point {
  constructor(x, y, z) {
    if (!Number.isFinite(x) || !Number.isFinite(y) || !Number.isFinite(z)) {
      throw new Error('Invalid point value');
    }

    this._x = x;
    this._y = y;
    this._z = z;
  }

  get magnitude() {
    return Math.abs(this._x) + Math.abs(this._y) + Math.abs(this._z);
  }

  add(other) {
    this._x += other._x;
    this._y += other._y;
    this._z += other._z;
  }

  compareMagnitude(other) {
    return this.magnitude - other.magnitude;
  }

  toString() {
    return `${this._x},${this._y},${this._z}`;
  }
}

class Particle {
  constructor(id, position, velocity, acceleration) {
    this._id = id;
    this._position = position;
    this._velocity = velocity;
    this._acceleration = acceleration;
    this._destroyed = false;
  }

  get id() {
    return this._id;
  }

  get position() {
    return this._position;
  }

  get destroyed() {
    return this._destroyed;
  }

  runTick() {
    this._velocity.add(this._acceleration);
    this._position.add(this._velocity);
  }

  compareMagnitude(other) {
    let magnitude = this._acceleration.compareMagnitude(other._acceleration);
    if (magnitude !== 0) {
      return magnitude;
    }

    magnitude = this._velocity.compareMagnitude(other._velocity);
    if (magnitude !== 0) {
      return magnitude;
    }

    return this._position.compareMagnitude(other._position);
  }

  destroy() {
    this._destroyed = true;
  }
}

class Day20 {
  static *parseParticles(str) {
    const regexp = /^p=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, v=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, a=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>$/gm;
    let currentId = 0;
    let result = null;

    do {
      result = regexp.exec(str);
      if (result !== null) {
        const position = new Point(
          Number.parseInt(result[1]),
          Number.parseInt(result[2]),
          Number.parseInt(result[3])
        );
        const velocity = new Point(
          Number.parseInt(result[4]),
          Number.parseInt(result[5]),
          Number.parseInt(result[6])
        );
        const acceleration = new Point(
          Number.parseInt(result[7]),
          Number.parseInt(result[8]),
          Number.parseInt(result[9])
        );

        yield new Particle(currentId++, position, velocity, acceleration);
      }
    } while (result != null);
  }

  static runPart1(str) {
    const particles = Array.from(this.parseParticles(str));
    particles.sort((a, b) => {
      return a.compareMagnitude(b);
    });

    return particles[0].id;
  }

  static runPart2(str) {
    const particles = Array.from(this.parseParticles(str));
    let particleCount = -1;
    let countSame = 0;

    for (;;) {
      const particleMap = new Map();
      let currentCount = 0;

      for (let i = 0; i < particles.length; i++) {
        const particle = particles[i];

        if (!particle.destroyed) {
          particle.runTick();
          const position = particle.position.toString();

          if (particleMap.has(position)) {
            particleMap.get(position).destroy();
            particle.destroy();
            currentCount--;
          } else {
            particleMap.set(position, particle);
            currentCount++;
          }
        }
      }

      if (particleCount === currentCount) {
        if (++countSame > 1000) {
          break;
        }
      } else {
        particleCount = currentCount;
        countSame = 0;
      }
    }

    return particleCount;
  }

  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');

    return [this.runPart1(fileContent), this.runPart2(fileContent)];
  }
}

module.exports = Day20;

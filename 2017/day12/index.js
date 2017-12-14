const fs = require('fs')

class ProgramNode {
  constructor (id) {
    this._id = id
    this._neighbors = []
  }

  get id () {
    return this._id
  }

  get neighbors () {
    return this._neighbors
  }

  addNeighbor (node) {
    this._neighbors.push(node)
  }
}

class ProgramGraph {
  constructor () {
    this._map = new Map()
  }

  addRawData (rawData) {
    const node = this.findOrCreateNode(rawData[0])

    for (const childId of rawData.slice(1)) {
      const neighbor = this.findOrCreateNode(childId)
      node.addNeighbor(neighbor)
      neighbor.addNeighbor(node)
    }
  }

  findOrCreateNode (id) {
    let node = this._map.get(id)

    if (node === undefined) {
      node = new ProgramNode(id)
      this._map.set(id, node)
    }

    return node
  }

  countNodesInGroup (id, visited = new Set()) {
    const processingQueue = []
    processingQueue.push(this.findOrCreateNode(id))

    while (processingQueue.length > 0) {
      const current = processingQueue.shift()
      visited.add(current.id)

      for (const neighbor of current.neighbors) {
        if (!visited.has(neighbor.id)) {
          visited.add(neighbor.id)
          processingQueue.push(neighbor)
        }
      }
    }

    return visited.size
  }

  countGroups () {
    const visited = new Set()
    let groupCount = 0

    for (const startingNode of this._map.values()) {
      if (!visited.has(startingNode.id)) {
        this.countNodesInGroup(startingNode.id, visited)
        groupCount++
      }
    }

    return groupCount
  }
}

class Day12 {
  static * parseRows (str) {
    let current = ''
    let row = []

    for (const ch of str) {
      if (ch === '\n' || ch === ' ' || ch === ',') {
        if (current.length > 0) {
          row.push(Number.parseInt(current))
          current = ''
        }

        if (ch === '\n') {
          yield row
          row = []
        }
      } else if (ch >= '0' && ch <= '9') {
        current += ch
      }
    }
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')

    const graph = new ProgramGraph()
    for (const row of this.parseRows(fileContent)) {
      graph.addRawData(row)
    }

    return [
      graph.countNodesInGroup(0),
      graph.countGroups()
    ]
  }
}

module.exports = Day12

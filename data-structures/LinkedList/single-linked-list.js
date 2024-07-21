class Node {
  constructor(value = null) {
    this.value = value;
    this.next = null;
    this.previous = null; // if list is double linked list.
  }
}

class LinkedList {
  constructor() {
    this.first = null;
  }

  addFirst(value) {
    const node = new Node(value);
    node.next = this.first;
    this.first = node;
  }

  addLast(value) {
    const node = new Node(value);
    if (this.first) {
      let currentNode = this.first;
      while (currentNode && currentNode.next) {
        currentNode = currentNode.next;
      }
      currentNode.next = node;
    } else {
      this.first = node;
    }
  }

  removeLast() {
    let target = null;
    let current = this.first;
    if (current && current.next) {
      while (current && current.next && current.next.next) {
        current = current.next;
      }
      target = current.next;
      current.next = null;
    } else {
      target = this.first;
    }
    if (target) {
      return target;
    }
  }

  removeFirst() {
    const first = this.first;
    if (first) {
      this.first = first.next;
      return first.value;
    }
  }

  contains(value) {
    let index = 0;
    for (let current = this.first; current; current = current.next) {
      if (current.value === value) {
        return index;
      }
      index++;
    }
    return index;
  }

  removeAt(nth) {
    if (nth === 0) {
      return this.removeFirst;
    }

    for (
      let current = this.first, index = 0;
      current;
      index++, current = current.next
    ) {
      if (index === nth) {
        if (!current.next) {
          return this.removeLast();
        }
        current.previous = current.next;
        return current.value;
      }
    }
  }
}

interface Bird {
  fly(): void;
  walk(): void;
}
class Duck implements Bird {
  fly() {
    console.log('duck is flying');
    // Duck can fly
  }
  walk() {
    console.log('duck is walking');
    // Duck can walk
  }
}
class Ostrich implements Bird {
  fly() {
    throw new Error('ostrich can not fly');
    // Ostrich can not fly
  }
  walk() {
    console.log('ostrich is walking');
    // Ostrich can walk
  }
}

// interface BirdFly {
//   fly(): void;
// }
// interface BirdWalk {
//   walk(): void;
// }
// class Duck implements BirdFly, BirdWalk {
//   fly() {
//     // Duck can fly
//     console.log('duck is flying');
//   }
//   walk() {
//     // Duck can walk
//     console.log('duck is walking');
//   }
// }
// class Ostrich implements BirdWalk {
//   walk() {
//     // Ostrich can walk
//     console.log('ostrich is walking');
//   }
// }

const o1 = new Ostrich();
const b1 = new Duck();

b1.walk();
b1.fly();

o1.walk();
o1.fly();

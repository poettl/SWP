interface BirdFly {
  fly(): void;
}
interface BirdWalk {
  walk(): void;
}
class Duck implements BirdFly, BirdWalk {
  fly() {
    // Duck can fly
    console.log('duck is flying');
  }
  walk() {
    // Duck can walk
    console.log('duck is walking');
  }
}
class Ostrich implements BirdWalk {
  walk() {
    // Ostrich can walk
    console.log('ostrich is walking');
  }
}

const o1 = new Ostrich();
const d1 = new Duck();

d1.walk();
d1.fly();

o1.walk();

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

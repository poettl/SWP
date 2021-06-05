var Duck = /** @class */ (function () {
    function Duck() {
    }
    Duck.prototype.fly = function () {
        // Duck can fly
        console.log('duck is flying');
    };
    Duck.prototype.walk = function () {
        // Duck can walk
        console.log('duck is walking');
    };
    return Duck;
}());
var Ostrich = /** @class */ (function () {
    function Ostrich() {
    }
    Ostrich.prototype.walk = function () {
        // Ostrich can walk
        console.log('ostrich is walking');
    };
    return Ostrich;
}());
var o1 = new Ostrich();
var d1 = new Duck();
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

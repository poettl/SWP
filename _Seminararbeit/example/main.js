var Duck = /** @class */ (function () {
    function Duck() {
    }
    Duck.prototype.fly = function () {
        console.log('duck is flying');
        // Duck can fly
    };
    Duck.prototype.walk = function () {
        console.log('duck is walking');
        // Duck can walk
    };
    return Duck;
}());
var Ostrich = /** @class */ (function () {
    function Ostrich() {
    }
    Ostrich.prototype.fly = function () {
        throw new Error('ostrich can not fly');
        // Ostrich can not fly
    };
    Ostrich.prototype.walk = function () {
        console.log('ostrich is walking');
        // Ostrich can walk
    };
    return Ostrich;
}());
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
var o1 = new Ostrich();
var b1 = new Duck();
b1.walk();
b1.fly();
o1.walk();
o1.fly();

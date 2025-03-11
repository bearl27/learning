"use strict";
const price = {
    apple: 100,
    banana: 200,
    orange: 300,
};
console.log(price.apple); // 100
price.grape = 400; // OK
const data = {
    foo: 123,
};
const num_foo = data.bar;
console.log(num_foo);
const obj__1 = { foo: "foo1", bar: 1 };
const obj__2 = { foo: "foo2", bar: 2, baz: true, };
console.log(obj__1.baz); // undefine
console.log(obj__2.baz); // true
const objjj = {
    foo: 123,
    bar: 456,
};
const obj2 = {
    foo: 987,
    bar: 654,
};
const commandList = ["attack", "defend", "run"];

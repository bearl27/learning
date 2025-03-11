"use strict";
const foo = {
    num: 1,
};
const bar = foo;
const zoo = { ...foo };
console.log(foo === bar); // true
console.log(foo === zoo); // false
console.log(bar.num); // 1
foo.num = 2;
console.log(bar.num); // 2
bar.num = 3;
console.log(foo.num); // 3
const foo2 = {
    num: 1,
};
const bar2 = { ...foo2 };
foo2["num"] = 3;
console.log(bar2["num"]);
const foo3 = {
    num: 1,
    obj: {
        s_num: 2,
    }
};
const bar3 = { ...foo3 };
foo3.obj.s_num = 3;
console.log(bar3.obj.s_num); // 3
let num_hoge = 1;
let num_fuga = num_hoge;
console.log("numfuga", num_fuga); // 1
num_hoge = 2;
console.log("numhoge", num_hoge); // 2
console.log("numfuga", num_fuga); // 1

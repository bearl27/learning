"use strict";
const obj_type = {
    name: "hoge",
    age: 20,
};
console.log(typeof obj_type["name"]); // string
console.log(typeof obj_type["age"]); // number
const user = {
    name: "hoge",
    age: 20,
};
console.log(typeof user); // object
const hoge = "hoge";
//const hoge2: str = 1; // Error
console.log(hoge); // hoge
console.log(typeof hoge); // string

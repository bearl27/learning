"use strict";
console.log(void console.log("test")); // undefined (testも出力)
let arr = [1, 2, 3, 4];
console.log(delete arr[1]); // true
delete arr[1];
console.log(arr); // [1, empty, 3, 4]
console.log(arr.length); // 4
console.log(arr[1]); // undefined
console.log(1 in arr); // true
console.log(2 in arr); // false
// console.log(null in arr); // error
// console.log(undefined in arr); // error
const type_number = 123;
console.log(typeof type_number); // number
console.log(typeof "123"); // string

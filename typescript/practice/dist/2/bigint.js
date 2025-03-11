"use strict";
const num = 100n;
console.log(num); // 100n
const bigint = BigInt(1e8);
console.log(bigint); // 100000000n
const a = BigInt(-1);
const b = -1;
console.log(BigInt(b)); // -1n
//const wrong = BigInt(1.5);
const googol = BigInt(10) ** BigInt(100); // 1000...n
const result = 5n / 2n; // 2n
console.log(googol);
console.log(result);

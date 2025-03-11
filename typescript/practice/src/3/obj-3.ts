type PriceData = {
    [key: string]: number,
}
const price: PriceData = {
    apple: 100,
    banana: 200,
    orange: 300,
};
console.log(price.apple); // 100

price.grape = 400; // OK
// price.melon = "500"; // Error

type MyData = {
    [key: string]: number,
}
const data: MyData = {
    foo: 123,
}
const num_foo: number = data.bar;
console.log(num_foo);

type ObjType = {
    foo: string,
    bar: number,
    baz?: boolean
}
const obj__1: ObjType = { foo:"foo1", bar:1 };
const obj__2: ObjType = { foo:"foo2", bar:2, baz: true, };

console.log(obj__1.baz); // undefine
console.log(obj__2.baz); // true

const objjj = {
    foo: 123,
    bar: 456,
}
type T = typeof objjj;
const obj2: T = {
    foo: 987,
    bar: 654,
}

const commandList = ["attack", "defend", "run"] as const;

type Command = typeof commandList[number];

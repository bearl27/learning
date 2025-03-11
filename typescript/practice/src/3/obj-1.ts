const obj = {
    name: "hoge",
    age: 20,
};
console.log(obj.name);
console.log(obj.age);
console.log(obj["name"]);
console.log(obj["age"]);

// const name = input ? input.name : "default";
// const user = {
//     name,
//     age: 20,
// }

const num_obj = {
    1: "one",
    3.14: "pi",
}
console.log(num_obj[1]);
console.log(num_obj[3.14]);
// console.log(num_obj.1);
// console.log(num_obj.3.14);

const propaty = "name";
const obj_1 = {
    [propaty]: "hoge",
    age: 20,
}
console.log(obj_1.name);

const obj_2 = {
    one: 10,
    2: "tow"
}
console.log(obj_2.one);
console.log(obj_2[2]);
obj_2.one = 1;
console.log(obj_2.one);
obj_2[2] = "two";
console.log(obj_2[2]);

const obj_3 = {
    name: "hoge",
    3.14: "pi",
}
console.log(obj_3["name"]);
console.log(obj_3["3.14"]);


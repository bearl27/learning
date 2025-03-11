const pre_obj1 = {
    name: "hoge",
    age: 20,
}
const obj1 = {
    ...pre_obj1,
    age: 30,
    sex: "man",
}

console.log(obj1.name);
console.log(obj1.age);



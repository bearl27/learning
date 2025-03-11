const obj_type: {
    name: string,
    age: number,
} = {
    name: "hoge",
    age: 20,
}

console.log(typeof obj_type["name"]); // string
console.log(typeof obj_type["age"]);  // number

type UserType = {
    name: string,
    age: number,
}

const user: UserType = {
    name: "hoge",
    age: 20,
}

console.log(typeof user); // object

type str = string;
const hoge: str = "hoge";
//const hoge2: str = 1; // Error
console.log(hoge); // hoge
console.log(typeof hoge); // string

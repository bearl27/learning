type FooBar = {
    foo: string;
    bar: number;
};

type FooBarBaz = {
    foo: string;
    bar: number;
    baz: boolean;
};

const obj_4: FooBarBaz = {
    foo: 'foo',
    bar: 1,
    baz: true,
};

const obj_41: FooBar = obj_4;
console.log(obj_41);

const obj_42: FooBar = {
    foo: 'foo',
    bar: 1,
}

// const obj_43: FooBarBaz = obj_42;
// console.log(obj_43);
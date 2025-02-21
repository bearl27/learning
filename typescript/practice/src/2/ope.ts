const str_num: string = "123";
console.log(typeof +str_num);
console.log(`orange < apple: ${"orange" < "apple"}`);

// ==と===の違い
const num_1: any = 1;
const str_1: any = "1";
const any_1: any = 1;
console.log(num_1 == str_1);
console.log(num_1 === str_1);
console.log(any_1 == "1");   //
console.log(any_1 === "1");

// nullとundefined
console.log(null == undefined);
console.log(null === undefined);

// NaN
const nan: number = NaN;
console.log(isNaN(nan));

const null_value: null = null;
const num_11: number = 123;
console.log(!null_value);
console.log(!num_11);

const str_empoty: string = "";
console.log(str_empoty || "name");
console.log(str_empoty && "name");

// ||と??の違い
const str_null: null = null;
const str_undef: undefined = undefined;
const str_empty: string = "";
const str_name: string = "name";
console.log("null","||",str_null || str_name, "??",str_null ?? str_name);
console.log("undef","||",str_undef || str_name, "??",str_undef ?? str_name);
console.log("empty","||",str_empty || str_name, "??",str_empty ?? str_name);
"use strict";
const input_number = 10;
if (input_number < 100)
    console.log("100以上の数字を入力してください");
let userName = "";
if (userName === "") {
    console.log("名前を入力してください");
    userName = "名無し";
}
else {
    const message = `Welcome, ${userName}!`;
    console.log(message);
}
if (input_number < 100) {
    console.log("100未満の数字です");
}
else if (input_number < 0) {
    console.log("0未満の数字です");
}
else {
    console.log("100以上の数字です");
}

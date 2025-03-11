"use strict";
const str1 = 'Hello, World!';
const str2 = 'Hello World!';
const str3 = `Hello,
World!`;
const hello = 'Hello';
const world = 'World';
const str4 = `${hello}, ${world}!`;
console.log(str1);
console.log(str2);
console.log(str3);
console.log(str4);
console.log(str3.length);
const bool = typeof str1 === typeof str2;
console.log(bool);
const num1 = 1;
const num2 = 2;
const str5 = `${num1} + ${num2} = ${num1 + num2}`;
console.log(str5);
console.log(str1 === str3);
// 全てのエスケープシーケンスのデモ
console.log("1. 改行テスト:\nこれは次の行です");
console.log("2. キャリッジリターン:\rこの文字で上書き");
console.log("3. タブ:\tスペースが空きます");
console.log("4. バックスペース:Hello\bWorld");
console.log("5. フォームフィード:\fページ区切り");
console.log('6. シングルクォート: It\'s a beautiful day');
console.log("7. ダブルクォート: 彼は\"こんにちは\"と言いました");
console.log("8. バックスラッシュ: C:\\Program Files\\App");
console.log("9. Unicode文字: \u3053\u3093\u306B\u3061\u306F"); // こんにちは
console.log("10. ASCII文字: \x41\x42\x43"); // ABC
console.log("11. ヌル文字: End\0of text");

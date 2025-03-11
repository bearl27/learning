import { createInterface } from "readline";
const rl = createInterface({
    input: process.stdin,
    output: process.stdout,
});
rl.question("input number: ", (answer) => {
    const num = +answer;
    num > 100 ? console.log("100より大きいです") : console.log("100以下です");
    rl.close();
});

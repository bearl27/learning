import { createInterface } from "readline";
const rl = createInterface({
    input: process.stdin,
    output: process.stdout
});
const messages = {
    good: "0以上の数値が入力されました",
    bad: "正の数値を入力してください"
};
rl.question("数値を入力してください:", (line) => {
    const num = Number(line);
    console.log(messages[num >= 0 ? "good" : "bad"]);
    rl.close();
});

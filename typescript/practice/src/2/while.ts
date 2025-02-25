let sum = 0;
while (sum < 10) {
    sum++;
}
console.log(sum); // 9

sum = -1;
while (sum < 10) {
    if(sum < 0){
        console.log("負の数です");
        break;
    }
    sum++;
}
console.log(sum); // -1

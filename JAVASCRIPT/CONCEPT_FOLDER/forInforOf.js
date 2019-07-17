//for...in
//오브젝트의 property에 루프를 돌수 있도록 설계됨
let obj = {a: 1, b:2, c:3};
for(let element in obj){
    if(!obj.hasOwnProperty(element)) continue;
    console.log(element);
    console.log(obj[element]);
}

//for...of
//ES6 문법 iterable, array 둘 다 가능
//인덱스를 몰라도 문제가 없을 경우
const hand = ["Brazil", "Korea", "America"];
for(let element of hand){
    console.log(`${element}`);
}
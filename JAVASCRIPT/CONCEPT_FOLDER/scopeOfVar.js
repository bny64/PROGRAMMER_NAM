//자바스크립트의 var변수는 단위가 블록
var i = 10;
for(; i<100; i++){
    console.log(i);
}

for(var j=0; j<100; j++){
    console.log(j);
}

console.log(j);
//for은 블록이 아니기 때문에 글로벌 변수가 된다.
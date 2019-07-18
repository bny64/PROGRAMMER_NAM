//배열을 복사할 때 자바스크립트는 얕은 복사를 한다.
let a = [1,2,3,4,5];
let b = a;

a[2] = 4;
console.log(a); //1,2,4,4,5
console.log(b); //1,2,4,4,5

//배열을 복사할 때 slice메서드를 사용해서 복사가능.
let a = [1,2,3,4,5];
let b = Array.prototype.slice.call(a);  //더 간단한 방법은 [].slice.call()
a[2] = 4;
console.log(a); //1,2,4,4,5
console.log(b); //1,2,3,4,5

//함수에서 체크
function func1(){
    let c = [].slice.call(arguments);
    console.log(c);
}

func1([1,2,3,4,5]);

//ES6에서는 확산연산자로 사용가능
function func2(...restParam){
    let c = restParam; //확산연산자는 배열이다.
    console.log(c);
}

func2([1,2,3,4,5]);

//다른 방법
//ES6 문법
function func3(){
    let c = Array.from(arguments);
    console.log(c);
}

func3([1,2,3,4,5]);
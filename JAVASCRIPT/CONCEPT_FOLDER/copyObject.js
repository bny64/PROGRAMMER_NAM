//原始型以外、String, Object, Array などなど
// 原始型以外(=)この符号は参照コピー
//얕은 복사
var foo = {key:"value"};
var bar = $.extend({}, foo);
//복사하고 싶은 오브젝트를 extend function의 파라미터로 넣는다
foo.key = "other value";

console.log(foo);
console.log(bar);

//vanila.js 방법
function clone(obj){

    //타입이 object가아닌 혹은 null인 경우
    if(obj === null && typeof(obj) !== "object") return;
  

    var copy = obj.constructor();//생성자
    
    for(var attr in obj){
        //프로퍼티를 한개씩 복사
        if(obj.hasOwnProperty(attr)){
        copy[attr] = obj[attr];
        }
    }
    return copy;
}

var foo = {key:"value"};
var bar = clone(foo);
foo.key = "other value";

console.log(foo);
console.log(bar);

//ES6부터 복사가 가능한 메서드가 만들어짐
const object = {a:1, b:2, c:3};

const object2 = Object.assign({c:5, d:6}, object);
//Object.assign(타켓、복사하려는 객체)
console.log(object2);
//Object{c=3, d:6, a=1, b=2}

//깊은 복사
// Deep Clone
obj1 = { a: 0 , b: { c: 0}};
let obj3 = JSON.parse(JSON.stringify(obj1));
obj1.a = 4;
obj1.b.c = 4;
console.log(JSON.stringify(obj3)); // { a: 0, b: { c: 0}}

const o1 = { a: 1, b: 1, c: 1 };
const o2 = { b: 2, c: 2 };
const o3 = { c: 3 };

const obj = Object.assign({}, o1, o2, o3);
console.log(obj); // { a: 1, b: 2, c: 3 }뒤에서 부터 겹친다.
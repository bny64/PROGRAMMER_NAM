var sym1 = Symbol();
var sym2 = Symbol("foo");

console.log(Symbol("foo")===Symbol("foo"));;
//기대한 출력값은 false

//symbol은 원시데이터 형태
//new Symbol()을 통해 만들 수 없다.
//명시적으로 객체를 생성해야 한다.

var sym = Symbol("foo");
var symObj = Object(sym);

//이러한 방법으로 객체가 생성된다.
//symbol은 문자열도 아니고 객체도 아니다.

const obj = {};
const mySymbol = Symbol("mySymbol");
obj[mySymbol] = 123;

console.log(obj);
console.log(obj[mySymbol]);

//symbol은 중복될 수 없기 때문에, 먼저 선언 후 사용해야 한다.

let myName = Symbol("Paolo");
console.log(myName.toString());
//Symbol(Paolo)
let obj2 = {};
obj2.myName = "Smith";
obj2[myName] = "Jenny";
console.log(obj2);
//{myName: "Smith", Symbol(Paolo): "Jenny"}
console.log(obj2.myName);
//Smith
console.log(obj2[myName]);
//Jenny
//보통 중괄호를 사용하면 프로퍼티에 직접접근
//하지만 객체에 symbol이 있는 경우 내 생각에는 심볼값에 접근하는 것 같다.

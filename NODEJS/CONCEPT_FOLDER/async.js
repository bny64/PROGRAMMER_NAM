//일반
let first = 10;
let second = 20;
let result = 0;

var add = (x, y) => x + y;

setTimeout(()=>{
    result = add(first, second);
    console.log(result);
}, 1000);

first = 20;

//고전적인 callback 함수 사용
let first = 10;
let second = 20;
let result = 0;

var add = (x, y) => x + y;

var getResult = callback => {
    setTimeout(()=>{
        result = add(first, second);
        console.log(result);
        callback();
    },1000);
}

getResult(()=>{
    first = 20;
});
//1초 뒤 30 출력

//ES6 promise
let first = 10;
let second = 20;
let result = 0;

var add = (x, y) => x + y;

const getResult = new Promise((resolve, reject)=>{
    setTimeout(()=>{
        result = add(first, second);
        console.log(result);
        resolve();
    }, 1000);
});

const getResult = () => {
    return new Promise((resolve, reject)=>{
        setTimeout(()=>{
            result = add(first, second);
            console.log(result);
            resolve();
        }, 1000);
    });
};

getResult().then(()=>{
    first = 20;
    console.log(`then`);
});

//ES7 async, await
let first = 10;
let second = 20;
let result = 0;

var add = (x, y) => x + y;

const getResult = async () => {
    return new Promise((resolve, reject)=>{
        setTimeout(()=>{
            result = add(first, second);
            console.log(result);
            resolve();
        }, 1000);
    });
};

(async ()=>{
    await getResult();
    first = 20;
    console.log(first);
})();
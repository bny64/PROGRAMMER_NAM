
class Vehicle{
    constructor(){
        this.passengers = [];
        console.log('Vehicle created');
    }
    addPassenger(p){
        this.passengers.push(p);
    }
}

class Car extends Vehicle{
    constructor(){
        super();
        console.log(`Car created`);       
    }
    deployAirbags(){
        console.log(`BWOOSH!`);
    }
}

const v = new Vehicle();
v.addPassenger('Frank');
v.addPassenger('Judy');
console.log(v.passenger);
const c = new Car();
c.addPassenger('Allice');
c.addPassenger('Cameron');
c.passenger;
v.deployAirbags();
c.deployAirbags();
I think the dice faces since they are fixed, can use arrays which are fixed in length, hence not takeing too much munneccessary memory.


### Arrays

declaring an array length of 3 which can hold 3 intigers. -> var nums [3]int

// I am not sure if an array can hold multiple data types, since some of the types are numbers,
// like the temparature what later on should be modified to can be changed between farenheit or 
// celsius.

How would I start with this project on javascript:

I have to store the dice faces like:

const temp-dice = [98; 85; 90; 92; 75; 100];

Math.random() should give us a number between 1-5 and we have to round it up? I think to give a whole
number insted of a floating point and we have to print out the number from the dice face whihc is in that position.

what I actually have to turn in to golang script is :

const dice = () => {
  const temp_dice = [95, 75, 90, 80, 85, 'your choice'];
  
  let res = temp_dice[Math.floor(Math.random()*6)];
  console.log(res);
}

dice();


The dices I have and I have to implement:

| Die   | Controls             | Faces                                                                                                                |
| ----- | -------------------- | -------------------------------------------------------------------------------------------------------------------- |
| **1** | Grind + contact time | Very fine / 30 sec; Fine / 60 sec; Medium-fine / 90 sec; Medium / 120 sec; Coarse / 4 min; Your choice               |
| **2** | Coffee + water       | 30g/200g + dilute; 24g/200g + dilute; 15g/250g + dilute; 15g/200g; 12g/200g; Your choice                             |
| **3** | Method + bloom       | Inverted + 30s/60g; Standard + 30s/60g; Inverted + 30s/30g; Standard + 30s/30g; Standard/no bloom; Inverted/no bloom |
| **4** | Agitation            | No stir; stir ×2 before press; CW once + CCW once before press; stir ×1 before press; N/S/E/W stir; Your choice      |
| **5** | Temperature          | 95°C; 75°C; 90°C; 80°C; 85°C; Your choice                                                                            |



Todos:

Time conversion
Farentheit conversion

output design

options to choose between farenheit and celsius

abstraction of code to multiple files maybe?

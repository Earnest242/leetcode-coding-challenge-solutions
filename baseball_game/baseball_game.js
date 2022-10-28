//function to perform the points calculation operation
var calPoints = function(ops){
    var result =null;
    let stack1 = []
    for (let i in ops) {
        if (ops[i] == "+"){
            let p = stack1[stack1.length-1] + stack1[stack1.length-2]
            stack1.push(p)
        }
        else if (ops[i] == "C"){
            stack1.pop()
        }
        else if (ops[i] == "D"){
            let n = stack1[stack1.length-1]*2
            stack1.push(n)
        }
        else {
            stack1.push(Number(ops[i]))
        }

    }
    for (let f in stack1){
        result =result + stack1[f]
    }
    
    return result;
}
//insert different test cases into this array
var ops3 =["5","2","C","D","+"]

//calling the calPoints function and outputting the result
console.log(calPoints(ops3))
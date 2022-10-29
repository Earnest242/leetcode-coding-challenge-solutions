
def calPoints(ops):
    result = 0
    stack1 = []

    for i in ops:
        length = len(stack1)
        if i == "+":
            p = stack1[length-1]+stack1[length-2]
            stack1.append(p)
        elif i == "C":
            stack1.pop()
        elif i == "D":
            n = stack1[length-1] * 2
            stack1.append(n)
        else :
            stack1.append(int(i))
    
    for f in stack1:
        result = result + f
    
    return result

ops3 = ["5","2","C","D","+"]
print(calPoints(ops3))
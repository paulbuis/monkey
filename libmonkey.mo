let concat = fn(a, b) {
    if (len(b) == 0) {
        a
    } else {
        concat( push(a, first(b)), rest(b) )
    };
};
let unshift = fn (x, a) { concat([x], a); };

let contains = fn(set) { 
    fn(x) {
        if (len(set) == 0) {
            false
        } else {
            if (first(set) == x) {
                true
            } else {
                contains(rest(set))(x)
            }
        };
    };
};
let setAdd = fn(set, x) {
    if (contains(set)(x)) { set } else { push(set, x) };
};

let filter = fn(predicate){
   fn(array) {
        let tailRecurse = fn(a, output){ 
            if (len(a) == 0) {
                output
            } else {
                if (predicate(first(a))) {
                    tailRecurse(rest(a), push(output, first(a)) )
                } else {
                    tailRecurse(rest(a), output)
                }
            };
        };
        tailRecurse(array, []);
   };
};

let not = fn(predicate) {
    fn(x) { !predicate(x); };
};
let or = fn(predA, predB){
    fn(x) {
        if (predA(x)) { true } else { predB(x) };
    };
};
let and = fn(predA, predB) {
    fn(x) {
        if (!predA(x)) { false } else { predB(x) };
    };
};

let reject = fn(predicate){ filter(not(predicate)); };
let intersect = fn(setA, setB){ filter(contains(setB)) (setA); };
let difference = fn(setA, setB){ filter(not(contains(setB))) (setA); };
let union = fn(setA, setB){ concat(setA, difference(setB, setA)); };

let map = fn(transform){
   fn(array) {
        let tailRecurse = fn(a, output){
            if (len(a) == 0) {
                output
            } else {
                tailRecurse(rest(a), push(output, transform(first(a)) ) )
            }
        };
        tailRecurse(array, []);
   };
};

let reduce =  fn(combine, init){
    fn(array) {
        let tailRecurse = fn(input, output){
            if (len(input) == 0) {
                output
            } else {
                tailRecurse(rest(input), combine(output, first(input)))
            };
        };
        tailRecurse(array, init);
    };
};
let sum = reduce( fn(x, y){ x + y } , 0);
let distinct = fn(array){
    reduce( fn(set, x){ setAdd(set, x) }, [] ) (array);
};

let zip = fn(arrayA, arrayB) {
    let tailRecurse = fn(A, B, output) {
        if (len(A) == 0) {
            output
        } else {
            tailRecurse(rest(A), rest(B), push(output, [first(A), first(B)] ))
        };
    };
    tailRecurse(arrayA, arrayB, []);
}

let drop = fn(n, a) {
    let tailRecurse = fn(n, a, output) {
        if (len(a) == 0) {
            output
        } else {
            if (n > 0) {
                tailRecurse(n-1, rest(a), output)
            } else {
                tailRecurse(0, rest(a), push(output, first(a)))
            }
        };
    };
    tailRecurse(n, a, []);
};

let take = fn(n, a) {
    let tailRecurse = fn(n, a, output){
        if (len(a) == 0) {
            output
        } else {
            if (n == 0) {
                output
            } else {
                tailRecurse(n-1, rest(a), push(output, first(a)))
            }
        };
    };
    tailRecurse(n, a, []);
};

let slice = fn(a, min, max) { take(max-min, drop(min, a)); };

let sort = fn(a) {
    let merge = fn(a, b) { concat(a,b) };
    if (len(a) < 2) {
        a
    } else {
        let middle = len(a)/2;
        let firstHalf = sort(take(middle, a));
        let secondHalf = sort(drop(middle, a));
        merge(firstHalf, secondHalf);
    };
};
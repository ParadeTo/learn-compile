
c = 2
function fun2 () {
  var a = 1
  return function fun1() {
    console.log(a)
  }
}

fun2()()
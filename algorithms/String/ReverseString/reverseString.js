//! Write an algorithm to reverse the String.

// time complexity O(n^2)
// Strings are immutable so every loop iteration new string is created with all previous characters mapped into new String.

function StringReverse0(str) {
  let reversed = "";
  for (let i = str.length; i >= 0; i--) {
    reversed += str[i];
  }
  return reversed;
}

// time complexity O(n+n+n) == 3 O(n) == O(n)
//
function StringReverse2(str) {
  let stringArr = str.split("");
  for (let i = 0, j = str.length - 1; i < j; i++, j--) {
    stringArr[i] = str[j];
    stringArr[j] = str[i];
  }
  return stringArr.join("");
}

// time complexity O(n+n) == O(2n) == 2 O(n) == O(n)
//
function StringReverse1(str) {
  let arr = new Array(str.length);
  for (let i = 0; i < str.length; i++) {
    arr[str.length - 1 - i] = str[i];
  }
  return arr.join("");
}

const dummy_string = "hello";
console.log(StringReverse2(dummy_string));

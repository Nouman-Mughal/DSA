//write an algorithm to replace all the spaces in a string with %20.
// time complexity O(n)
const urlify = (str) => {
  const arr = [];

  for (let i = 0; i < str.length; i++) {
    if (str[i] === " ") {
      arr[i] = "%20";
    } else {
      arr[i] = str[i];
    }
  }
  return arr.join("");
};
console.log(
  urlify("hello world this is the world no one ever found happiness"),
);

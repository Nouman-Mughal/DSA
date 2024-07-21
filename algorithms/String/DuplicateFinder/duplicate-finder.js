// Algorithm to check does strings have duplicates.
// constraint: No extra data structure should be used.
//
// time complexity O(n)
const duplicateFoundUsingSet = (str) => {
  const set = new Set();
  for (let i = 0; i < str.length; i++) {
    set.add(str[i]);
  }
  if (str.length === set.size) {
    return false;
  } else return true;
};
// time complexity O(n^2)
const duplicateFound = (str) => {
  for (let i = 0; i < str.length; i++) {
    for (let j = i + 1; j < str.length; j++) {
      if (str[i] === str[j]) {
        return true;
      }
    }
  }
  return false;
};

console.log(duplicateFoundUsingSet("hello"));

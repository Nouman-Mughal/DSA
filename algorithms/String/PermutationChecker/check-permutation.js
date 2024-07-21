// Give two strings check if one string is permutation of other string.
/*
  first we need to ask the interviewer that if check have to case sensitive or Not. In our case its case sensitive
  e.g. Few and wef are not palindromic.
  also ask if whitespace significant or not.In Our case they are significant.
*/
// Time Complexity is O(N)
function palindromeChecker(str1, str2) {
  let sum1 = 0;
  let sum2 = 0;
  // needed optimization:
  if (str1.length !== str1.length) {
    return false;
  }
  for (let i = 0; i < str1.length; i++) {
    sum1 += str1.charCodeAt(i);
  }

  for (let j = 0; j < str2.length; j++) {
    sum2 += str2.charCodeAt(j);
  }

  if (sum1 === sum2) return true;

  return false;
}

console.log(palindromeChecker("hello     ", "lloeh"));

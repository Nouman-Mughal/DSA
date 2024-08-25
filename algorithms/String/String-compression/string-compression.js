function compressString(input) {
  if (input.length === 0) {
    return " ";
  }

  let result = "";
  let currentChar = input[0];
  let charCount = 1;

  for (let i = 1; i < input.length; i++) {
    if (input[i] === currentChar) {
      charCount++;
    } else {
      result += writeCharAndCount(currentChar, charCount);
      currentChar = input[i];
      charCount = 1;
    }
  }
  // Write the last character and its count because last character not have been written.
  result += writeCharAndCount(currentChar, charCount);

  return result;
}

function writeCharAndCount(char, count) {
  return char + (count > 1 ? count : "");
}

// Example usage:
console.log(compressString("aabbccccccdddddddddd")); // Output: "a2b2c6d10"

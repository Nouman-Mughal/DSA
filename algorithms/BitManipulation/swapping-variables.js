function swapWithXOR(a, b) {
  a = a ^ b;
  b = a ^ b;
  a = a ^ b;
  console.log("result => a: " + a + ", b: " + b);
}

console.log(swapWithXOR(4, 1));

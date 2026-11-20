function processInput(val) {
  const hasLetters = /[a-zA-Z@]/.test(val);
  if (hasLetters) {
    // If it has letters, strip phone formatting residue in case they transitioned from phone to email
    return val.replace(/[\(\)\-\s]/g, "");
  }

  let digits = val.replace(/\D/g, "");
  
  // 11 digits starting with 1 -> strip the 1
  if (digits.length === 11 && digits.startsWith("1")) {
    digits = digits.substring(1);
  }

  // Exactly 10 digits -> format
  if (digits.length === 10) {
    const match = digits.match(/^(\d{3})(\d{3})(\d{4})$/);
    if (match) {
      return `(${match[1]}) ${match[2]}-${match[3]}`;
    }
  }

  // Otherwise, just raw digits
  return digits;
}

console.log(processInput("5550001111"));
console.log(processInput("15550001111"));
console.log(processInput("2(555) 000-1111")); // User adds a 2 at the front
console.log(processInput("(555) 1000-1111")); // User adds a 1 in the middle
console.log(processInput("555000111"));
console.log(processInput("test@test.com"));
console.log(processInput("(555) 000-1111a"));


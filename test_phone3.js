function processInput(val) {
  // If the user types an email, let them type whatever they want, 
  // but strip () and spaces just in case they are leftover from a formatted phone number.
  let stripped = val.replace(/[\(\)\-\s]/g, "");
  
  // If they paste +1, handle it
  if (stripped.startsWith("+1")) {
    stripped = stripped.replace("+1", "1");
  }

  // If the remaining string is ONLY digits, check if we should format it
  if (/^\d+$/.test(stripped)) {
    if (stripped.length === 11 && stripped.startsWith("1")) {
      stripped = stripped.substring(1);
    }
    if (stripped.length === 10) {
      // Format it!
      const match = stripped.match(/^(\d{3})(\d{3})(\d{4})$/);
      return `(${match[1]}) ${match[2]}-${match[3]}`;
    }
  }
  
  // For anything else (emails, partial numbers, numbers with dots), just return the stripped version
  return stripped;
}

console.log(processInput("5550001111"));
console.log(processInput("15550001111"));
console.log(processInput("2(555) 000-1111")); 
console.log(processInput("(555) 1000-1111")); 
console.log(processInput("123.456")); 
console.log(processInput("test+alias@gmail.com"));
console.log(processInput("+1 (555) 000-1111")); 


function processInput(val) {
  const hasLettersOrAt = /[a-zA-Z@]/.test(val);
  
  if (hasLettersOrAt) {
    // It's an email or contains letters. Leave it EXACTLY as typed.
    return val;
  }
  
  // It only has numbers and symbols, so treat it as a potential phone number.
  let stripped = val.replace(/[\(\)\-\s]/g, "");
  if (stripped.startsWith("+1")) {
    stripped = stripped.replace("+1", "1");
  }

  // Only digits remaining?
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
  
  return stripped;
}

console.log(processInput("5550001111"));
console.log(processInput("15550001111"));
console.log(processInput("2(555) 000-1111")); 
console.log(processInput("(555) 1000-1111")); 
console.log(processInput("123.456")); 
console.log(processInput("test-alias@gmail.com"));
console.log(processInput("+1 (555) 000-1111")); 


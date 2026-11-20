import re

with open('packages/utils/src/index.ts', 'r') as f:
    content = f.read()

old_format = """export const formatPhone = (value: string): string => {
  if (!value) return value;
  
  // Clean the input for any non-digit values.
  let phoneNumber = value.replace(/[^\\d]/g, "");
  
  if (phoneNumber.length > 10 && phoneNumber.startsWith("1")) {
    phoneNumber = phoneNumber.slice(1);
  } else if (phoneNumber.length === 11 && !phoneNumber.startsWith("1")) {
    return `+${phoneNumber}`;
  }

  const phoneNumberLength = phoneNumber.length;
  if (phoneNumberLength === 0) return "";
  
  if (phoneNumberLength < 4) return `+1 (${phoneNumber}`;
  if (phoneNumberLength < 7) {
    return `+1 (${phoneNumber.slice(0, 3)}) ${phoneNumber.slice(3)}`;
  }
  return `+1 (${phoneNumber.slice(0, 3)}) ${phoneNumber.slice(3, 6)}-${phoneNumber.slice(6, 10)}`;
};"""

new_format = """export const formatPhone = (value: string): string => {
  if (!value) return value;
  
  // Clean the input for any non-digit values.
  let phoneNumber = value.replace(/[^\\d]/g, "");
  
  // If they pasted a number with a leading 1, strip it so we only handle the 10 digits
  if (phoneNumber.length === 11 && phoneNumber.startsWith("1")) {
    phoneNumber = phoneNumber.slice(1);
  }

  const phoneNumberLength = phoneNumber.length;
  if (phoneNumberLength === 0) return "";
  
  if (phoneNumberLength < 4) return `(${phoneNumber}`;
  if (phoneNumberLength < 7) {
    return `(${phoneNumber.slice(0, 3)}) ${phoneNumber.slice(3)}`;
  }
  return `(${phoneNumber.slice(0, 3)}) ${phoneNumber.slice(3, 6)}-${phoneNumber.slice(6, 10)}`;
};"""

content = content.replace(old_format, new_format)

with open('packages/utils/src/index.ts', 'w') as f:
    f.write(content)


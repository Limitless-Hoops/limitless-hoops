/**
 * Validates if a string is a valid email address.
 */
export const isValidEmail = (email: string): boolean => {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
};

/**
 * Validates if a string is a valid 10-digit US phone number.
 * Assumes the string might be formatted.
 */
export const isValidPhone = (phone: string): boolean => {
  const digitsOnly = phone.replace(/\D/g, "");
  return digitsOnly.length === 10;
};

/**
 * Validates if a string is a valid email or phone number.
 */
export const isValidEmailOrPhone = (identifier: string): boolean => {
  const val = identifier.trim();
  if (!val) return false;
  
  if (val.includes("@")) {
    return isValidEmail(val);
  }
  return isValidPhone(val);
};

/**
 * Validates names to only allow letters, spaces, hyphens, and apostrophes.
 * Eliminates numbers and security risk characters.
 */
export const isValidName = (name: string): boolean => {
  const val = name.trim();
  if (!val) return false;
  return /^[a-zA-Z\s\-']+$/.test(val);
};

/**
 * Formats a numeric string or partially formatted phone string into (xxx) xxx-xxxx
 */
export const formatPhone = (value: string): string => {
  if (!value) return value;
  
  // Clean the input for any non-digit values.
  const phoneNumber = value.replace(/[^\d]/g, "");
  const phoneNumberLength = phoneNumber.length;
  
  // We return the raw value if it's empty
  if (phoneNumberLength === 0) return "";
  
  // Format based on length
  if (phoneNumberLength < 4) return phoneNumber;
  if (phoneNumberLength < 7) {
    return `(${phoneNumber.slice(0, 3)}) ${phoneNumber.slice(3)}`;
  }
  return `(${phoneNumber.slice(0, 3)}) ${phoneNumber.slice(3, 6)}-${phoneNumber.slice(6, 10)}`;
};

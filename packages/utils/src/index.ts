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
  return digitsOnly.length === 10 || (digitsOnly.length === 11 && digitsOnly.startsWith("1"));
};

/**
 * Normalizes a phone number to strict E.164 format (+1XXXXXXXXXX) for API payloads.
 */
export const normalizePhoneForAPI = (phone: string): string => {
  let digitsOnly = phone.replace(/\D/g, "");
  if (digitsOnly.length === 10) {
    digitsOnly = "1" + digitsOnly;
  }
  return `+${digitsOnly}`;
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
  let phoneNumber = value.replace(/[^\d]/g, "");
  
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
};

/**
 * Returns the maximum date string (YYYY-MM-DD) allowed for a given minimum age.
 * E.g., getDateOfBirthMax(18) ensures the selected date is at least 18 years ago.
 * Passing 0 returns today's date (cannot be born in the future).
 */
export const getDateOfBirthMax = (minAgeInYears: number): string => {
  const d = new Date();
  d.setFullYear(d.getFullYear() - minAgeInYears);
  return d.toISOString().split('T')[0];
};

import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

old_onchange = """                  onChange={(e) => {
                    let val = e.target.value;
                    let digits = val.replace(/\\D/g, "");
                    const hasLetters = /[a-zA-Z@]/.test(val);
                    
                    // If they pasted/typed an 11-digit number starting with 1, strip the 1
                    // so it formats neatly into our 10-digit UI pattern
                    if (digits.length === 11 && digits.startsWith("1") && !hasLetters) {
                      val = digits.substring(1);
                      digits = val;
                    }
                    
                    const isTenDigitsPhone = digits.length === 10 && !hasLetters;
                    
                    if (isTenDigitsPhone) {
                      setIdentifier(formatPhone(val));
                    } else if (hasLetters) {
                      // Strip phone formatting if they start typing an email address
                      setIdentifier(val.replace(/[\\(\\)\\-\\s]/g, ""));
                    } else {
                      setIdentifier(val);
                    }
                    if (error) setError("");
                  }}"""

new_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    const hasLetters = /[a-zA-Z@]/.test(val);
                    
                    if (hasLetters) {
                      setIdentifier(val.replace(/[\\(\\)\\-\\s]/g, ""));
                      if (error) setError("");
                      return;
                    }
                    
                    let digits = val.replace(/\\D/g, "");
                    
                    // Allow 11-digit numbers ONLY if they start with 1 (strip the 1 for UI formatting)
                    if (digits.length === 11 && digits.startsWith("1")) {
                      digits = digits.substring(1);
                    }
                    
                    if (digits.length === 10) {
                      setIdentifier(formatPhone(digits));
                    } else {
                      // Harden formatting: if it's not exactly 10 valid digits, fallback to raw digits
                      // so we don't leave broken formatting like "2(555) 000-1111" hanging around.
                      setIdentifier(digits);
                    }
                    if (error) setError("");
                  }}"""

content = content.replace(old_onchange, new_onchange)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


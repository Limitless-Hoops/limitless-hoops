import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

old_onchange = """                  onChange={(e) => {
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

new_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    const hasLettersOrAt = /[a-zA-Z@]/.test(val);
                    
                    if (hasLettersOrAt) {
                      setIdentifier(val);
                      if (error) setError("");
                      return;
                    }
                    
                    let stripped = val.replace(/[\\(\\)\\-\\s]/g, "");
                    if (stripped.startsWith("+1")) {
                      stripped = stripped.replace("+1", "1");
                    }
                    
                    if (/^\\d+$/.test(stripped)) {
                      if (stripped.length === 11 && stripped.startsWith("1")) {
                        stripped = stripped.substring(1);
                      }
                      if (stripped.length === 10) {
                        setIdentifier(formatPhone(stripped));
                      } else {
                        setIdentifier(stripped);
                      }
                    } else {
                      setIdentifier(stripped);
                    }
                    
                    if (error) setError("");
                  }}"""

content = content.replace(old_onchange, new_onchange)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


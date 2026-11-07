import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

old_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    // Format ONLY if it's exactly 10 digits without any letters or @
                    const isTenDigits = /^\\d{10}$/.test(val.replace(/\\D/g, "")) && !/[a-zA-Z@]/.test(val);
                    setIdentifier(isTenDigits ? formatPhone(val) : val);
                    if (error) setError("");
                  }}"""
new_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    const digits = val.replace(/\\D/g, "");
                    const isTenDigitsPhone = digits.length === 10 && !/[a-zA-Z@]/.test(val);
                    const hasLetters = /[a-zA-Z@]/.test(val);
                    
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
content = content.replace(old_onchange, new_onchange)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

# Fix isPhoneInput and onChange
old_onchange = """                  onChange={(e) => {
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

new_onchange = """                  onChange={(e) => {
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
content = content.replace(old_onchange, new_onchange)

old_prefix = """                  prefixText={isPhoneInput ? "+1" : undefined}"""
new_prefix = """                  prefixText={identifier.replace(/\\D/g, "").length === 10 && !/[a-zA-Z@]/.test(identifier) ? "+1" : undefined}"""
content = content.replace(old_prefix, new_prefix)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


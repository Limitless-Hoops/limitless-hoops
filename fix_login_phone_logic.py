import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

# Fix isPhoneInput: exactly 10 digits and NO letters or @
old_is_phone = 'const isPhoneInput = /^\\d|^\\(/.test(identifier);'
new_is_phone = 'const isPhoneInput = /^\\d{10}$/.test(identifier.replace(/\\D/g, "")) && !/[a-zA-Z@]/.test(identifier);'
content = content.replace(old_is_phone, new_is_phone)

# Fix onChange
old_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    setIdentifier(/^\\d|^\\(/.test(val) ? formatPhone(val) : val);
                    if (error) setError("");
                  }}"""
new_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    // Format ONLY if it's exactly 10 digits without any letters or @
                    const isTenDigits = /^\\d{10}$/.test(val.replace(/\\D/g, "")) && !/[a-zA-Z@]/.test(val);
                    setIdentifier(isTenDigits ? formatPhone(val) : val);
                    if (error) setError("");
                  }}"""
content = content.replace(old_onchange, new_onchange)

# Add disclaimer text below the form
old_terms = 'By continuing, you agree to our{" "}'
new_terms = """Only 10-digit US/Canadian phone numbers (+1) are accepted at this time.<br />
          By continuing, you agree to our{" "}"""
content = content.replace(old_terms, new_terms)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


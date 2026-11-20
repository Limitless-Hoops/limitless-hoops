import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

# I will add `const isPhoneInput = /^\\d/.test(identifier);`
if 'const isPhoneInput' not in content:
    content = content.replace('  const [identifier, setIdentifier] = useState("");', '  const [identifier, setIdentifier] = useState("");\n  const isPhoneInput = /^\\d/.test(identifier);')

# I need to conditionally apply formatPhone to identifier if isPhoneInput
# Actually, the user can just type and onChange handles it.
old_onchange = """                  onChange={(e) => {
                    setIdentifier(e.target.value);
                    if (error) setError("");
                  }}"""
# I need to import formatPhone
if 'formatPhone' not in content:
    content = content.replace('import { isValidEmailOrPhone, normalizePhoneForAPI } from "@lh/utils";', 'import { isValidEmailOrPhone, normalizePhoneForAPI, formatPhone } from "@lh/utils";')

new_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    setIdentifier(/^\\d/.test(val) ? formatPhone(val) : val);
                    if (error) setError("");
                  }}"""

content = content.replace(old_onchange, new_onchange)

# Add prefixText
old_input = """                  value={identifier}
                  error={error}"""
new_input = """                  value={identifier}
                  error={error}
                  prefixText={isPhoneInput ? "+1" : undefined}"""

content = content.replace(old_input, new_input)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


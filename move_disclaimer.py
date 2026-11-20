import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

# 1. Remove it from the bottom
old_terms = """Only 10-digit US/Canadian phone numbers (+1) are accepted at this time.<br />
          By continuing, you agree to our{" "}"""
new_terms = 'By continuing, you agree to our{" "}'
content = content.replace(old_terms, new_terms)

# 2. Add it under the input field
old_input_div = """                  }}
                />
              </div>
              <Button"""
new_input_div = """                  }}
                />
                <p className="text-[11px] text-white/40 pt-1">Only 10-digit US/Canadian numbers are allowed.</p>
              </div>
              <Button"""
content = content.replace(old_input_div, new_input_div)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


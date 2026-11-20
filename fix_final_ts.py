import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

content = content.replace('  const [editingEmailId, setEditingEmailId] = useState<string | null>(null);\n  const [editEmailValue, setEditEmailValue] = useState("");\n  const [editingPhoneId, setEditingPhoneId] = useState<string | null>(null);\n  const [editPhoneValue, setEditPhoneValue] = useState("");', '')

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

if 'formatPhone' not in content:
    content = content.replace('normalizePhoneForAPI } from "@lh/utils";', 'normalizePhoneForAPI, formatPhone } from "@lh/utils";')

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)


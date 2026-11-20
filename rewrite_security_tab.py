import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

match = re.search(r'// ─── SecurityTab ───.*?(?=(?:// ─── NotificationsTab ───))', content, re.DOTALL)
if match:
    pass

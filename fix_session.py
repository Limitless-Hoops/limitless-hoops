import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# I will just insert const { session } = useAuth(); into SecurityTab, NotificationsTab, and BillingTab
content = content.replace('function SecurityTab({ household }: { household: any }) {', 'function SecurityTab({ household }: { household: any }) {\n  const { session } = useAuth();')
content = content.replace('function NotificationsTab({ household }: { household: any }) {', 'function NotificationsTab({ household }: { household: any }) {\n  const { session } = useAuth();')
content = content.replace('function BillingTab({ household }: { household: any }) {', 'function BillingTab({ household }: { household: any }) {\n  const { session } = useAuth();')

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


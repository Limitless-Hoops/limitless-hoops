import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Add to SecurityTab
content = re.sub(
    r'function SecurityTab\(\{ household: _household \}: \{ household: any \}\) \{\s*', 
    'function SecurityTab({ household: _household }: { household: any }) {\n  const { session } = useAuth();\n  ', 
    content
)

# Remove unused from BillingTab (error 1110)
content = re.sub(
    r'function BillingTab\(\{ household \}: \{ household: any \}\) \{\n  const \{ session \} = useAuth\(\);',
    r'function BillingTab({ household }: { household: any }) {',
    content
)

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


import re

with open('packages/ui/src/components/ui/button.tsx', 'r') as f:
    content = f.read()

content = content.replace('disabled:pointer-events-none', 'disabled:cursor-not-allowed')

with open('packages/ui/src/components/ui/button.tsx', 'w') as f:
    f.write(content)


import re

with open('/home/premo/.gemini/antigravity/brain/967128c8-22df-4100-85bf-34a7100de4b2/task.md', 'r') as f:
    content = f.read()

content = content.replace('- [ ]', '- [x]')

with open('/home/premo/.gemini/antigravity/brain/967128c8-22df-4100-85bf-34a7100de4b2/task.md', 'w') as f:
    f.write(content)


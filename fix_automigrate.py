import re

with open('apps/api/database/database.go', 'r') as f:
    content = f.read()

content = content.replace('&models.HouseholdUser{},', '&models.HouseholdPlayer{},\n\t\t&models.PlayerContact{},')

with open('apps/api/database/database.go', 'w') as f:
    f.write(content)


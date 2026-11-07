import re

with open('apps/api/main.go', 'r') as f:
    content = f.read()

content = content.replace('api.Post("/households/:id/guardians", handlers.AddGuardian)', 'api.Post("/players/:id/contacts", handlers.AddPlayerContact)')
content = content.replace('api.Delete("/households/:id/guardians/:userId", handlers.RemoveGuardian)', 'api.Delete("/players/:id/contacts/:contactId", handlers.RemovePlayerContact)')

with open('apps/api/main.go', 'w') as f:
    f.write(content)


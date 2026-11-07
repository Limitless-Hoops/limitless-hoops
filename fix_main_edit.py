import re

with open('apps/api/main.go', 'r') as f:
    content = f.read()

if 'api.Put("/users/me/emails/:id", handlers.UpdateEmail)' not in content:
    content = content.replace('api.Delete("/users/me/emails/:id", handlers.DeleteEmail)', 'api.Delete("/users/me/emails/:id", handlers.DeleteEmail)\n\tapi.Put("/users/me/emails/:id", handlers.UpdateEmail)')

if 'api.Put("/users/me/phones/:id", handlers.UpdatePhone)' not in content:
    content = content.replace('api.Delete("/users/me/phones/:id", handlers.DeletePhone)', 'api.Delete("/users/me/phones/:id", handlers.DeletePhone)\n\tapi.Put("/users/me/phones/:id", handlers.UpdatePhone)')

with open('apps/api/main.go', 'w') as f:
    f.write(content)


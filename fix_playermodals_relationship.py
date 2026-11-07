import re

with open('apps/app/src/modules/account/components/modals/PlayerModals.tsx', 'r') as f:
    content = f.read()

# I already added the state `relationship`? Let's check.
# Wait, did I? In fix_forms.py I added graduationYear. Let's add relationship.
if 'const [relationship, setRelationship] = useState("");' not in content:
    content = content.replace('const [graduationYear, setGraduationYear] = useState("");', 'const [graduationYear, setGraduationYear] = useState("");\n  const [relationship, setRelationship] = useState("");')

# Add relationship to reset
if 'setRelationship("");' not in content:
    content = content.replace('setGraduationYear("");', 'setGraduationYear("");\n      setRelationship("");')

# Add to fetch payload
old_payload = """          email,
          phone: normalizePhoneForAPI(phone),
        }),"""
new_payload = """          email,
          phone: normalizePhoneForAPI(phone),
          relationship,
        }),"""
content = content.replace(old_payload, new_payload)

# Add to UI
if 'label="Your Relationship' not in content:
    # Insert it right before "Email Address"
    content = content.replace('<FormField label="Email Address"', '<FormField label="Your Relationship (e.g. Mother, Father)" type="text" value={relationship} onChange={(e) => setRelationship(e.target.value)} />\n        <FormField label="Email Address"')

with open('apps/app/src/modules/account/components/modals/PlayerModals.tsx', 'w') as f:
    f.write(content)


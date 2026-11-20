import re

with open('apps/app/src/modules/account/components/modals/GuardianModals.tsx', 'r') as f:
    content = f.read()

# Add formatPhone import
content = content.replace('import { getDateOfBirthMax } from "@lh/utils";', 'import { getDateOfBirthMax, formatPhone, normalizePhoneForAPI } from "@lh/utils";')

# Update phone field
old_phone = '<FormField label="Phone Number" type="tel" value={phone} onChange={(e) => setPhone(e.target.value)} />'
new_phone = '<FormField label="Phone Number" type="tel" value={phone} onChange={(e) => setPhone(formatPhone(e.target.value))} />'
content = content.replace(old_phone, new_phone)

# Update payload phone
old_payload = '          phone,\n        }),'
new_payload = '          phone: normalizePhoneForAPI(phone),\n        }),'
content = content.replace(old_payload, new_payload)

with open('apps/app/src/modules/account/components/modals/GuardianModals.tsx', 'w') as f:
    f.write(content)


with open('apps/app/src/modules/account/components/modals/PlayerModals.tsx', 'r') as f:
    content = f.read()

content = content.replace('import { getDateOfBirthMax } from "@lh/utils";', 'import { getDateOfBirthMax, formatPhone, normalizePhoneForAPI } from "@lh/utils";')

old_phone2 = '<FormField label="Phone Number" type="tel" placeholder="Optional" value={phone} onChange={(e) => setPhone(e.target.value)} />'
new_phone2 = '<FormField label="Phone Number" type="tel" placeholder="Optional" value={phone} onChange={(e) => setPhone(formatPhone(e.target.value))} />'
content = content.replace(old_phone2, new_phone2)

old_payload2 = '          email,\n          phone,\n        }),'
new_payload2 = '          email,\n          phone: normalizePhoneForAPI(phone),\n        }),'
content = content.replace(old_payload2, new_payload2)

# Add graduationYear to AddPlayerModal
old_state = '  const [dob, setDob] = useState("");\n  const [_isSubmitting, setIsSubmitting] = useState(false);'
new_state = '  const [dob, setDob] = useState("");\n  const [graduationYear, setGraduationYear] = useState("");\n  const [_isSubmitting, setIsSubmitting] = useState(false);'
content = content.replace(old_state, new_state)

old_reset = '      setDob("");\n      setEmail("");\n      setPhone("");'
new_reset = '      setDob("");\n      setGraduationYear("");\n      setEmail("");\n      setPhone("");'
content = content.replace(old_reset, new_reset)

# Remove hardcoded 2030 and replace with graduationYear variable
content = content.replace('graduation_year: "2030", // Assuming DOB -> Grad Year logic later', 'graduation_year: graduationYear,')

# Add UI for graduation year
old_ui = '        <FormField label="Date of Birth" type="date" max={getDateOfBirthMax(4)} value={dob} onChange={(e) => setDob(e.target.value)} />'
new_ui = '        <div className="grid grid-cols-2 gap-4">\n          <FormField label="Date of Birth" type="date" max={getDateOfBirthMax(4)} value={dob} onChange={(e) => setDob(e.target.value)} />\n          <FormField label="Graduation Year" type="text" placeholder="e.g. 2030" value={graduationYear} onChange={(e) => setGraduationYear(e.target.value)} />\n        </div>'
content = content.replace(old_ui, new_ui)


with open('apps/app/src/modules/account/components/modals/PlayerModals.tsx', 'w') as f:
    f.write(content)


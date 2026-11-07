import re

with open('apps/app/src/modules/account/components/modals/GuardianModals.tsx', 'r') as f:
    content = f.read()
content = content.replace('getDateOfBirthMax, ', '')
with open('apps/app/src/modules/account/components/modals/GuardianModals.tsx', 'w') as f:
    f.write(content)


with open('apps/app/src/modules/account/components/modals/PlayerModals.tsx', 'r') as f:
    content = f.read()

if 'const [graduationYear' not in content:
    content = content.replace('  const [dob, setDob] = useState("");', '  const [dob, setDob] = useState("");\n  const [graduationYear, setGraduationYear] = useState("");\n  const [relationship, setRelationship] = useState("");')

with open('apps/app/src/modules/account/components/modals/PlayerModals.tsx', 'w') as f:
    f.write(content)


with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()
# Let's just remove session if it's unused in HouseholdTab. But wait, it's used in PersonalInfoTab too.
# The error was: src/modules/account/pages/AccountPage.tsx:1000:9
# Line 1000 is probably inside HouseholdTab.
match = re.search(r'function HouseholdTab\(\{ household, fetchHousehold, loading \}: \{ household: any; fetchHousehold: \(\) => void; loading: boolean \}\) \{\s*const \{ session \} = useAuth\(\);', content)
if match:
    content = content.replace('const { session } = useAuth();', '')
    
with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


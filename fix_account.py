import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# M10: PersonalInfoTab primaryUser
old_m10 = """const primaryUser = household?.users?.find((u: any) => 
    u.emails?.some((e: any) => e.email === session?.user?.email)
  ) || household?.users?.[0];"""
new_m10 = """const primaryUser = household?.users?.find((u: any) => 
    u.id === session?.user?.id
  ) || household?.users?.[0];"""
content = content.replace(old_m10, new_m10)

# M11: HouseholdTab isMe
old_m11 = """            // Check if this is the currently logged in user (via email match)
            const isMe = emailStr === session?.user?.email;"""
new_m11 = """            // Check if this is the currently logged in user (via user id)
            const isMe = user.id === session?.user?.id;"""
content = content.replace(old_m11, new_m11)

# M15: handleSetPrimaryEmail
old_m15 = """  const handleSetPrimaryEmail = async (id: string) => {
    await fetch(`http://localhost:3000/api/users/me/emails/${id}/set-primary`, {
      method: "PUT",
      headers: authHeader,
    });
    setEmails((prev) => prev.map((e) => ({ ...e, is_primary: e.id === id })));
  };"""
new_m15 = """  const handleSetPrimaryEmail = async (id: string) => {
    const res = await fetch(`http://localhost:3000/api/users/me/emails/${id}/set-primary`, {
      method: "PUT",
      headers: authHeader,
    });
    if (res.ok) {
      setEmails((prev) => prev.map((e) => ({ ...e, is_primary: e.id === id })));
    }
  };"""
content = content.replace(old_m15, new_m15)

# M16: handleSetPrimaryPhone
old_m16 = """  const handleSetPrimaryPhone = async (id: string) => {
    await fetch(`http://localhost:3000/api/users/me/phones/${id}/set-primary`, {
      method: "PUT",
      headers: authHeader,
    });
    setPhones((prev) => prev.map((p) => ({ ...p, is_primary: p.id === id })));
  };"""
new_m16 = """  const handleSetPrimaryPhone = async (id: string) => {
    const res = await fetch(`http://localhost:3000/api/users/me/phones/${id}/set-primary`, {
      method: "PUT",
      headers: authHeader,
    });
    if (res.ok) {
      setPhones((prev) => prev.map((p) => ({ ...p, is_primary: p.id === id })));
    }
  };"""
content = content.replace(old_m16, new_m16)

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


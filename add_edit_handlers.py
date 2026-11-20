import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Add state
if 'const [editingEmailId' not in content:
    content = content.replace('  const [emails, setEmails] = useState<UserEmail[]>([]);', '  const [emails, setEmails] = useState<UserEmail[]>([]);\n  const [editingEmailId, setEditingEmailId] = useState<string | null>(null);\n  const [editEmailValue, setEditEmailValue] = useState("");')
    content = content.replace('  const [phones, setPhones] = useState<UserPhone[]>([]);', '  const [phones, setPhones] = useState<UserPhone[]>([]);\n  const [editingPhoneId, setEditingPhoneId] = useState<string | null>(null);\n  const [editPhoneValue, setEditPhoneValue] = useState("");')

# Add handlers
handlers = """
  const handleSaveEmailEdit = async (id: string) => {
    if (!editEmailValue.trim()) return;
    try {
      const res = await fetch(`http://localhost:3000/api/users/me/emails/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json", ...authHeader },
        body: JSON.stringify({ value: editEmailValue })
      });
      if (res.ok) {
        setEditingEmailId(null);
        fetchContacts();
      }
    } catch (err) { console.error(err); }
  };

  const handleSavePhoneEdit = async (id: string) => {
    if (!editPhoneValue.trim()) return;
    try {
      const res = await fetch(`http://localhost:3000/api/users/me/phones/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json", ...authHeader },
        body: JSON.stringify({ value: editPhoneValue })
      });
      if (res.ok) {
        setEditingPhoneId(null);
        fetchContacts();
      }
    } catch (err) { console.error(err); }
  };
"""

if 'handleSaveEmailEdit' not in content:
    content = content.replace('  const handleDeleteEmail = async', handlers + '\n  const handleDeleteEmail = async')

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


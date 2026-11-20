import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Add state
if 'const [editingEmailId, setEditingEmailId]' not in content:
    content = content.replace('  // Contact list state', '  // Contact list state\n  const [editingEmailId, setEditingEmailId] = useState<string | null>(null);\n  const [editEmailValue, setEditEmailValue] = useState("");\n  const [editingPhoneId, setEditingPhoneId] = useState<string | null>(null);\n  const [editPhoneValue, setEditPhoneValue] = useState("");')

# Add handleSaveEdit
if 'const handleSaveEmailEdit' not in content:
    content = content.replace('  const handleDeleteEmail = async (id: string) => {', """
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

  const handleDeleteEmail = async (id: string) => {""")

# Update rendering of emails
old_email_render = """              <ListCard key={emailItem.id}>
                <div>
                  <p className="font-semibold text-white">{emailItem.email}</p>
                  <p className="text-xs text-white/50">{emailItem.is_primary ? "Primary Email" : "Secondary Email"}</p>
                </div>"""
new_email_render = """              <ListCard key={emailItem.id}>
                {editingEmailId === emailItem.id ? (
                  <div className="flex-1 flex gap-2">
                    <Input 
                      className="flex-1" 
                      value={editEmailValue} 
                      onChange={(e) => setEditEmailValue(e.target.value)} 
                      autoFocus 
                      onKeyDown={(e) => { if (e.key === 'Enter') handleSaveEmailEdit(emailItem.id); else if (e.key === 'Escape') setEditingEmailId(null); }}
                    />
                    <Button variant="outline" size="sm" onClick={() => handleSaveEmailEdit(emailItem.id)}>Save</Button>
                    <Button variant="ghost" size="sm" onClick={() => setEditingEmailId(null)}>Cancel</Button>
                  </div>
                ) : (
                <>
                <div>
                  <p className="font-semibold text-white">{emailItem.email}</p>
                  <p className="text-xs text-white/50">{emailItem.is_primary ? "Primary Email" : "Secondary Email"}</p>
                </div>"""
content = content.replace(old_email_render, new_email_render)

# Close fragment and add Edit button
old_email_buttons = """                  <Button
                    variant="ghost"
                    className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                    onClick={() => handleDeleteEmail(emailItem.id)}
                  >
                    Remove
                  </Button>"""
new_email_buttons = """                  {!emailItem.source_provider || emailItem.source_provider === 'manual' ? (
                    <Button
                      variant="ghost"
                      className="text-white/40 hover:text-white hover:bg-white/10 text-xs h-8 px-3"
                      onClick={() => {
                        setEditingEmailId(emailItem.id);
                        setEditEmailValue(emailItem.email);
                      }}
                    >
                      Edit
                    </Button>
                  ) : null}
                  <Button
                    variant="ghost"
                    className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                    onClick={() => handleDeleteEmail(emailItem.id)}
                  >
                    Remove
                  </Button>"""
content = content.replace(old_email_buttons, new_email_buttons)
# Make sure to close the fragment for emails
content = content.replace("""              </ListCard>
            );
          })}""", """                </>
                )}
              </ListCard>
            );
          })}""")

# Same for phones
old_phone_render = """              <ListCard key={phoneItem.id}>
                <div>
                  <p className="font-semibold text-white">{formatPhone(phoneItem.phone)}</p>
                  <p className="text-xs text-white/50">{phoneItem.is_primary ? "Primary Phone" : "Secondary Phone"}</p>
                </div>"""
new_phone_render = """              <ListCard key={phoneItem.id}>
                {editingPhoneId === phoneItem.id ? (
                  <div className="flex-1 flex gap-2">
                    <Input 
                      className="flex-1" 
                      value={editPhoneValue} 
                      onChange={(e) => setEditPhoneValue(formatPhone(e.target.value))} 
                      prefixText="+1"
                      autoFocus 
                      onKeyDown={(e) => { if (e.key === 'Enter') handleSavePhoneEdit(phoneItem.id); else if (e.key === 'Escape') setEditingPhoneId(null); }}
                    />
                    <Button variant="outline" size="sm" onClick={() => handleSavePhoneEdit(phoneItem.id)}>Save</Button>
                    <Button variant="ghost" size="sm" onClick={() => setEditingPhoneId(null)}>Cancel</Button>
                  </div>
                ) : (
                <>
                <div>
                  <p className="font-semibold text-white">{formatPhone(phoneItem.phone)}</p>
                  <p className="text-xs text-white/50">{phoneItem.is_primary ? "Primary Phone" : "Secondary Phone"}</p>
                </div>"""
content = content.replace(old_phone_render, new_phone_render)

old_phone_buttons = """                  <Button
                    variant="ghost"
                    className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                    onClick={() => handleDeletePhone(phoneItem.id)}
                  >
                    Remove
                  </Button>"""
new_phone_buttons = """                  <Button
                    variant="ghost"
                    className="text-white/40 hover:text-white hover:bg-white/10 text-xs h-8 px-3"
                    onClick={() => {
                      setEditingPhoneId(phoneItem.id);
                      setEditPhoneValue(phoneItem.phone);
                    }}
                  >
                    Edit
                  </Button>
                  <Button
                    variant="ghost"
                    className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                    onClick={() => handleDeletePhone(phoneItem.id)}
                  >
                    Remove
                  </Button>"""
content = content.replace(old_phone_buttons, new_phone_buttons)

# Also close the fragment for phones! But wait, my previous replacement for emails replaced the closing tag.
# I need to specifically target the closing tag inside the phone map.
content = re.sub(r'(\s*)</ListCard>\n(\s*)\);\n(\s*)}\)', r'\1</>\n\1  )}\n\1</ListCard>\n\2);\n\3})', content)

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


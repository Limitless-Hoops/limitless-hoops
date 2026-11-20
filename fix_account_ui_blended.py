import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Replace Guardians section
old_guardians = """      {/* Additional Contacts */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <div>
              <h2 className="text-xl font-bold text-white mb-1">Guardians</h2>
              <p className="text-sm text-white/50">Manage parents/guardians for this household.</p>
            </div>
            <Button variant="outline" onClick={() => setIsAddGuardianOpen(true)} className="border-white/20 hover:bg-white/10">+ Add Guardian</Button>
          </div>

          {household?.users?.map((user: any) => {
            // Note: DB returns email via the join so user.emails array isn't needed here if we rely on the primary email field from GetMyHousehold (wait, user_emails are nested now. Let's handle both)
            const emailStr = user.emails?.find((e: any) => e.is_primary)?.email || user.email;
            const phoneStr = formatPhone(user.phones?.find((p: any) => p.is_primary)?.phone || user.phone);
            
            // Check if this is the currently logged in user (via user id)
            const isMe = user.id === session?.user?.id;

            return (
              <ListCard key={user.id}>
                <div>
                  <div className="flex items-center gap-2">
                    <p className="font-semibold text-white">{user.first_name} {user.last_name}</p>
                    {isMe && (
                      <span className="text-[10px] bg-white/15 text-white font-bold px-1.5 py-0.5 rounded tracking-wide">
                        YOU
                      </span>
                    )}
                  </div>
                  <p className="text-xs text-white/50 mt-0.5">{emailStr} &bull; {phoneStr}</p>
                </div>
                {isMe ? (
                  <CustomTooltip content="You cannot remove yourself from the household." align="right">
                    <span>
                      <Button variant="ghost" disabled className="text-white/40 text-xs h-8 px-3 disabled:opacity-30 disabled:cursor-not-allowed">Remove</Button>
                    </span>
                  </CustomTooltip>
                ) : (
                  <Button variant="ghost" onClick={() => setGuardianToRemove(`${user.first_name} ${user.last_name}`)} className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3">Remove</Button>
                )}
              </ListCard>
            );
          })}
        </CardContent>
      </Card>"""

new_guardians = """      {/* Primary Account & Notification Contacts */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <div>
              <h2 className="text-xl font-bold text-white mb-1">Notification Contacts</h2>
              <p className="text-sm text-white/50">Manage who receives updates for your dependents.</p>
            </div>
            <Button variant="outline" onClick={() => setIsAddGuardianOpen(true)} className="border-white/20 hover:bg-white/10">+ Add Contact</Button>
          </div>

          {/* Primary Owner Card */}
          {household?.owner && (
            <ListCard>
              <div>
                <div className="flex items-center gap-2">
                  <p className="font-semibold text-white">{household.owner.first_name} {household.owner.last_name}</p>
                  <span className="text-[10px] bg-white/15 text-white font-bold px-1.5 py-0.5 rounded tracking-wide">
                    ACCOUNT OWNER
                  </span>
                </div>
                <p className="text-xs text-white/50 mt-0.5">
                  {household.owner.emails?.find((e: any) => e.is_primary)?.email || household.owner.email} &bull; {formatPhone(household.owner.phones?.find((p: any) => p.is_primary)?.phone || household.owner.phone)}
                </p>
              </div>
              <CustomTooltip content="You are the primary account owner." align="right">
                <span>
                  <Button variant="ghost" disabled className="text-white/40 text-xs h-8 px-3 disabled:opacity-30 disabled:cursor-not-allowed">Remove</Button>
                </span>
              </CustomTooltip>
            </ListCard>
          )}

          {/* Secondary Contacts */}
          {household?.players?.flatMap((p: any) => p.contacts?.map((c: any) => ({ ...c, playerName: p.first_name })) || []).map((contact: any) => (
            <ListCard key={contact.id}>
              <div>
                <div className="flex items-center gap-2">
                  <p className="font-semibold text-white">{contact.first_name} {contact.last_name}</p>
                  <span className="text-[10px] bg-white/10 text-white/70 px-1.5 py-0.5 rounded tracking-wide border border-white/10">
                    {contact.relationship}
                  </span>
                </div>
                <p className="text-xs text-white/50 mt-0.5">
                  {contact.email} &bull; {formatPhone(contact.phone)}
                </p>
                <p className="text-xs text-emerald-400 mt-1">Receives updates for {contact.playerName}</p>
              </div>
              <Button variant="ghost" onClick={() => setGuardianToRemove(contact)} className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3">Remove</Button>
            </ListCard>
          ))}
        </CardContent>
      </Card>"""

content = content.replace(old_guardians, new_guardians)

# Fix AddGuardianModal passing household to pass players
content = content.replace('<AddGuardianModal isOpen={isAddGuardianOpen} onClose={() => { setIsAddGuardianOpen(false); fetchHousehold(); }} householdId={household?.id} />', 
                          '<AddGuardianModal isOpen={isAddGuardianOpen} onClose={() => { setIsAddGuardianOpen(false); fetchHousehold(); }} players={household?.players || []} />')

# Fix RemoveGuardianModal passing contact object
content = content.replace('setGuardianToRemove(`${user.first_name} ${user.last_name}`)', 'setGuardianToRemove(contact)')
content = content.replace('guardianName={guardianToRemove || ""}', 'contact={guardianToRemove}')
content = content.replace('const [guardianToRemove, setGuardianToRemove] = useState<string | null>(null);', 'const [guardianToRemove, setGuardianToRemove] = useState<any>(null);')

# Also PersonalInfoTab has primaryUser calculation which depended on household.users
old_personal = """  const primaryUser = household?.users?.find((u: any) => 
    u.id === session?.user?.id
  ) || household?.users?.[0];"""
new_personal = """  const primaryUser = household?.owner;"""
content = content.replace(old_personal, new_personal)

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


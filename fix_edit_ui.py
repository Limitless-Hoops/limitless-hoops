import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Replace emails ListCard content
old_email_render = """                      <div className="flex flex-wrap items-center gap-2">
                        <span className="text-sm text-white font-medium">{emailItem.email}</span>"""
new_email_render = """                      <div className="flex flex-wrap items-center gap-2 flex-1">
                        {editingEmailId === emailItem.id ? (
                          <div className="flex items-center gap-2 w-full max-w-sm">
                            <Input 
                              value={editEmailValue} 
                              onChange={(e) => setEditEmailValue(e.target.value)}
                              className="h-8 text-sm px-2 bg-black/40 border-white/20 text-white"
                              autoFocus
                              onKeyDown={(e) => { if (e.key === 'Enter') handleSaveEmailEdit(emailItem.id); else if (e.key === 'Escape') setEditingEmailId(null); }}
                            />
                            <Button size="sm" className="h-8" onClick={() => handleSaveEmailEdit(emailItem.id)}>Save</Button>
                            <Button variant="ghost" size="sm" className="h-8 text-white/50 hover:text-white" onClick={() => setEditingEmailId(null)}>Cancel</Button>
                          </div>
                        ) : (
                          <span className="text-sm text-white font-medium">{emailItem.email}</span>
                        )}"""

content = content.replace(old_email_render, new_email_render)

old_email_remove = """                        ) : (
                          <Button
                            variant="ghost"
                            onClick={() => handleRemoveEmail(emailItem.id)}
                            className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                          >
                            Remove
                          </Button>
                        )}
                      </div>"""
new_email_remove = """                        ) : (
                          <>
                            {(!emailItem.source_provider || emailItem.source_provider === 'manual') && editingEmailId !== emailItem.id && (
                              <Button
                                variant="ghost"
                                onClick={() => { setEditingEmailId(emailItem.id); setEditEmailValue(emailItem.email); }}
                                className="text-white/40 hover:text-white hover:bg-white/10 text-xs h-8 px-3"
                              >
                                Edit
                              </Button>
                            )}
                            <Button
                              variant="ghost"
                              onClick={() => handleRemoveEmail(emailItem.id)}
                              className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                            >
                              Remove
                            </Button>
                          </>
                        )}
                      </div>"""

content = content.replace(old_email_remove, new_email_remove)


# Replace phones ListCard content
old_phone_render = """                      <div className="flex flex-wrap items-center gap-2">
                        <span className="text-sm text-white font-medium">{formatPhone(phoneItem.phone)}</span>"""

new_phone_render = """                      <div className="flex flex-wrap items-center gap-2 flex-1">
                        {editingPhoneId === phoneItem.id ? (
                          <div className="flex items-center gap-2 w-full max-w-sm">
                            <Input 
                              value={editPhoneValue} 
                              onChange={(e) => setEditPhoneValue(formatPhone(e.target.value))}
                              prefixText="+1"
                              className="h-8 text-sm px-2 bg-black/40 border-white/20 text-white"
                              autoFocus
                              onKeyDown={(e) => { if (e.key === 'Enter') handleSavePhoneEdit(phoneItem.id); else if (e.key === 'Escape') setEditingPhoneId(null); }}
                            />
                            <Button size="sm" className="h-8" onClick={() => handleSavePhoneEdit(phoneItem.id)}>Save</Button>
                            <Button variant="ghost" size="sm" className="h-8 text-white/50 hover:text-white" onClick={() => setEditingPhoneId(null)}>Cancel</Button>
                          </div>
                        ) : (
                          <span className="text-sm text-white font-medium">{formatPhone(phoneItem.phone)}</span>
                        )}"""

content = content.replace(old_phone_render, new_phone_render)

old_phone_remove = """                        ) : (
                          <Button
                            variant="ghost"
                            onClick={() => handleRemovePhone(phoneItem.id)}
                            className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                          >
                            Remove
                          </Button>
                        )}
                      </div>"""

new_phone_remove = """                        ) : (
                          <>
                            {editingPhoneId !== phoneItem.id && (
                              <Button
                                variant="ghost"
                                onClick={() => { setEditingPhoneId(phoneItem.id); setEditPhoneValue(phoneItem.phone); }}
                                className="text-white/40 hover:text-white hover:bg-white/10 text-xs h-8 px-3"
                              >
                                Edit
                              </Button>
                            )}
                            <Button
                              variant="ghost"
                              onClick={() => handleRemovePhone(phoneItem.id)}
                              className="text-white/40 hover:text-red-400 hover:bg-red-900/20 text-xs h-8 px-3"
                            >
                              Remove
                            </Button>
                          </>
                        )}
                      </div>"""

content = content.replace(old_phone_remove, new_phone_remove)

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)


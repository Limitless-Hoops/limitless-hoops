import { useState, useEffect } from "react";
import DashboardLayout from "@/modules/dashboard/components/DashboardLayout";
import { Card, CardContent, Input, Button, FormField, CustomTooltip, ListCard } from "@lh/ui";
import { formatPhone, normalizePhoneForAPI, isValidName, getDateOfBirthMax, isValidEmail, isValidPhone } from "@lh/utils";
import { useSearchParams } from "react-router-dom";
import { supabase } from "@/lib/supabase";
import { useAuth } from "@/modules/auth/contexts/AuthContext";

import { DisconnectSocialModal } from "@/modules/account/components/modals/DisconnectSocialModal";
import { VerifyContactModal } from "@/modules/account/components/modals/VerifyContactModal";
import { AddGuardianModal, RemoveGuardianModal } from "@/modules/account/components/modals/GuardianModals";
import { AddPlayerModal, RemovePlayerModal } from "@/modules/account/components/modals/PlayerModals";
import { DeleteAccountModal } from "@/modules/account/components/modals/DeleteAccountModal";
import { StripeMockModal } from "@/modules/account/components/modals/StripeMockModal";

// ─── Constants ───────────────────────────────────────────────────────────────

const TABS = [
  { id: "personal", label: "Personal Information" },
  { id: "security", label: "Login & Security" },
  { id: "household", label: "Household & Contacts" },
  { id: "notifications", label: "Notifications" },
  { id: "billing", label: "Subscription & Billing" },
  { id: "advanced", label: "Advanced" },
];

/** Shared card class applied to every section card */
const cardClass = "border-white/10 bg-black/40 backdrop-blur-2xl shadow-xl shadow-black/50";

// ─── Types ────────────────────────────────────────────────────────────────────

interface UserEmail {
  id: string;
  email: string;
  is_primary: boolean;
  is_verified: boolean;
  source_provider: string; // 'manual' | 'google' | 'facebook'
  label?: string;
}

interface UserPhone {
  id: string;
  phone: string;
  is_primary: boolean;
  is_verified: boolean;
  label?: string;
}

interface NotificationPreference {
  id: string;
  target_type: "email" | "phone";
  target_id: string;
  category: "billing" | "events" | "announcements" | "security";
  enabled: boolean;
  // Joined fields from backend
  contact_value?: string;
  is_primary?: boolean;
}

// ─── Tabs ─────────────────────────────────────────────────────────────────────

function PersonalInfoTab({ household }: { household: any }) {
  const { session, refreshHousehold } = useAuth();
  const primaryUser = household?.owner;

  const [firstName, setFirstName] = useState(primaryUser?.first_name || "");
  const [lastName, setLastName] = useState(primaryUser?.last_name || "");
  const [dob, setDob] = useState(primaryUser?.dob || "");
  const [isSaving, setIsSaving] = useState(false);
  const [saveSuccess, setSaveSuccess] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");
  
  useEffect(() => {
    if (primaryUser) {
      setFirstName(primaryUser.first_name || "");
      setLastName(primaryUser.last_name || "");
      setDob(primaryUser.dob || "");
    }
  }, [primaryUser]);

  const firstNameError = firstName && !isValidName(firstName) ? "Name can only contain letters, spaces, hyphens, and apostrophes." : undefined;
  const lastNameError = lastName && !isValidName(lastName) ? "Name can only contain letters, spaces, hyphens, and apostrophes." : undefined;

  const handleSave = async () => {
    if (firstNameError || lastNameError || !firstName || !lastName || !dob) {
      return;
    }
    
    setIsSaving(true);
    setErrorMsg("");
    setSaveSuccess(false);

    try {
      const res = await fetch("http://localhost:3000/api/users/me", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${session?.access_token}`
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          dob: dob
        })
      });

      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Failed to save changes");
      }

      setSaveSuccess(true);
      await refreshHousehold();
      
      // Hide success message after 3 seconds
      setTimeout(() => setSaveSuccess(false), 3000);
    } catch (err: any) {
      setErrorMsg(err.message);
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <Card className={cardClass}>
      <CardContent className="p-6 sm:p-8 space-y-6">
        <div>
          <h2 className="text-xl font-bold text-white mb-1">Personal Information</h2>
          <p className="text-sm text-white/50">Manage your basic profile details.</p>
        </div>

        {errorMsg && (
          <div className="bg-red-900/20 border border-red-500/50 rounded-lg p-3">
            <p className="text-sm text-red-400">{errorMsg}</p>
          </div>
        )}
        
        {saveSuccess && (
          <div className="bg-green-900/20 border border-green-500/50 rounded-lg p-3">
            <p className="text-sm text-green-400">✅ Changes saved successfully!</p>
          </div>
        )}

        <div className="grid gap-6 sm:grid-cols-2">
          <FormField
            label="First Name"
            value={firstName}
            error={firstNameError}
            onChange={(e) => setFirstName(e.target.value)}
            className="bg-white/5 border-white/10 text-white"
          />
          <FormField
            label="Last Name"
            value={lastName}
            error={lastNameError}
            onChange={(e) => setLastName(e.target.value)}
            className="bg-white/5 border-white/10 text-white"
          />
        </div>

        <FormField
          label="Date of Birth"
          type="date"
          value={dob}
          max={getDateOfBirthMax(18)}
          onChange={(e) => setDob(e.target.value)}
          hint="Used to verify you are 18 or older to manage your dependents."
          className="bg-white/5 border-white/10 text-white w-full sm:w-1/2"
        />

        <div className="pt-4 flex justify-end">
          <Button 
            onClick={handleSave} 
            disabled={isSaving || !!firstNameError || !!lastNameError || !firstName || !lastName || !dob}
          >
            {isSaving ? "Saving..." : "Save Changes"}
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}

// ─── SecurityTab ──────────────────────────────────────────────────────────────

function SecurityTab({ household: _household }: { household: any }) {
  const { session } = useAuth();
  // Contact list state

  const [emails, setEmails] = useState<UserEmail[]>([]);
  const [editingEmailId, setEditingEmailId] = useState<string | null>(null);
  const [editEmailValue, setEditEmailValue] = useState("");
  const [phones, setPhones] = useState<UserPhone[]>([]);
  const [editingPhoneId, setEditingPhoneId] = useState<string | null>(null);
  const [editPhoneValue, setEditPhoneValue] = useState("");
  const [contactsLoading, setContactsLoading] = useState(false);

  // Add-email form state
  const [showAddEmail, setShowAddEmail] = useState(false);
  const [newEmail, setNewEmail] = useState("");
  const [addEmailLoading, setAddEmailLoading] = useState(false);
  const [addEmailError, setAddEmailError] = useState("");

  // Add-phone form state
  const [showAddPhone, setShowAddPhone] = useState(false);
  const [newPhone, setNewPhone] = useState("");
  const [addPhoneLoading, setAddPhoneLoading] = useState(false);
  const [addPhoneError, setAddPhoneError] = useState("");

  // Disconnect modal state
  const [isDisconnectOpen, setIsDisconnectOpen] = useState(false);
  const [disconnectProvider, setDisconnectProvider] = useState("");

  // Verify modal state
  const [verifyContact, setVerifyContact] = useState<{ id: string, value: string, type: "emails" | "phones" } | null>(null);

  const authHeader: Record<string, string> = session?.access_token
    ? { Authorization: `Bearer ${session.access_token}` }
    : {};


  // Real identities from Supabase Auth
  const [identities, setIdentities] = useState<any[]>([]);

  // ── Fetch contacts ──
  const fetchContacts = async () => {
    if (!session?.access_token) return;
    try {
      // 1. Fetch from our API
      const [emailRes, phoneRes] = await Promise.all([
        fetch("http://localhost:3000/api/users/me/emails", { headers: authHeader }),
        fetch("http://localhost:3000/api/users/me/phones", { headers: authHeader }),
      ]);
      if (emailRes.ok) setEmails(await emailRes.json());
      if (phoneRes.ok) setPhones(await phoneRes.json());

      // 2. Fetch fresh user data from Supabase to get up-to-date identities
      const { data: { user } } = await supabase.auth.getUser();
      if (user?.identities) {
        setIdentities(user.identities);
      }
    } catch (err) {
      console.error("Failed to fetch contacts:", err);
    }
  };

  useEffect(() => {
    const init = async () => {
      setContactsLoading(true);
      await fetchContacts();
      setContactsLoading(false);
    };
    init();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [session?.access_token]);

  // ── Email actions ──
  const handleSetPrimaryEmail = async (id: string) => {
    const res = await fetch(`http://localhost:3000/api/users/me/emails/${id}/set-primary`, {
      method: "PUT",
      headers: authHeader,
    });
    if (res.ok) {
      setEmails((prev) => prev.map((e) => ({ ...e, is_primary: e.id === id })));
    }
  };

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
  const handleRemoveEmail = async (id: string) => {
    const res = await fetch(`http://localhost:3000/api/users/me/emails/${id}`, {
      method: "DELETE",
      headers: authHeader,
    });
    if (res.ok) setEmails((prev) => prev.filter((e) => e.id !== id));
  };

  const handleAddEmail = async () => {
    if (!newEmail) return;
    setAddEmailLoading(true);
    setAddEmailError("");
    try {
      const res = await fetch("http://localhost:3000/api/users/me/emails", {
        method: "POST",
        headers: { ...authHeader, "Content-Type": "application/json" },
        body: JSON.stringify({ email: newEmail }),
      });
      if (res.ok) {
        const data = await res.json();
        setShowAddEmail(false);
        await fetchContacts(); // Refresh list immediately
        
        // Auto-open modal for the newly added email
        if (data.email?.id) {
          setVerifyContact({ id: data.email.id, value: newEmail, type: "emails" });
        }
        setNewEmail("");
      } else {
        const data = await res.json();
        setAddEmailError(data.error || "Failed to add email.");
      }
    } catch (err) {
      console.error("Failed to add email:", err);
      setAddEmailError("A network error occurred. Please try again.");
    } finally {
      setAddEmailLoading(false);
    }
  };

  // ── Phone actions ──
  const handleSetPrimaryPhone = async (id: string) => {
    const res = await fetch(`http://localhost:3000/api/users/me/phones/${id}/set-primary`, {
      method: "PUT",
      headers: authHeader,
    });
    if (res.ok) {
      setPhones((prev) => prev.map((p) => ({ ...p, is_primary: p.id === id })));
    }
  };

  const handleRemovePhone = async (id: string) => {
    const res = await fetch(`http://localhost:3000/api/users/me/phones/${id}`, {
      method: "DELETE",
      headers: authHeader,
    });
    if (res.ok) setPhones((prev) => prev.filter((p) => p.id !== id));
  };

  const handleAddPhone = async () => {
    if (!newPhone) return;
    setAddPhoneLoading(true);
    setAddPhoneError("");
    try {
      const res = await fetch("http://localhost:3000/api/users/me/phones", {
        method: "POST",
        headers: { ...authHeader, "Content-Type": "application/json" },
        body: JSON.stringify({ phone: normalizePhoneForAPI(newPhone) }),
      });
      if (res.ok) {
        const data = await res.json();
        setShowAddPhone(false);
        await fetchContacts(); // Refresh list immediately
        
        // Auto-open modal for the newly added phone
        if (data.phone?.id) {
          setVerifyContact({ id: data.phone.id, value: newPhone, type: "phones" });
        }
        setNewPhone("");
      } else {
        const data = await res.json();
        setAddPhoneError(data.error || "Failed to add phone.");
      }
    } catch (err) {
      console.error("Failed to add phone:", err);
      setAddPhoneError("A network error occurred. Please try again.");
    } finally {
      setAddPhoneLoading(false);
    }
  };

  // ── Connected Accounts helpers ──


  // Disconnect is only safe when there is more than one active login method
  // We count verified manual emails and distinct provider identities

  return (
    <Card className={cardClass}>
      <CardContent className="p-6 sm:p-8 space-y-8">
        <div>
          <h2 className="text-xl font-bold text-white mb-1">Login & Security</h2>
          <p className="text-sm text-white/50">Manage your email addresses, phone numbers, and connected accounts.</p>
        </div>

        {/* ── Email Addresses ── */}
        <div className="space-y-4">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-white/80 border-b border-white/10 pb-2">
            Email Addresses
          </h3>

          {contactsLoading ? (
            <p className="text-sm text-white/40">Loading...</p>
          ) : (
            <div className="flex flex-col gap-2">
              {emails.map((emailItem) => {
                const isProvider = emailItem.source_provider !== "manual";
                const isAuthIdentity = emailItem.email === session?.user?.email;
                const canRemove =
                  !emailItem.is_primary &&
                  emails.length > 1 &&
                  !isProvider &&
                  !isAuthIdentity;
                const canSetPrimary = !emailItem.is_primary && emailItem.is_verified;

                // Build a human-readable reason why Remove is disabled
                const removeDisabledReason = emailItem.is_primary
                  ? "This is your primary email. Set another verified email as primary before removing this one."
                  : emails.length <= 1
                  ? "You must have at least one email on your account."
                  : isAuthIdentity
                  ? "You cannot remove the email you are currently signed in with."
                  : isProvider
                  ? `Provider-linked emails can only be removed by disconnecting your ${emailItem.source_provider.charAt(0).toUpperCase() + emailItem.source_provider.slice(1)} account.`
                  : undefined;

                return (
                  <ListCard key={emailItem.id} className="flex-col sm:flex-col items-stretch sm:items-stretch">
                    <div className="flex flex-col sm:flex-row sm:items-center justify-between w-full gap-3">
                      {/* Left: address + badges */}
                      <div className="flex flex-wrap items-center gap-2 flex-1">
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
                        )}
                        {emailItem.is_primary && (
                          <span className="text-xs bg-white/15 text-white font-semibold px-2 py-0.5 rounded">
                            PRIMARY
                          </span>
                        )}
                        {isAuthIdentity && (
                          <CustomTooltip content="You are currently signed in with this email.">
                            <span className="text-xs text-amber-400/80 cursor-help">
                              🔑 Current Session
                            </span>
                          </CustomTooltip>
                        )}
                        {!emailItem.is_verified && (
                          <span className="text-xs text-yellow-500/80">⚠ Unverified</span>
                        )}
                      </div>

                      {/* Right: actions */}
                      <div className="flex items-center gap-2 shrink-0">
                        {!emailItem.is_primary && (
                          <Button
                            variant="outline"
                            onClick={() => canSetPrimary && handleSetPrimaryEmail(emailItem.id)}
                            disabled={!canSetPrimary}
                            className="border-white/20 hover:bg-white/10 text-xs h-8 px-3 disabled:opacity-30 disabled:cursor-not-allowed"
                          >
                            Set as Primary
                          </Button>
                        )}
                        
                        {!emailItem.is_verified && (
                          <Button
                            variant="outline"
                            onClick={() => {
                              setVerifyContact({ id: emailItem.id, value: emailItem.email, type: "emails" });
                            }}
                            className="border-yellow-500/30 text-yellow-500/80 hover:bg-yellow-500/10 hover:text-yellow-400 hover:border-yellow-500/50 text-xs h-8 px-3"
                          >
                            Verify
                          </Button>
                        )}

                        {!canRemove && removeDisabledReason ? (
                          <CustomTooltip content={removeDisabledReason} align="right">
                            <span>
                              <Button
                                variant="ghost"
                                disabled
                                className="text-white/40 text-xs h-8 px-3 disabled:opacity-30 disabled:cursor-not-allowed"
                              >
                                Remove
                              </Button>
                            </span>
                          </CustomTooltip>
                        ) : (
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
                      </div>
                    </div>

                  </ListCard>
                );
              })}
            </div>
          )}

          {/* Add Email inline form */}
          {showAddEmail ? (
            <div className="border border-white/10 rounded-lg p-4 space-y-3 bg-white/5">
              <FormField
                label="New Email Address"
                type="email"
                value={newEmail}
                onChange={(e) => setNewEmail(e.target.value)}
                error={newEmail && !isValidEmail(newEmail) ? "Please enter a valid email address." : undefined}
                className="bg-white/5 border-white/10 text-white"
              />
              {addEmailError && <p className="text-sm text-red-400 font-medium">{addEmailError}</p>}
              <div className="flex gap-2">
                <Button
                  onClick={handleAddEmail}
                  disabled={addEmailLoading || !newEmail || !isValidEmail(newEmail)}
                  className="text-sm"
                >
                  {addEmailLoading ? "Sending..." : "Send Verification"}
                </Button>
                <Button
                  variant="ghost"
                  onClick={() => { setShowAddEmail(false); setNewEmail(""); }}
                  className="text-white/50 hover:text-white text-sm"
                >
                  Cancel
                </Button>
              </div>
            </div>
          ) : (
            emails.length >= 2 ? (
              <p className="text-xs text-white/50 italic">Maximum of 2 emails allowed per account.</p>
            ) : (
              <Button
                variant="outline"
                onClick={() => setShowAddEmail(true)}
                className="border-white/20 hover:bg-white/10 text-sm"
              >
                + Add Email
              </Button>
            )
          )}
        </div>

        {/* ── Phone Numbers ── */}
        <div className="space-y-4">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-white/80 border-b border-white/10 pb-2">
            Phone Numbers
          </h3>

          {contactsLoading ? (
            <p className="text-sm text-white/40">Loading...</p>
          ) : (
            <div className="flex flex-col gap-2">
              {phones.map((phoneItem) => {
                const isAuthIdentity = phoneItem.phone === session?.user?.phone;
                const canRemove = !phoneItem.is_primary && phones.length > 1 && !isAuthIdentity;
                const canSetPrimary = !phoneItem.is_primary && phoneItem.is_verified;

                const removeDisabledReason = phoneItem.is_primary
                  ? "This is your primary phone. Set another verified phone as primary before removing this one."
                  : phones.length <= 1
                  ? "You must have at least one phone number on your account."
                  : isAuthIdentity
                  ? "You cannot remove the phone number you are currently signed in with."
                  : undefined;

                return (
                  <ListCard key={phoneItem.id} className="flex-col sm:flex-col items-stretch sm:items-stretch">
                    <div className="flex flex-col sm:flex-row sm:items-center justify-between w-full gap-3">
                      {/* Left: number + badges */}
                      <div className="flex flex-wrap items-center gap-2 flex-1">
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
                        )}
                        {phoneItem.is_primary && (
                          <span className="text-xs bg-white/15 text-white font-semibold px-2 py-0.5 rounded">
                            PRIMARY
                          </span>
                        )}
                        {isAuthIdentity && (
                          <CustomTooltip content="You are currently signed in with this phone number.">
                            <span className="text-xs text-amber-400/80 cursor-help">
                              🔑 Current Session
                            </span>
                          </CustomTooltip>
                        )}
                        {!phoneItem.is_verified && (
                          <span className="text-xs text-yellow-500/80">⚠ Unverified</span>
                        )}
                      </div>

                      {/* Right: actions */}
                      <div className="flex items-center gap-2 shrink-0">
                        {!phoneItem.is_primary && (
                          <Button
                            variant="outline"
                            onClick={() => canSetPrimary && handleSetPrimaryPhone(phoneItem.id)}
                            disabled={!canSetPrimary}
                            className="border-white/20 hover:bg-white/10 text-xs h-8 px-3 disabled:opacity-30 disabled:cursor-not-allowed"
                          >
                            Set as Primary
                          </Button>
                        )}
                        
                        {!phoneItem.is_verified && (
                          <Button
                            variant="outline"
                            onClick={() => {
                              setVerifyContact({ id: phoneItem.id, value: phoneItem.phone, type: "phones" });
                            }}
                            className="border-yellow-500/30 text-yellow-500/80 hover:bg-yellow-500/10 hover:text-yellow-400 hover:border-yellow-500/50 text-xs h-8 px-3"
                          >
                            Verify
                          </Button>
                        )}

                        {!canRemove && removeDisabledReason ? (
                          <CustomTooltip content={removeDisabledReason} align="right">
                            <span>
                              <Button
                                variant="ghost"
                                disabled
                                className="text-white/40 text-xs h-8 px-3 disabled:opacity-30 disabled:cursor-not-allowed"
                              >
                                Remove
                              </Button>
                            </span>
                          </CustomTooltip>
                        ) : (
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
                      </div>
                    </div>
                  </ListCard>
                );
              })}
            </div>
          )}

          {/* Add Phone inline form */}
          {showAddPhone ? (
            <div className="border border-white/10 rounded-lg p-4 space-y-3 bg-white/5">
              <FormField
                label="New Phone Number"
                type="tel" prefixText="+1"
                value={newPhone}
                onChange={(e) => setNewPhone(formatPhone(e.target.value))}
                error={newPhone && !isValidPhone(newPhone) ? "Please enter a valid 10-digit phone number." : undefined}
                className="bg-white/5 border-white/10 text-white"
              />
              {addPhoneError && <p className="text-sm text-red-400 font-medium">{addPhoneError}</p>}
              <div className="flex gap-2">
                <Button
                  onClick={handleAddPhone}
                  disabled={addPhoneLoading || !newPhone || !isValidPhone(newPhone)}
                  className="text-sm"
                >
                  {addPhoneLoading ? "Sending..." : "Send Verification"}
                </Button>
                <Button
                  variant="ghost"
                  onClick={() => { setShowAddPhone(false); setNewPhone(""); }}
                  className="text-white/50 hover:text-white text-sm"
                >
                  Cancel
                </Button>
              </div>
            </div>
          ) : (
            phones.length >= 2 ? (
              <p className="text-xs text-white/50 italic">Maximum of 2 phone numbers allowed per account.</p>
            ) : (
              <Button
                variant="outline"
                onClick={() => setShowAddPhone(true)}
                className="border-white/20 hover:bg-white/10 text-sm"
              >
                + Add Phone
              </Button>
            )
          )}
        </div>

        {/* ── Linked Accounts ── */}
        <div className="space-y-4">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-white/80 border-b border-white/10 pb-2">Linked Accounts</h3>
          <p className="text-xs text-white/50 leading-relaxed max-w-2xl">
            You can sign in to your Limitless Hoops account using any linked provider.{" "}
            <strong className="text-white/70 block mt-1">
              Important: In order to link an external account (like Facebook), it cannot already be associated with another Limitless Hoops account. If it is, you must log into that account and unlink it first.
            </strong>
          </p>

          <div className="flex flex-col gap-3">
            {(["google", "facebook"] as const).map((provider) => {
              const providerName = provider.charAt(0).toUpperCase() + provider.slice(1);
              const providerIdentities = identities.filter((i) => i.provider === provider);
              
              // Helper to render the provider icon
              const ProviderIcon = () => (
                <div className="bg-white p-1.5 rounded-full flex items-center justify-center w-8 h-8 shrink-0">
                  {provider === "google" && (
                    <svg className="h-5 w-5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                      <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                      <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                      <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                    </svg>
                  )}
                  {provider === "facebook" && (
                    <svg className="h-5 w-5" viewBox="0 0 24 24" fill="#1877F2" xmlns="http://www.w3.org/2000/svg"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.469h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.469h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
                  )}

                </div>
              );

              // If not connected, show the connect row
              if (providerIdentities.length === 0) {
                return (
                  <ListCard key={provider}>
                    <div className="flex items-center gap-4">
                      <ProviderIcon />
                      <div>
                        <p className="font-semibold text-white">{providerName}</p>
                        <p className="text-xs text-white/30">Not Connected</p>
                      </div>
                    </div>
                    <Button
                      variant="outline"
                      className="border-white/20 hover:bg-white/10"
                      onClick={() => supabase.auth.linkIdentity({ 
                        provider: provider as any, 
                        options: { redirectTo: window.location.origin + '/account' } 
                      })}
                    >
                      Connect
                    </Button>
                  </ListCard>
                );
              }

              // If connected, show a row for EACH connected identity of this provider
              return providerIdentities.map((identity) => {
                const identityEmail = identity.identity_data?.email;
                const identityName = identity.identity_data?.full_name || identity.identity_data?.name;
                const disconnectDisabled = identities.length <= 1;

                return (
                  <ListCard key={identity.identity_id}>
                    <div className="flex items-center gap-4">
                      <ProviderIcon />
                      <div>
                        <p className="font-semibold text-white">{providerName}</p>
                        <div className="flex flex-col mt-0.5">
                          <p className="text-xs text-green-400 font-medium flex items-center gap-1">
                            <span className="h-1.5 w-1.5 bg-green-500 rounded-full inline-block"></span>
                            Connected
                          </p>
                          {(identityName || identityEmail) && (
                            <p className="text-[11px] text-white/50 mt-1">
                              {identityName && <span>{identityName}</span>}
                              {identityName && identityEmail && <span className="mx-1.5 opacity-50">•</span>}
                              {identityEmail && <span>{identityEmail}</span>}
                            </p>
                          )}
                        </div>
                      </div>
                    </div>
                    {disconnectDisabled ? (
                      <CustomTooltip content="Since you created your account using this provider, it is your permanent root identity and cannot be disconnected unless you link another provider like Facebook." align="right">
                        <span>
                          <Button
                            variant="outline"
                            disabled
                            className="border-white/20 disabled:opacity-40 disabled:cursor-not-allowed"
                          >
                            Disconnect
                          </Button>
                        </span>
                      </CustomTooltip>
                    ) : (
                      <Button
                        variant="outline"
                        onClick={() => {
                          // Note: we still just pass the provider string (e.g. 'google'). 
                          // If they have multiple, the modal's unlinkIdentity might grab the first one, 
                          // but typically they only have one per provider.
                          setDisconnectProvider(provider);
                          setIsDisconnectOpen(true);
                        }}
                        className="border-white/20 hover:bg-red-900/30 hover:text-red-400 hover:border-red-900/50"
                      >
                        Disconnect
                      </Button>
                    )}
                  </ListCard>
                );
              });
            })}
          </div>
        </div>
      </CardContent>

      <DisconnectSocialModal
        isOpen={isDisconnectOpen}
        onClose={() => setIsDisconnectOpen(false)}
        provider={disconnectProvider}
        onSuccess={() => fetchContacts()}
      />
      {verifyContact && (
        <VerifyContactModal
          isOpen={true}
          onClose={() => setVerifyContact(null)}
          contactId={verifyContact.id}
          contactValue={verifyContact.value}
          contactType={verifyContact.type}
          onSuccess={() => fetchContacts()}
        />
      )}
    </Card>
  );
}

// ─── NotificationsTab ─────────────────────────────────────────────────────────

const NOTIFICATION_CATEGORIES: { key: NotificationPreference["category"]; label: string }[] = [
  { key: "billing", label: "Billing" },
  { key: "events", label: "Events" },
  { key: "announcements", label: "Announcements" },
  { key: "security", label: "Security" },
];

function NotificationsTab({ household }: { household: any }) {
  const { session } = useAuth();
  
  const [preferences, setPreferences] = useState<NotificationPreference[]>([]);
  const [notifLoading, setNotifLoading] = useState(false);

  const authHeader: Record<string, string> = session?.access_token
    ? { Authorization: `Bearer ${session.access_token}` }
    : {};

  useEffect(() => {
    const fetchPreferences = async () => {
      if (!session?.access_token) return;
      setNotifLoading(true);
      try {
        const res = await fetch("http://localhost:3000/api/users/me/notifications", {
          headers: authHeader,
        });
        if (res.ok) setPreferences(await res.json());
      } catch (err) {
        console.error("Failed to fetch notification preferences:", err);
      } finally {
        setNotifLoading(false);
      }
    };
    fetchPreferences();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [session?.access_token]);

  const handleToggle = async (pref: NotificationPreference) => {
    // Security notifications on the primary contact cannot be disabled
    if (pref.category === "security" && pref.is_primary) return;

    const newEnabled = !pref.enabled;
    // Optimistic update
    setPreferences((prev) =>
      prev.map((p) => (p.id === pref.id ? { ...p, enabled: newEnabled } : p))
    );

    try {
      await fetch(`http://localhost:3000/api/users/me/notifications/${pref.id}`, {
        method: "PUT",
        headers: { ...authHeader, "Content-Type": "application/json" },
        body: JSON.stringify({ enabled: newEnabled }),
      });
    } catch (err) {
      console.error("Failed to update preference:", err);
      // Revert on error
      setPreferences((prev) =>
        prev.map((p) => (p.id === pref.id ? { ...p, enabled: pref.enabled } : p))
      );
    }
  };

  // Deduplicate contact IDs preserving order
  const contactIds = Array.from(new Set(preferences.map((p) => p.target_id)));

  // Other guardians in the household (everyone except index 0, the current user)
  const otherGuardians = household?.users?.slice(1) ?? [];

  return (
    <div className="space-y-8">
      {/* Section 1 — My Notification Preferences */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div>
            <h2 className="text-xl font-bold text-white mb-1">My Notification Preferences</h2>
            <p className="text-sm text-white/50">
              Choose which notifications each of your contacts receives.
            </p>
          </div>

          {notifLoading ? (
            <p className="text-sm text-white/40">Loading preferences...</p>
          ) : contactIds.length === 0 ? (
            <p className="text-sm text-white/40 italic">
              No notification preferences found. Add emails or phone numbers in Login &amp; Security.
            </p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full text-sm">
                <thead>
                  <tr className="border-b border-white/10">
                    <th className="text-left text-white/60 font-medium pb-3 pr-4 w-1/3">Contact</th>
                    {NOTIFICATION_CATEGORIES.map((cat) => (
                      <th key={cat.key} className="text-center text-white/60 font-medium pb-3 px-2">
                        {cat.label}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody>
                  {contactIds.map((contactId) => {
                    const contactPrefs = preferences.filter((p) => p.target_id === contactId);
                    const sample = contactPrefs[0];
                    const contactLabel = sample?.contact_value ?? contactId;
                    const isPrimary = sample?.is_primary ?? false;

                    return (
                      <tr key={contactId} className="border-b border-white/5 last:border-0">
                        <td className="py-4 pr-4">
                          <div className="flex flex-wrap items-center gap-2">
                            <span className="text-white font-medium">{contactLabel}</span>
                            {isPrimary && (
                              <span className="text-xs bg-white/15 text-white font-semibold px-2 py-0.5 rounded">
                                PRIMARY
                              </span>
                            )}
                          </div>
                        </td>
                        {NOTIFICATION_CATEGORIES.map((cat) => {
                          const pref = contactPrefs.find((p) => p.category === cat.key);
                          const isSecurityPrimary = cat.key === "security" && isPrimary;

                          return (
                            <td key={cat.key} className="py-4 px-2 text-center">
                              {isSecurityPrimary ? (
                                <CustomTooltip content="Cannot disable security notifications on your primary contact." align="right">
                                  <span className="text-lg cursor-help">
                                    🔒
                                  </span>
                                </CustomTooltip>
                              ) : pref ? (
                                <label className="inline-flex items-center cursor-pointer">
                                  <input
                                    type="checkbox"
                                    checked={pref.enabled}
                                    onChange={() => handleToggle(pref)}
                                    className="w-4 h-4 accent-primary cursor-pointer"
                                  />
                                </label>
                              ) : (
                                <span className="text-white/20">—</span>
                              )}
                            </td>
                          );
                        })}
                      </tr>
                    );
                  })}
                </tbody>
              </table>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Section 2 — Household Contacts */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div>
            <h2 className="text-xl font-bold text-white mb-1">Household Contacts</h2>
            <p className="text-sm text-white/50">
              Other guardians manage their own notification preferences independently.
            </p>
          </div>

          {otherGuardians.length === 0 ? (
            <p className="text-sm text-white/40 italic">
              No other guardians in this household.
            </p>
          ) : (
            <div className="flex flex-col gap-3">
              {otherGuardians.map((user: any) => (
                <ListCard key={user.id}>
                  <div>
                    <p className="font-semibold text-white">
                      {user.first_name} {user.last_name}
                    </p>
                    <p className="text-xs text-white/40 mt-0.5">
                      Manages their own notification preferences.
                    </p>
                  </div>
                </ListCard>
              ))}
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}

// ─── HouseholdTab ─────────────────────────────────────────────────────────────

function HouseholdTab({ household, fetchHousehold, loading }: { household: any, fetchHousehold: () => void, loading: boolean }) {
  
  const [isAddGuardianOpen, setIsAddGuardianOpen] = useState(false);
  const [isAddPlayerOpen, setIsAddPlayerOpen] = useState(false);
  const [guardianToRemove, setGuardianToRemove] = useState<any>(null);
  const [playerToRemove, setPlayerToRemove] = useState<string | null>(null);

  const handleAddPlayerClose = () => {
    setIsAddPlayerOpen(false);
    fetchHousehold(); // Refresh data!
  };

  if (loading) {
    return <div className="text-white">Loading household data...</div>;
  }

  return (
    <div className="space-y-8">
      {/* Primary Account & Notification Contacts */}
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
      </Card>

      {/* Dependents */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <div>
              <h2 className="text-xl font-bold text-white mb-1">Dependents</h2>
              <p className="text-sm text-white/50">Manage the players in your household.</p>
            </div>
            <Button onClick={() => setIsAddPlayerOpen(true)}>+ Add Player</Button>
          </div>

          {household?.players?.map((dep: any) => (
            <ListCard key={dep.id}>
              <div>
                <p className="font-semibold text-white">{dep.first_name} {dep.last_name}</p>
                <p className="text-xs text-white/50">Graduation Year {dep.graduation_year}</p>
              </div>
              <Button variant="ghost" onClick={() => setPlayerToRemove(dep.first_name)} className="text-white/40 hover:text-red-400 hover:bg-red-900/20">Remove</Button>
            </ListCard>
          ))}
          {(!household?.players || household.players.length === 0) && (
            <p className="text-sm text-white/40 italic">No players added yet.</p>
          )}
        </CardContent>
      </Card>

      <AddGuardianModal isOpen={isAddGuardianOpen} onClose={() => { setIsAddGuardianOpen(false); fetchHousehold(); }} players={household?.players || []} />
      <RemoveGuardianModal isOpen={!!guardianToRemove} onClose={() => setGuardianToRemove(null)} contact={guardianToRemove} />
      
      <AddPlayerModal isOpen={isAddPlayerOpen} onClose={handleAddPlayerClose} householdId={household?.id} />
      <RemovePlayerModal isOpen={!!playerToRemove} onClose={() => setPlayerToRemove(null)} playerName={playerToRemove || ""} />
    </div>
  );
}

// ─── BillingTab ───────────────────────────────────────────────────────────────

function BillingTab({ household }: { household: any }) {
  const [stripeAction, setStripeAction] = useState<"upgrade" | "cancel" | null>(null);
  const [loading, setLoading] = useState(false);

  const handleCheckout = async (planType: string) => {
    const playerId = household?.players?.[0]?.id;
    if (!playerId) {
      alert("Please add a player to your household before upgrading your plan.");
      return;
    }

    setLoading(true);
    try {
      const res = await fetch("http://localhost:3000/api/checkout", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          player_id: playerId,
          plan_type: planType,
        }),
      });

      if (!res.ok) throw new Error("Failed to create checkout session");

      const data = await res.json();
      if (data.url) {
        window.location.href = data.url; // Redirect to Stripe
      }
    } catch (err) {
      console.error(err);
      alert("Error initiating checkout");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card className={cardClass}>
      <CardContent className="p-6 sm:p-8 space-y-6">
        <div>
          <h2 className="text-xl font-bold text-white mb-1">Subscription & Billing</h2>
          <p className="text-sm text-white/50">Manage your membership plan and payment methods.</p>
        </div>

        <div className="bg-linear-to-r from-primary/20 to-transparent border border-primary/30 rounded-lg p-6">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
            <div>
              <h3 className="text-sm font-bold uppercase tracking-widest text-primary mb-1">Current Plan</h3>
              <p className="text-2xl font-extrabold text-white">Player Roster Spot (Early Pricing)</p>
              <p className="text-sm text-white/60 mt-1">$100.00 / week (15 Installments) &bull; Next payment Oct 1, 2026</p>
            </div>
            <Button 
              variant="outline" 
              onClick={() => handleCheckout("player_early_installments")} 
              disabled={loading}
              className="w-full md:w-auto border-white/20 hover:bg-white/10"
            >
              {loading ? "Redirecting..." : "Manage Plan"}
            </Button>
          </div>
        </div>

        <div className="pt-4 border-t border-white/10 flex justify-end">
          <Button variant="ghost" onClick={() => handleCheckout("volunteer_200_upfront")} className="text-white/40 hover:text-white hover:bg-white/5">Pay $200 Deposit</Button>
        </div>
      </CardContent>

      <StripeMockModal 
        isOpen={!!stripeAction} 
        onClose={() => setStripeAction(null)} 
        action={stripeAction || "upgrade"} 
      />
    </Card>
  );
}

// ─── AdvancedTab ──────────────────────────────────────────────────────────────

function AdvancedTab() {
  const [isDeleteOpen, setIsDeleteOpen] = useState(false);

  return (
    <Card className="border-red-900/30 bg-black/40 backdrop-blur-2xl shadow-xl shadow-black/50">
      <CardContent className="p-6 sm:p-8 space-y-6">
        <div>
          <h2 className="text-xl font-bold text-red-500 mb-1">Danger Zone</h2>
          <p className="text-sm text-white/50">Irreversible and destructive actions.</p>
        </div>

        <div className="border border-red-900/50 bg-red-950/20 rounded-lg p-6 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <div>
            <h3 className="font-semibold text-white">Delete Account</h3>
            <p className="text-xs text-white/50 mt-1 max-w-md">
              This action will permanently delete your account and archive all dependent data. You will immediately lose access to the platform.
            </p>
          </div>
          <Button variant="destructive" onClick={() => setIsDeleteOpen(true)} className="whitespace-nowrap shrink-0">Delete Account</Button>
        </div>
      </CardContent>

      <DeleteAccountModal isOpen={isDeleteOpen} onClose={() => setIsDeleteOpen(false)} />
    </Card>
  );
}

// ─── Page ─────────────────────────────────────────────────────────────────────

export default function AccountPage() {
  const [searchParams, setSearchParams] = useSearchParams();
  const activeTab = searchParams.get("tab") || "personal";
  const setActiveTab = (tab: string) => setSearchParams({ tab }, { replace: true });
  const { household, loading, refreshHousehold: fetchHousehold } = useAuth();
  const [globalLinkError, setGlobalLinkError] = useState("");

  useEffect(() => {
    // Parse OAuth errors from the hash (e.g. from a failed identity link attempt)
    if (window.location.hash) {
      const hashParams = new URLSearchParams(window.location.hash.substring(1));
      const error = hashParams.get("error");
      const errorDesc = hashParams.get("error_description");
      if (error || errorDesc) {
        setGlobalLinkError(errorDesc ? decodeURIComponent(errorDesc).replace(/\+/g, " ") : "Failed to link identity.");
        // Force the user to the security tab so they see where the error came from
        setActiveTab("security");
        window.history.replaceState(null, "", window.location.pathname + window.location.search);
      }
    }
  }, []);

  return (
    <DashboardLayout>
      {/* Global Error Alert */}
      {globalLinkError && (
        <div className="w-full mb-8">
          <div className="bg-red-500/10 border border-red-500/50 text-red-400 px-4 py-3 rounded-lg flex items-center justify-between shadow-sm animate-in fade-in slide-in-from-top-2 duration-300">
            <div className="flex items-center gap-3">
              <svg className="w-5 h-5 text-red-500 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
              <div>
                <h3 className="text-sm font-bold text-white">Identity Linking Failed</h3>
                <p className="text-sm mt-0.5">{globalLinkError}</p>
              </div>
            </div>
            <button onClick={() => setGlobalLinkError("")} className="text-white/50 hover:text-white transition-colors p-2">
              <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M6 18L18 6M6 6l12 12" /></svg>
            </button>
          </div>
        </div>
      )}
      <div className="flex flex-col lg:flex-row gap-8 pb-12">

        {/* Sidebar Nav */}
        <aside className="w-full lg:w-64 shrink-0">
          <div className="sticky top-24">
            <h1 className="text-2xl font-bold tracking-tight text-white uppercase mb-6 hidden lg:block">Account Settings</h1>

            <nav className="flex lg:flex-col gap-2 overflow-x-auto pb-4 lg:pb-0">
              {TABS.map((tab) => (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`cursor-pointer whitespace-nowrap rounded-lg px-4 py-3 text-left text-sm font-medium transition-colors ${
                    activeTab === tab.id
                      ? "bg-primary/20 text-white border border-primary/50"
                      : "text-white/60 hover:bg-white/5 hover:text-white border border-transparent"
                  }`}
                >
                  {tab.label}
                </button>
              ))}
            </nav>
          </div>
        </aside>

        {/* Main Content Area */}
        <div className="flex-1 min-w-0">
          <h1 className="text-2xl font-bold tracking-tight text-white uppercase mb-6 lg:hidden">Account Settings</h1>

          <div key={activeTab} className="animate-in fade-in slide-in-from-bottom-4 duration-300">
            {activeTab === "personal"       && <PersonalInfoTab household={household} />}
            {activeTab === "security"       && <SecurityTab household={household} />}
            {activeTab === "household"      && <HouseholdTab household={household} fetchHousehold={fetchHousehold} loading={loading} />}
            {activeTab === "notifications"  && <NotificationsTab household={household} />}
            {activeTab === "billing"        && <BillingTab household={household} />}
            {activeTab === "advanced"       && <AdvancedTab />}
          </div>
        </div>

      </div>
    </DashboardLayout>
  );
}

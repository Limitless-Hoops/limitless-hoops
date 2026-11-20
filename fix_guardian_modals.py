import re

with open('apps/app/src/modules/account/components/modals/GuardianModals.tsx', 'r') as f:
    content = f.read()

old_add_guardian = """interface AddGuardianModalProps {
  isOpen: boolean;
  onClose: () => void;
  householdId?: string;
}

export function AddGuardianModal({ isOpen, onClose, householdId }: AddGuardianModalProps) {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [dob, setDob] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSave = async () => {
    if (!householdId) {
      console.warn("No household ID found");
      onClose();
      return;
    }
    
    setIsSubmitting(true);
    try {
      const { data: { session } } = await supabase.auth.getSession();
      const res = await fetch(`http://localhost:3000/api/households/${householdId}/guardians`, {
        method: "POST",
        headers: { 
          "Content-Type": "application/json",
          "Authorization": `Bearer ${session?.access_token}`
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          dob: dob,
          email,
          phone: normalizePhoneForAPI(phone),
        }),
      });

      if (!res.ok) throw new Error("Failed to add guardian");
      
      setFirstName("");
      setLastName("");
      setDob("");
      setEmail("");
      setPhone("");
      onClose();
    } catch (err) {
      console.error(err);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Dialog isOpen={isOpen} onClose={onClose} title="Add Guardian" description="Add a co-parent or guardian to receive notifications and manage dependents.">
      <div className="mt-6 space-y-4">
        <div className="grid grid-cols-2 gap-4">
          <FormField label="First Name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
          <FormField label="Last Name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
        </div>
        <FormField label="Date of Birth" type="date" max={getDateOfBirthMax(18)} value={dob} onChange={(e) => setDob(e.target.value)} />
        <FormField label="Email Address" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        <FormField label="Phone Number" type="tel" value={phone} onChange={(e) => setPhone(formatPhone(e.target.value))} />
        <div className="mt-8 flex justify-end gap-3">
          <Button variant="ghost" onClick={onClose} disabled={isSubmitting} className="text-white hover:bg-white/10">Cancel</Button>
          <Button onClick={handleSave} disabled={isSubmitting}>
            {isSubmitting ? "Adding..." : "Add Guardian"}
          </Button>
        </div>
      </div>
    </Dialog>
  );
}"""

new_add_guardian = """interface AddGuardianModalProps {
  isOpen: boolean;
  onClose: () => void;
  players?: any[];
}

export function AddGuardianModal({ isOpen, onClose, players = [] }: AddGuardianModalProps) {
  const [playerId, setPlayerId] = useState("");
  const [relationship, setRelationship] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSave = async () => {
    if (!playerId) {
      return;
    }
    
    setIsSubmitting(true);
    try {
      const { data: { session } } = await supabase.auth.getSession();
      const res = await fetch(`http://localhost:3000/api/players/${playerId}/contacts`, {
        method: "POST",
        headers: { 
          "Content-Type": "application/json",
          "Authorization": `Bearer ${session?.access_token}`
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          relationship: relationship,
          email,
          phone: normalizePhoneForAPI(phone),
          receive_sms: true,
          receive_email: true
        }),
      });

      if (!res.ok) throw new Error("Failed to add contact");
      
      setFirstName("");
      setLastName("");
      setRelationship("");
      setPlayerId("");
      setEmail("");
      setPhone("");
      onClose();
    } catch (err) {
      console.error(err);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Dialog isOpen={isOpen} onClose={onClose} title="Add Notification Contact" description="Add an emergency contact or family member to receive updates for a dependent.">
      <div className="mt-6 space-y-4">
        <div>
          <label className="block text-sm font-medium text-white/70 mb-2">Select Player</label>
          <select 
            className="w-full bg-black/40 border border-white/20 rounded-lg px-4 py-2.5 text-white placeholder:text-white/30 focus:outline-none focus:ring-2 focus:ring-primary/50 transition-shadow appearance-none"
            value={playerId}
            onChange={(e) => setPlayerId(e.target.value)}
          >
            <option value="" disabled>Choose a dependent...</option>
            {players.map(p => <option key={p.id} value={p.id}>{p.first_name} {p.last_name}</option>)}
          </select>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <FormField label="First Name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
          <FormField label="Last Name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
        </div>
        <FormField label="Relationship (e.g. Grandparent)" value={relationship} onChange={(e) => setRelationship(e.target.value)} />
        <FormField label="Email Address" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        <FormField label="Phone Number" type="tel" value={phone} onChange={(e) => setPhone(formatPhone(e.target.value))} />
        <div className="mt-8 flex justify-end gap-3">
          <Button variant="ghost" onClick={onClose} disabled={isSubmitting} className="text-white hover:bg-white/10">Cancel</Button>
          <Button onClick={handleSave} disabled={isSubmitting || !playerId}>
            {isSubmitting ? "Adding..." : "Add Contact"}
          </Button>
        </div>
      </div>
    </Dialog>
  );
}"""

content = content.replace(old_add_guardian, new_add_guardian)

old_remove_guardian = """interface RemoveGuardianModalProps {
  isOpen: boolean;
  onClose: () => void;
  guardianName: string;
}

export function RemoveGuardianModal({ isOpen, onClose, guardianName }: RemoveGuardianModalProps) {
  const handleRemove = () => {
    console.log("Removing guardian", guardianName);
    onClose();
  };

  return (
    <Dialog isOpen={isOpen} onClose={onClose} title="Remove Guardian" description={`Are you sure you want to remove ${guardianName} from your household? They will lose access to manage players.`}>
      <div className="mt-8 flex justify-end gap-3">
        <Button variant="ghost" onClick={onClose} className="text-white hover:bg-white/10">Cancel</Button>
        <Button variant="destructive" onClick={handleRemove}>Remove Guardian</Button>
      </div>
    </Dialog>
  );
}"""

new_remove_guardian = """interface RemoveGuardianModalProps {
  isOpen: boolean;
  onClose: () => void;
  contact: any;
}

export function RemoveGuardianModal({ isOpen, onClose, contact }: RemoveGuardianModalProps) {
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleRemove = async () => {
    if (!contact) return;
    setIsSubmitting(true);
    try {
      const { data: { session } } = await supabase.auth.getSession();
      const res = await fetch(`http://localhost:3000/api/players/${contact.player_id}/contacts/${contact.id}`, {
        method: "DELETE",
        headers: { "Authorization": `Bearer ${session?.access_token}` },
      });
      if (!res.ok) throw new Error("Failed to remove contact");
      onClose();
    } catch (err) {
      console.error(err);
    } finally {
      setIsSubmitting(false);
    }
  };

  if (!contact) return null;

  return (
    <Dialog isOpen={isOpen} onClose={onClose} title="Remove Contact" description={`Are you sure you want to remove ${contact.first_name} ${contact.last_name}? They will no longer receive notifications for ${contact.playerName}.`}>
      <div className="mt-8 flex justify-end gap-3">
        <Button variant="ghost" onClick={onClose} disabled={isSubmitting} className="text-white hover:bg-white/10">Cancel</Button>
        <Button variant="destructive" onClick={handleRemove} disabled={isSubmitting}>
          {isSubmitting ? "Removing..." : "Remove Contact"}
        </Button>
      </div>
    </Dialog>
  );
}"""

content = content.replace(old_remove_guardian, new_remove_guardian)

with open('apps/app/src/modules/account/components/modals/GuardianModals.tsx', 'w') as f:
    f.write(content)


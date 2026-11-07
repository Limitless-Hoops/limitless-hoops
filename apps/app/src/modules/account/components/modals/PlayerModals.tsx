import { useState } from "react";
import { Dialog, Button, FormField } from "@lh/ui";
import { getDateOfBirthMax, formatPhone, normalizePhoneForAPI } from "@lh/utils";
import { supabase } from "@/lib/supabase";

interface AddPlayerModalProps {
  isOpen: boolean;
  onClose: () => void;
  householdId?: string; // Passed down from dashboard
}

export function AddPlayerModal({ isOpen, onClose, householdId }: AddPlayerModalProps) {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [dob, setDob] = useState("");
  const [graduationYear, setGraduationYear] = useState("");
  const [relationship, setRelationship] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [_isSubmitting, setIsSubmitting] = useState(false);

  const handleSave = async () => {
    if (!householdId) {
      console.warn("No household ID found");
      onClose();
      return;
    }
    
    setIsSubmitting(true);
    try {
      const { data: { session } } = await supabase.auth.getSession();
      const res = await fetch(`http://localhost:3000/api/households/${householdId}/players`, {
        method: "POST",
        headers: { 
          "Content-Type": "application/json",
          "Authorization": `Bearer ${session?.access_token}`
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          dob: dob,
          graduation_year: graduationYear,
          email,
          phone: normalizePhoneForAPI(phone),
          relationship,
        }),
      });

      if (!res.ok) throw new Error("Failed to add player");

      console.log("Player added successfully!");
      setFirstName("");
      setLastName("");
      setDob("");
      setGraduationYear("");
      setRelationship("");
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
    <Dialog isOpen={isOpen} onClose={onClose} title="Add Player" description="Add a dependent to your household to register them for events.">
      <div className="mt-6 space-y-4">
        <div className="grid grid-cols-2 gap-4">
          <FormField label="First Name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
          <FormField label="Last Name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
        </div>
        <FormField label="Date of Birth" type="date" value={dob} max={getDateOfBirthMax(0)} onChange={(e) => setDob(e.target.value)} />
        <FormField label="Email Address (Optional)" type="email" value={email} onChange={(e) => setEmail(e.target.value)} hint="If provided, player will receive their own notifications." />
        <FormField label="Phone Number (Optional)" type="tel" prefixText="+1" value={phone} onChange={(e) => setPhone(formatPhone(e.target.value))} />
        <div className="mt-8 flex justify-end gap-3">
          <Button variant="ghost" onClick={onClose} className="text-white hover:bg-white/10">Cancel</Button>
          <Button onClick={handleSave}>Add Player</Button>
        </div>
      </div>
    </Dialog>
  );
}

interface RemovePlayerModalProps {
  isOpen: boolean;
  onClose: () => void;
  playerName: string;
}

export function RemovePlayerModal({ isOpen, onClose, playerName }: RemovePlayerModalProps) {
  const handleRemove = () => {
    console.log("Removing player", playerName);
    onClose();
  };

  return (
    <Dialog isOpen={isOpen} onClose={onClose} title="Remove Player" description={`Are you sure you want to remove ${playerName} from your household? This may affect their active registrations.`}>
      <div className="mt-8 flex justify-end gap-3">
        <Button variant="ghost" onClick={onClose} className="text-white hover:bg-white/10">Cancel</Button>
        <Button variant="destructive" onClick={handleRemove}>Remove Player</Button>
      </div>
    </Dialog>
  );
}

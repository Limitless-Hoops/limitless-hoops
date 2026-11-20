import { useState } from "react";
import { Dialog, Button, FormField } from "@lh/ui";

interface DeleteAccountModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export function DeleteAccountModal({ isOpen, onClose }: DeleteAccountModalProps) {
  const [confirmText, setConfirmText] = useState("");

  const handleDelete = () => {
    console.log("Deleting account permanently");
    onClose();
  };

  return (
    <Dialog 
      isOpen={isOpen} 
      onClose={onClose}
      title="Delete Account"
      description="This action is permanent and cannot be undone. You will lose access to the platform immediately, and no refunds will be issued for active memberships."
    >
      <div className="mt-6 space-y-4">
        <div className="rounded-lg bg-red-950/30 border border-red-900/50 p-4">
          <p className="text-sm text-red-200">
            If you just want to stop being charged, you can <span className="font-bold underline cursor-pointer hover:text-white" onClick={() => console.log('redirect to Stripe')}>cancel your subscription</span> instead.
          </p>
        </div>

        <FormField
          label='Type "DELETE" to confirm'
          value={confirmText}
          onChange={(e) => setConfirmText(e.target.value)}
          placeholder="DELETE"
        />

        <div className="mt-8 flex justify-end gap-3">
          <Button variant="ghost" onClick={onClose} className="text-white hover:bg-white/10">
            Keep Account
          </Button>
          <Button 
            variant="destructive" 
            onClick={handleDelete}
            disabled={confirmText !== "DELETE"}
          >
            Permanently Delete
          </Button>
        </div>
      </div>
    </Dialog>
  );
}
